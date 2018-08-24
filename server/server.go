package server

import (
	"net/http"
	"time"

	"github.com/5imili/kugo/server/controller"
	"github.com/5imili/kugo/server/controller/task"
	"github.com/5imili/kugo/server/service"
	"github.com/facebookgo/httpdown"
	"github.com/gorilla/mux"
)

//Server xxx
type Server interface {
	http.Handler
	ListenAndServe() error
}

// Options contains the parameters needed by a api server
type Options struct {
	ListenAddr      string
	EnableHTTPDebug bool
	CtrlOpts        *controller.Options
	//StaticPath string
}

type server struct {
	opt    Options
	router *mux.Router
}

// New will create a new api server
func New(opt Options) Server {
	router := mux.NewRouter()
	// add controllers
	v2Router := router.PathPrefix("/kugo/api/v2").Subrouter()
	//register sub routers
	task.New(opt.CtrlOpts).Register(v2Router)
	opt.CtrlOpts.Service = service.New(&service.Options{
		DB: opt.CtrlOpts.DB,
	})
	return &server{
		opt:    opt,
		router: v2Router,
	}
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) ListenAndServe() error {
	httpServer := &http.Server{
		Addr:    s.opt.ListenAddr,
		Handler: s.router,
	}
	hd := &httpdown.HTTP{
		StopTimeout: time.Second,
		KillTimeout: time.Second,
	}
	return httpdown.ListenAndServe(httpServer, hd)
}
