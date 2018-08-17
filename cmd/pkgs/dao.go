package pkgs

import (
	"sync"

	"github.com/5imili/kugo/pkg/dao"
	"github.com/5imili/kugo/pkg/dao/mysql"
	"github.com/spf13/viper"
)

const (
	DaoConnection = "dao.connection"
	DaoMaxConns   = "dao.maxConns"
	DaoIdleConns  = "dao.idleConns"
)

var (
	mysqlDao  dao.Storage
	mysqlOnce sync.Once
)

func init() {
	initDaoDefault()
}

func initDaoDefault() {
	viper.SetDefault(DaoConnection, "test:test@/kugo?charset=utf8&parseTime=true&loc=Asia%2FShanghai")
	viper.SetDefault(DaoMaxConns, 100)
	viper.SetDefault(DaoIdleConns, 50)
}

//GetDao xxx
func GetDao() dao.Storage {
	mysqlOnce.Do(func() {
		mysqlDao = mysql.New(&mysql.Options{
			DbConnStr:   viper.GetString(DaoConnection),
			DbMaxConns:  viper.GetInt(DaoMaxConns),
			DBIdleConns: viper.GetInt(DaoIdleConns),
		})
		if mysqlDao == nil {
			panic("connect mysql failed")
		}
	})
	return mysqlDao
}
