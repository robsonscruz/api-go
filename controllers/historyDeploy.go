package controllers

import (
	"net/http"
	"github.com/gernest/utron/controller"
	"github.com/robsonscruz/api-go/models"
	"time"
	"github.com/gorilla/schema"
)

var decoder = schema.NewDecoder()

//HistoryDeploy is a controller for HistoryDeploy list
type HistoryDeploy struct {
	controller.BaseController
	Routes []string
}

func ReponseSuccess() string {
	return `{"response":"success", "message":"data successfully registered"}`
}

//Home renders a todo list
func (t *HistoryDeploy) Home() {
	histories := []*models.HistoryDeploy{}
	t.Ctx.DB.Order("created_at desc").Limit(10).Find(&histories)
	t.Ctx.Data["List"] = histories
	t.Ctx.Data["CurrentDate"] = time.Now()
	t.Ctx.Template = "deploy"
	t.HTML(http.StatusOK)
}

//Home renders a todo list
func (t *HistoryDeploy) Download() {
	histories := []*models.HistoryDeploy{}
	t.Ctx.DB.Order("created_at desc").Find(&histories)
	t.Ctx.Data["List"] = histories
	t.Ctx.Template = "export"
	t.HTML(http.StatusOK)
}

//Create creates a todo  item
func (t *HistoryDeploy) Create() {
	history := &models.HistoryDeploy{}
	req := t.Ctx.Request()
	_ = req.ParseForm()
	if err := decoder.Decode(history, req.PostForm); err != nil {
		t.Ctx.Data["Message"] = err.Error()
		t.Ctx.Template = "error"
		t.HTML(http.StatusInternalServerError)
		return
	}

	t.Ctx.DB.Create(history)
	t.Ctx.Data["Message"] = ReponseSuccess()
	t.Ctx.Template = "json"
	t.HTML(http.StatusOK)
}

//NewTodo returns a new  todo list controller
func NewHistoryDeploy() controller.Controller {
	return &HistoryDeploy{
		Routes: []string{
			"get;/;Home",
			"post;/create;Create",
			"get;/download;Download",
		},
	}
}