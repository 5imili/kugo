package task

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/5imili/kugo/server/controller"
	"github.com/5imili/kugo/server/utils"
	"github.com/gorilla/mux"
	"github.com/leopoldxx/go-utils/middleware"
	"github.com/leopoldxx/go-utils/trace"
)

type task struct {
	opt *controller.Options
}

// New will create an app controller
func New(opt *controller.Options) controller.Controller {
	return &task{opt: opt}
}

func (t *task) Register(router *mux.Router) {
	subrouter := router.PathPrefix("/namespaces/{namespace}").Subrouter()
	//subrouter.Use(utils.AuthenticateMW)
	subrouter.Use(utils.LoggingMiddleware)
	subrouter.Methods("GET").Path("/tasks").HandlerFunc(
		middleware.RecoverWithTrace("listTask").HandlerFunc(
			utils.AuthenticateMW().HandlerFunc(t.listTask),
		),
	)

	subrouter.Methods("POST").Path("/tasks").HandlerFunc(
		middleware.Chain(
			middleware.RecoverWithTrace("createTask"),
		).HandlerFunc(t.createTask))
	subrouter.Methods("GET").Path("/tasks/{task}").HandlerFunc(
		middleware.Chain(
			middleware.RecoverWithTrace("getTask"),
		).HandlerFunc(t.getTask))
	subrouter.Methods("DELETE").Path("/tasks").HandlerFunc(
		middleware.Chain(
			middleware.RecoverWithTrace("deleteTask"),
		).HandlerFunc(t.deleteTask))
}

func (t *task) listTask(w http.ResponseWriter, r *http.Request) {
	tracer := trace.GetTraceFromRequest(r)
	tracer.Info("list getTasks")
	t.opt.Service.ListTask(r.Context())
	fmt.Fprintln(w, "hello boy")
}

func (t *task) createTask(w http.ResponseWriter, r *http.Request) {
	tracer := trace.GetTraceFromRequest(r)
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		tracer.Error(err)
		utils.CommReply(w, r, http.StatusBadRequest, err.Error())
		return
	}
	tracer.Info(string(data))
	tracer.Info("createTask")
	t.opt.Service.CreateTask(r.Context())
	utils.CommReply(w, r, http.StatusOK, "success")
}

func (t *task) deleteTask(w http.ResponseWriter, r *http.Request) {
	tracer := trace.GetTraceFromRequest(r)
	tracer.Info("deleteTask")
	t.opt.Service.DeleteTask(r.Context())
	utils.CommReply(w, r, http.StatusOK, "success")
}

func (t *task) getTask(w http.ResponseWriter, r *http.Request) {
	tracer := trace.GetTraceFromRequest(r)
	tracer.Info("getTask")
	t.opt.Service.GetTask(r.Context())
	utils.CommReply(w, r, http.StatusOK, "success")
}
