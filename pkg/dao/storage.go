package dao

import "context"

//Storage xxx
type Storage interface {
	ListTask(ctx context.Context)
	CreateTask(ctx context.Context)
	GetTask(ctx context.Context)
	DeleteTask(ctx context.Context)
}
