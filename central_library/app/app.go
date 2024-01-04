package app

import (
	"central_library/config"
	"github.com/gin-gonic/gin"
	"log"
)

type App struct {
	Config config.Config
	Router *gin.Engine
}

func NewApp() (*App, error) {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	app := &App{
		Config: config,
	}

	err = app.CreateRoutersAndSetRoutes()
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	return app, nil
}
