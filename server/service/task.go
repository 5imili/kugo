package service

import (
	"context"

	"github.com/leopoldxx/go-utils/trace"
)

func (s *service) ListTask(ctx context.Context) error {
	tracer := trace.GetTraceFromContext(ctx)
	tracer.Info("call listTask by service")
	return nil
}

func (s *service) CreateTask(ctx context.Context) error {
	tracer := trace.GetTraceFromContext(ctx)
	tracer.Info("call CreateTask by service")
	return nil
}

func (s *service) GetTask(ctx context.Context) error {
	tracer := trace.GetTraceFromContext(ctx)
	tracer.Info("call GetTask by service")
	return nil
}

func (s *service) DeleteTask(ctx context.Context) error {
	tracer := trace.GetTraceFromContext(ctx)
	tracer.Info("call DeleteTask by service")
	return nil
}
