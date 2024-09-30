package api

import (
	"gin-gorm-crud/internal/api/handlers"
	"gin-gorm-crud/internal/middleware"
	"gin-gorm-crud/internal/repository"
	"gin-gorm-crud/internal/service"
	"gin-gorm-crud/internal/ws"
	"time"

	"gorm.io/gorm"

	docs "gin-gorm-crud/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000"
		},
		MaxAge: 12 * time.Hour,
	}))

	v1 := router.Group("/api/v1")
	authRoutes := v1.Group("/auth")
	{
		authRoutes.POST("/login", userHandler.Login)
		authRoutes.POST("/signup", userHandler.CreateUser)
	}

	userRoutes := v1.Group("/users")
	{
		userRoutes.GET("/email/:email", middleware.AuthMiddleware(), userHandler.GetUserByEmail)
	}

	// WebSocket Routes
	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)
	go hub.Run()

	wsRoutes := v1.Group("/ws")
	{
		wsRoutes.POST("/createRoom", wsHandler.CreateRoom)
		wsRoutes.GET("/joinRoom/:roomId", wsHandler.JoinRoom)
		wsRoutes.GET("/getRooms", wsHandler.GetRooms)
		wsRoutes.GET("/getClients/:roomId", wsHandler.GetClients)
	}

}
