package main

import (
	"log"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/kconde2/vote-app/api/controllers"
	"github.com/kconde2/vote-app/api/db"
	"github.com/kconde2/vote-app/api/middleware"
	"github.com/itsjamie/gin-cors"
)

func main() {
	log.Println("Starting server...")

	db.Initialize()
	db.CreateSystemAdmin()


	r := setupRouter()
	r.Run(":8080")
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.Middleware(cors.Config{
		Origins:        "*",
		Methods:        "GET, PUT, POST, DELETE",
		RequestHeaders: "Origin, Authorization, Content-Type",
		ExposedHeaders: "",
		MaxAge: 50 * time.Second,
		Credentials: true,
		ValidateHeaders: false,
	}))


	route := r.Group("/")

	// Manage login (auth + generate JWT)
	authMiddleware, err := middleware.AuthMiddleware()
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	route.POST("/login", authMiddleware.LoginHandler)

	// Manage protected routes
	route.Use(authMiddleware.MiddlewareFunc())
	{
		route.GET("/user/:email", controllers.RetrieveUserByEmail)
		users := route.Group("/users")
		{
			users.GET(":uuid", controllers.RetrieveUser)
			users.GET("", controllers.GetUsers)
			users.POST("", controllers.CreateUser)
			users.PUT(":uuid", controllers.UpdateUser)
			users.DELETE(":uuid", controllers.DeleteUser)
		}

		votes := route.Group("/votes")
		{
			votes.GET("", controllers.GetVotes)
			votes.POST("", controllers.CreateVote)
			votes.GET(":uuid", controllers.RetrieveVote)
			votes.PUT(":uuid", controllers.UpdateVote)
		}
	}
	return r
}
