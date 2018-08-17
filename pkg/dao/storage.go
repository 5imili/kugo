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
}
