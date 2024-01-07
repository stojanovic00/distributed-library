package app

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"local_library/core/domain"
	"local_library/handler"
	"local_library/persistence"
	"log"
	"time"
)

func (a *App) CreateRoutersAndSetRoutes() error {
	//DEPENDENCIES
	userHandler := handler.NewUserHandler(&a.Config)
	pgClient := a.initPGClient()
	issuingRepo := persistence.NewIssuingRepoPg(pgClient)
	issuingHandler := handler.NewIssuingHandler(&a.Config, issuingRepo)

	// ROUTES
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Endpoint doesn't exist"})
	})

	cityGroup := router.Group(a.Config.CityPath)

	userGroup := cityGroup.Group("/user")
	userGroup.POST("", userHandler.Register)

	bookGroup := cityGroup.Group("/book")
	bookGroup.POST("/issue", issuingHandler.RecordBookIssue)
	bookGroup.PUT(":isbn/return", issuingHandler.RecordBookReturn)

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
