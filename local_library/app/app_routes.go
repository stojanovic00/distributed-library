package app

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"local_library/core/domain"
	"local_library/handler"
	"local_library/persistence"
	"log"
)

func (a *App) CreateRoutersAndSetRoutes() error {
	//DEPENDENCIES
	userHandler := handler.NewUserHandler(&a.Config)
	pgClient := a.initPGClient()
	issuingRepo := persistence.NewIssuingRepoPg(pgClient)
	issuingHandler := handler.NewIssuingHandler(&a.Config, issuingRepo)

	// ROUTES
	router := gin.Default()

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Endpoint doesn't exist"})
	})

	userGroup := router.Group("/user")
	userGroup.POST("", userHandler.Register)

	bookGroup := router.Group("/book")
	bookGroup.POST("/issue", issuingHandler.RecordBookIssue)
	bookGroup.PUT(":isbn/return", issuingHandler.RecordBookReturn)

	a.Router = router
	return nil
}

func (a *App) initPGClient() *gorm.DB {
	client, err := persistence.GetPostgresClient(
		a.Config.DbHost, a.Config.DbUser,
		a.Config.DbPass, a.Config.DbName,
		a.Config.DbPort)
	if err != nil {
		log.Fatal(err)
	}
	err = client.AutoMigrate(
		&domain.IssuingRecord{},
	)
	if err != nil {
		log.Fatal(err)
	}

	return client
}
