package taskscheduler

import (
	"context"
	"errors"

	"github.com/5imili/kugo/pkg/dao"
	"github.com/5imili/kugo/pkg/enum"
	"github.com/5imili/kugo/pkg/task/scheduler"
	tasktypes "github.com/5imili/kugo/pkg/task/taskscheduler/types"
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
	t, err := convertSchedTaskToTask(task)
	if err != nil {
		tracer.Errorf("convert task failed: %s, %+v", err, *task)
		return err
	}
	tracer.Infof("get a task, Common:%+v, Spec:%+v, Status:%+v",
		t.Common, t.Spec, t.Status)

	if len(t.Status.State) == 0 {
		// init status
		t.Status.State = enum.TaskPending
		t.Status.TryTimes = 0
	}
	tracer.Info(ctx, string(t.Status.State))
	switch t.Status.State {
	case enum.TaskPending:
		tracer.Info("current is ", enum.TaskPending)
		return sched.taskPending(ctx, t)
	case enum.TaskDoing:
		tracer.Info("current is ", enum.TaskDoing)
		return sched.taskDoing(ctx, t)
	case enum.TaskDone:
		tracer.Info("current is ", enum.TaskDone)
		return sched.taskDone(ctx, t)
	default:
		tracer.Errorf("unknown status %s of task %v", string(t.Status.State), *task)
		return errors.New("unknown state of the task")
	}
}

func (sched *taskScheduler) taskPending(ctx context.Context, t *tasktypes.Task) error {
	t.Status.State = enum.TaskDoing
	return sched.updateTaskStatus(ctx, t)
}

func (sched *taskScheduler) taskDoing(ctx context.Context, t *tasktypes.Task) error {
	t.Status.State = enum.TaskDone
	return sched.updateTaskStatus(ctx, t)
}

func (sched *taskScheduler) taskDone(ctx context.Context, t *tasktypes.Task) error {
	t.Status.State = enum.TaskDoing
	t.Common.Close = true
	return sched.updateTaskStatus(ctx, t)
}

func (sched *taskScheduler) updateTaskStatus(ctx context.Context, task *tasktypes.Task) error {
	tracer := trace.GetTraceFromContext(ctx)
	dbtask := convertTaskToDBTask(task)
	err := sched.dao.UpdateTask(ctx, dbtask)
	if err != nil {
		tracer.Error(err)
		return err
	}
	return nil
}
