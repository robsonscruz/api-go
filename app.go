package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gernest/utron"
	c "github.com/robsonscruz/api-go/controllers"
	"github.com/robsonscruz/api-go/models"
)

func main() {

	// Start the MVC App
	app, err := utron.NewMVC()
	if err != nil {
		log.Fatal(err)
	}

	// Register Models
	app.Model.Register(&models.Todo{})
	app.Model.Register(&models.HistoryDeploy{})

	// CReate Models tables if they dont exist yet
	app.Model.AutoMigrateAll()

	// Register Controller
	app.AddController(c.NewTodo)
	app.AddController(c.NewHistoryDeploy)

	// Start the server
	port := fmt.Sprintf(":%d", app.Config.Port)
	app.Log.Info("staring server on port", port)
	log.Fatal(http.ListenAndServe(port, app))
}