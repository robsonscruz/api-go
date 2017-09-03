package controllers

import (
	"net/http"
	"strconv"

	"github.com/gernest/utron/controller"
	"github.com/gorilla/schema"
	"github.com/robsonscruz/api-go/models"
)

var decoder = schema.NewDecoder()

//HistoryDeploy is a controller for HistoryDeploy list
type HistoryDeploy struct {
	controller.BaseController
	Routes []string
}

//Home renders a todo list
func (t *HistoryDeploy) Home() {
	histories := []*models.HistoryDeploy{}
	t.Ctx.DB.Order("created_at desc").Find(&histories)
	t.Ctx.Data["List"] = histories
	t.Ctx.Template = "deploy"
	t.HTML(http.StatusOK)
}

//NewTodo returns a new  todo list controller
func NewHistoryDeploy() controller.Controller {
	return &HistoryDeploy{
		Routes: []string{
			"get;/deploy;Home",
		},
	}
}