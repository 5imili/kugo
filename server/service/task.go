package service

import (
	"context"

	"github.com/5imili/kugo/pkg/dao/mysql/types"
	"github.com/5imili/kugo/pkg/enum"
	"github.com/leopoldxx/go-utils/trace"
)

func (s *service) ListTask(ctx context.Context) error {
	tracer := trace.GetTraceFromContext(ctx)
	tracer.Info("call service listTask by service")
	return nil
}

func (s *service) CreateTask(ctx context.Context, namespace string, info *Task) (*types.Task, error) {
	tracer := trace.GetTraceFromContext(ctx)
	tracer.Info("call service CreateTask by service")
	task := &types.Task{
		Resource:  info.Resource,
		Type:      string(enum.TaskCreate),
		NameSpace: namespace,
	}
	taskID, err := s.opt.DB.CreateTask(ctx, task)
	if err != nil {
		tracer.Errorf("failed insert task")
		return nil, err
	}
	task.ID = taskID
	return task, nil
}

func (s *service) GetTask(ctx context.Context) error {
	tracer := trace.GetTraceFromContext(ctx)
	tracer.Info("call service GetTask by service")
	return nil
}

func (s *service) DeleteTask(ctx context.Context) error {
	tracer := trace.GetTraceFromContext(ctx)
	tracer.Info("call service DeleteTask by service")
	return nil
}
