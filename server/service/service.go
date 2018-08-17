package service

import (
	"context"

	"github.com/5imili/kugo/pkg/dao"
	"github.com/5imili/kugo/pkg/dao/mysql/types"
)

//Operation xxx
type Operation interface {
	ListTask(ctx context.Context) error
	CreateTask(ctx context.Context, namespace string, info *Task) (*types.Task, error)
	GetTask(ctx context.Context) error
	DeleteTask(ctx context.Context) error
}

// Options contains other pkgs used for console operations
// like db or kubernetes etc.
type Options struct {
	DB dao.Storage
}

type service struct {
	opt *Options
}

// New will create a console implementaion of Operation
func New(opt *Options) Operation {
	return &service{
		opt: opt,
	}
}
