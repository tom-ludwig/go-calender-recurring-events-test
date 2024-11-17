package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tom-ludwig/calender-events-test/handlers"
	"github.com/tom-ludwig/calender-events-test/services"
)

func InitilizeUserRoutes(r *gin.RouterGroup, db_pool *pgxpool.Pool) {
	userService := services.NewUserService(db_pool)
	userHandler := handlers.NewUserHandler(userService)

	userGroup := r.Group("/user")

	// TODO: Check if only certain routes can use certain middleware i.e some routes with user shouldn't need authentication
	// userGroup.Use(userHandler.CheckRequiredRole)

	userGroup.GET("/", userHandler.GetUser).POST("/", userHandler.CreateUser).PUT("/", userHandler.UpdateUser).DELETE("/", userHandler.DeleteUser)
	userGroup.POST("/login", userHandler.LogIn)
	userGroup.GET("/all", userHandler.GetAllUsers)
}
