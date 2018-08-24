package pkgs

import (
	"context"
	"fmt"
	"sync"

	"github.com/5imili/kugo/pkg/dao"
	"github.com/5imili/kugo/pkg/task/scheduler"
	"github.com/5imili/kugo/pkg/task/taskscheduler"
	"github.com/leopoldxx/go-utils/trace"
)

var (
	manager     *scheduler.Manager
	managerOnce sync.Once
)

//GetScheduler xxx
func GetScheduler(sto dao.Storage) *scheduler.Manager {
	managerOnce.Do(func() {
		var err error
		ctx := trace.WithTraceForContext(context.TODO(), "task-scheduler")
		manager, err = scheduler.NewManager(ctx, sto)
		if err != nil {
			panic(fmt.Sprintf("create task scheduler manager failed: %s", err))
		}
		if err = manager.InitSchedulers(taskscheduler.Scheduler()); err != nil {
			panic(fmt.Sprintf("init scheduler failed: %s", err))
		}
	})
	return manager
}
