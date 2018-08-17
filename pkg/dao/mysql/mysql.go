package mysql

import (
	"github.com/5imili/kugo/pkg/dao"
	"github.com/jmoiron/sqlx"
)

// Options stores mysql needed parameters
type Options struct {
	DbConnStr   string
	DbMaxConns  int
	DBIdleConns int
	// obsoleted
	DbName string
}

type mysql struct {
	opt *Options
	db  *sqlx.DB
}

//New create dao interface
func New(opt *Options) dao.Storage {
	s := mysql{
		opt: opt,
	}
	return &s
}
