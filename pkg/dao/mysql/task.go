package mysql

import (
	"context"
	"database/sql"

	"github.com/5imili/kugo/pkg/dao/mysql/types"
	"github.com/leopoldxx/go-utils/trace"
)

func (m *mysql) CreateTask(ctx context.Context, task *types.Task) (int64, error) {
	const (
		sqlTpl = `
INSERT INTO task (
	namespace,
	resource,
	task_type,
	spec,
	status,
	op_user,
	create_time)
VALUES (:namespace,
	:resource,
	:task_type,
	:spec,
	:status,
	:op_user,
	NOW());`
	)
	tracer := trace.GetTraceFromContext(ctx)

	var (
		res sql.Result
		err error
	)
	if m.db != nil {
		res, err = m.db.Exec(sqlTpl, task.NameSpace, task.Resource, task.Type, task.Spec, task.Status, task.OpUser)
		if err != nil {
			tracer.Errorf("failed to insert task: %s", err)
			return 0, err
		}
	}

	tracer.Info("insert task successfully")
	lastID, err := res.LastInsertId()
	if err != nil {
		tracer.Errorf("failed to get lastid of the task: %s", err)
		return 0, err
	}
	return lastID, err
}

func (m *mysql) ListTask(ctx context.Context) {
	tracer := trace.GetTraceFromContext(ctx)
	tracer.Info("CreateTask")
}

func (m *mysql) GetTask(ctx context.Context) {
	tracer := trace.GetTraceFromContext(ctx)
	tracer.Info("CreateTask")
}

func (m *mysql) DeleteTask(ctx context.Context) {
	tracer := trace.GetTraceFromContext(ctx)
	tracer.Info("CreateTask")
}
