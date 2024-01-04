package app

import (
	"central_library/core/domain"
	"central_library/handler"
	"central_library/persistence"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

func (a *App) CreateRoutersAndSetRoutes() error {
	//DEPENDENCIES
	pgClient := a.initPGClient()
	userRepo := persistence.NewUserRepoPg(pgClient)
	userHandler := handler.NewUserHandler(userRepo)

	// ROUTES
	router := gin.Default()

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Endpoint doesn't exist"})
	})

	userGroup := router.Group("/user")
	userGroup.POST("", userHandler.Register)
	userGroup.PUT("/:id/record-issue", userHandler.RecordBookIssue)
	userGroup.PUT("/:id/record-return", userHandler.RecordBookReturn)

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
		&domain.User{},
	)
	if err != nil {
		log.Fatal(err)
	}

	return client
}
