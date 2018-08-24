package taskscheduler

import (
	"context"
	"errors"

	"github.com/5imili/kugo/pkg/dao"
	"github.com/5imili/kugo/pkg/task/scheduler"
	schedtypes "github.com/5imili/kugo/pkg/task/types"
	"github.com/leopoldxx/go-utils/trace"
)

const (
	maxRetryTimes = 3
)

type taskScheduler struct {
	dao dao.Storage
}

var (
	task = taskScheduler{}
)

// Scheduler return the global task scheduler
func Scheduler() scheduler.Scheduler {
	return &task
}

func (sched *taskScheduler) GetName() string {
	return string("task")
}
func (sched *taskScheduler) Init(cfg schedtypes.InitConfigs) error {
	if sched == nil {
		return errors.New("sched is nil")
	}
	sched.dao = cfg.Dao
	return nil
}

func (sched *taskScheduler) Schedule(ctx context.Context, task *schedtypes.Task) error {
	tracer := trace.GetTraceFromContext(ctx)
	tracer.Infof("get an podvmgroup task, Common:%+v, Spec:%+v, Status:%+v",
		task.Common, task.Spec, task.Status)

	return nil
}
