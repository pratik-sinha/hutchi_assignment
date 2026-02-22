package routes

import (
	"assignment/internal/controller"

	"github.com/gin-gonic/gin"
)

func DogRoutes(route *gin.RouterGroup, controller controller.Controller) {
	route.GET("/getdogs", controller.GetDogs)
	route.POST("/createdog", controller.CreateDog)
	route.PUT("/updatedog", controller.UpdateDog)
	route.DELETE("/deletedog", controller.DeleteDog)

}
