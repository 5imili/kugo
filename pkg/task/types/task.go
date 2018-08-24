package types

import (
	"time"

	"github.com/5imili/kugo/pkg/dao"
)

// InitConfigs for backend scheduler
type InitConfigs struct {
	Dao dao.Storage
}

// Common contains some common properties of a task
type Common struct {
	ID        int64
	Namespace string
	// CategoryType     string
	// CategoryID       int64
	Resource         string
	Type             string
	Pause            bool
	SkipPause        bool
	Close            bool
	IsClosedManually bool
	LastUpdateTime   time.Time
	//State          State
}

// Task is a generic task type, can get specific task from Raw field
type Task struct {
	Common
	Spec   string
	Status string
}
