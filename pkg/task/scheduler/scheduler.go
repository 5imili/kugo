package scheduler

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/5imili/kugo/pkg/dao"
	daotypes "github.com/5imili/kugo/pkg/dao/mysql/types"
	"github.com/5imili/kugo/pkg/task/types"
	"github.com/5imili/kugo/pkg/task/utils"
	"github.com/leopoldxx/go-utils/trace"
)

const (
	defaultScheduleCycle = time.Second * 1
	defaultLockPrefix    = "/github.com/reboot/kugo/pkg/task/scheduler/lock"
)

// Scheduler will schedule different types and states task into their corresponding processor
type Scheduler interface {
	GetName() string
	Init(cfg types.InitConfigs) error
	Schedule(ctx context.Context, task *types.Task) error
}

//Manager xx
type Manager struct {
	schedulers map[string]Scheduler
	dao        dao.Storage
	lockPrefix string
	ticker     *time.Ticker
	ctx        context.Context
	cancel     context.CancelFunc
	wg         sync.WaitGroup
}

type options struct {
	lockPrefix    string
	scheduleCycle time.Duration
}

//Option config scheduler manager
type Option func(opt *options)

// WithLockPrefix will set user defined lock prefix
func WithLockPrefix(prefix string) Option {
	return func(opt *options) {
		opt.lockPrefix = prefix
	}
}

// WithScheduleCycle will set user defined shedule cycle
func WithScheduleCycle(duration time.Duration) Option {
	return func(opt *options) {
		opt.scheduleCycle = duration
	}
}

// NewManager creates a Scheduler manager
func NewManager(ctx context.Context,
	dao dao.Storage) (*Manager, error) {
	ops := &options{
		lockPrefix:    defaultLockPrefix,
		scheduleCycle: defaultScheduleCycle,
	}
	ctx, cancel := context.WithCancel(trace.WithTraceForContext(ctx, "task-scheduler"))
	return &Manager{
		schedulers: map[string]Scheduler{},
		ctx:        ctx,
		cancel:     cancel,
		dao:        dao,
		lockPrefix: ops.lockPrefix,
		ticker:     time.NewTicker(ops.scheduleCycle),
	}, nil
}

// InitSchedulers will init all resource schedulers
func (m *Manager) InitSchedulers(schedulers ...Scheduler) error {
	for _, scheduler := range schedulers {
		if err := scheduler.Init(types.InitConfigs{
			Dao: m.dao,
		}); err != nil {
			return err
		}
		m.schedulers[scheduler.GetName()] = scheduler
	}
	return nil
}

func (m *Manager) Stop() {
	if m == nil {
		return
	}
	tracer := trace.GetTraceFromContext(m.ctx)
	m.ticker.Stop()
	m.cancel()
	tracer.Info("stopping scheduler mananger")
	m.wg.Wait()
	tracer.Info("stopped scheduler manager")
}

// Schedule a task to a specific scheduler
func (m *Manager) Schedule() error {
	if m == nil {
		return errors.New("task scheduler is not created")
	}
	// get task from db
	if m.dao == nil {
		return errors.New("nil dao in task scheduler")
	}
	tracer := trace.GetTraceFromContext(m.ctx)
	tracer.Info("start scheduling the pending task")
	for range m.ticker.C {
		tasks, err := m.dao.ListOpenTasks(m.ctx)
		if err != nil {
			if err == context.Canceled {
				tracer.Info("task scheduler has been stopped.")
				break
			}
			tracer.Warnf("list open task failed: %s", err)
			time.Sleep(time.Second)
			continue
		}
		tracer.Infof("total get %d pending tasks", len(tasks))
		for _, task := range tasks {
			scheduler, exists := m.schedulers[task.Resource]
			if !exists {
				tracer.Warnf("resource [%s]'s scheduler not exists: %v",
					task.Resource, task)
				continue
			}
			m.wg.Add(1)
			go func(task daotypes.Task) {
				defer m.wg.Done()
				newCtx := trace.WithTraceForContext(m.ctx,
					fmt.Sprintf("schedTask:%s:%s:%d", task.Resource, task.Type, task.ID))
				tracer := trace.GetTraceFromContext(newCtx)
				lockKey := fmt.Sprintf("%s/%d", m.lockPrefix, task.ID)
				// unlock, newCtx2, err := m.locker.Trylock(newCtx, lockKey)
				// if err != nil {
				// 	if err != context.DeadlineExceeded {
				// 		tracer.Errorf("lock task failed:%s", err)
				// 	}
				// 	return
				// }
				tracer.Infof("task has been locked by %s \n", lockKey)
				// defer func() {
				// 	time.Sleep(time.Second)
				// 	unlock()
				// 	tracer.Info("task has been unlocked")
				// }()
				newTask, err := m.dao.GetOpenTaskByTaskID(m.ctx, task.ID)
				if err != nil {
					tracer.Errorf("task can not be scheduled now: %d, err: %v", task.ID, err)
					return
				}
				combinedCtx, combinedCancel := context.WithCancel(newCtx)
				go func() {
					select {
					//case <-newCtx2.Done():
					case <-combinedCtx.Done():
					}
					combinedCancel()
				}()
				err = scheduler.Schedule(combinedCtx, utils.ConvertDBTaskToSchedulerTask(newTask))
				if err != nil {
					if err == context.Canceled {
						tracer.Info("task schedule has been canceled.")
						return
					}
					tracer.Warnf("schedule task failed: %s", err)
				}
			}(task)
		}
	}
	return nil
}
