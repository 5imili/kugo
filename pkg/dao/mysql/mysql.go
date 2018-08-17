package mysql

import (
	"fmt"

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
	db, err := sqlx.Open("mysql", opt.DbConnStr)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	db.SetMaxOpenConns(opt.DbMaxConns)
	db.SetMaxIdleConns(opt.DBIdleConns)
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	s := mysql{
		opt: opt,
		db:  db,
	}
	return &s
}
