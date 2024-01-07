package app

import (
	"central_library/core/domain"
	"central_library/handler"
	"central_library/persistence"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"time"
)

func (a *App) CreateRoutersAndSetRoutes() error {
	//DEPENDENCIES
	pgClient := a.initPGClient()
	userRepo := persistence.NewUserRepoPg(pgClient)
	userHandler := handler.NewUserHandler(userRepo)
	// ROUTES
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Endpoint doesn't exist"})
	})

	userGroup := router.Group("/user")
	userGroup.POST("", userHandler.Register)
	userGroup.POST("/:id/record-issue", userHandler.RecordBookIssue)
	userGroup.POST("/:id/record-return", userHandler.RecordBookReturn)

	a.Router = router
	return nil
}

func (a *App) initPGClient() *gorm.DB {
	var (
		client *gorm.DB
		err    error
	)

	for attempt := 1; attempt <= 30; attempt++ {
		client, err = persistence.GetPostgresClient(
			a.Config.DbHost, a.Config.DbUser,
			a.Config.DbPass, a.Config.DbName,
			a.Config.DbPort)
		if err == nil {
			break
		}
		log.Printf("[attempt %d/30] Failed to connect to db, retrying in 3 seconds...", attempt)
		time.Sleep(3 * time.Second)
	}

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Connected to database!")

	err = client.AutoMigrate(
		&domain.User{},
	)
	if err != nil {
		log.Fatal(err)
	}

	return client
}
