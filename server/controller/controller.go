package controller

import (
	"github.com/5imili/kugo/pkg/dao"
	"github.com/5imili/kugo/server/service"
	"github.com/gorilla/mux"
)

//Controller xxx
type Controller interface {
	Register(router *mux.Router)
}

// Options of controller
type Options struct {
	Service service.Operation // 进行对应操作
	DB      dao.Storage
}
