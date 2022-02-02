package main

import (
	"backend/controller"
	_ "backend/model"

	"github.com/gin-gonic/gin"
)

func RegisterRoute(r *gin.Engine) *gin.Engine {
	customerRoutes := r.Group("/customer")
	customerRoutes.POST("/register", controller.Register)
	customerRoutes.POST("/login", controller.Login)
	customerRoutes.GET("/delete", controller.Delete)
	return r
}

func RatingRoute(r *gin.Engine) *gin.Engine {
	customerRoutes := r.Group("/rating")
	customerRoutes.POST("/create", controller.CreateRating)
	customerRoutes.GET("/get", controller.GetRating)
	return r
}

func CuisineRoute(r *gin.Engine) *gin.Engine {
	modelRoutes := r.Group("/cuisine")
	modelRoutes.POST("/create", controller.Create)
	// modelRoutes.POST("/delete", controller.Delete)
	// modelRoutes.POST("/update", controller.Update)
	// modelRoutes.GET("/read", controller.Read)

	return r
}
