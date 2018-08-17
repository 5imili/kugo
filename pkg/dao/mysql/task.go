package mysql

import (
	"context"

	"github.com/leopoldxx/go-utils/trace"
)

func (m *mysql) CreateTask(ctx context.Context) {
	tracer := trace.GetTraceFromContext(ctx)
	tracer.Info("CreateTask")
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
