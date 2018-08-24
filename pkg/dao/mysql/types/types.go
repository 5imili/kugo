package types

import "time"

// Task represents a user defined task, including task types, status etc.
// It's not the bottom storage schema, it is a higher level view.
type Task struct {
	ID                int64     `db:"id"`
	NameSpace         string    `db:"namespace"`
	Resource          string    `db:"resource"`
	Type              string    `db:"task_type"`
	Spec              string    `db:"spec"`
	Status            string    `db:"status"`
	IsCanceled        bool      `db:"is_canceled"`
	IsPaused          bool      `db:"is_paused"`
	IsSkipPaused      bool      `db:"is_skip_paused"`
	IsUrgentSkipped   bool      `db:"is_urgent_skipped"`
	UrgentSkipComment string    `db:"urgent_skip_comment"`
	IsClosed          bool      `db:"is_closed"`
	IsClosedManually  bool      `db:"is_closed_manually"`
	OpUser            string    `db:"op_user"`
	CreateTime        time.Time `db:"create_time"`
	LastUpdateTime    time.Time `db:"last_update_time"`
}

type Field string

const (
	FieldID                = Field("id")
	FieldUUID              = Field("uuid")
	FieldStatus            = Field("status")
	FieldIsSkipped         = Field("is_skipped")
	FieldIsWithdrawn       = Field("is_withdrawn")
	FieldIsApproved        = Field("is_approved")
	FieldIsExecuted        = Field("is_executed")
	FieldIsClosed          = Field("is_closed")
	FieldIsClosedManually  = Field("is_closed_manually")
	FieldIsPaused          = Field("is_paused")
	FieldIsSkipPaused      = Field("is_skip_paused")
	FieldIsUrgentSkipped   = Field("is_urgent_skipped")
	FieldUrgentSkipComment = Field("urgent_skip_comment")
)

type Value interface{}
