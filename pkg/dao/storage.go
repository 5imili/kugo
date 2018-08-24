package dao

import (
	"context"

	"github.com/5imili/kugo/pkg/dao/mysql/types"
)

//Storage xxx
type Storage interface {
	ListTask(ctx context.Context)
	CreateTask(ctx context.Context, task *types.Task) (int64, error)
	GetTask(ctx context.Context)
	DeleteTask(ctx context.Context)
	ListOpenTasks(ctx context.Context) ([]types.Task, error)
	GetOpenTaskByTaskID(ctx context.Context, id int64) (*types.Task, error)
	UpdateTask(ctx context.Context, task *types.Task) error
}
