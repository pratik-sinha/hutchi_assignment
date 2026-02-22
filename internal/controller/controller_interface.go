package controller

import "github.com/gin-gonic/gin"

type Controller interface {
	GetDogs(c *gin.Context)
	CreateDog(c *gin.Context)
	UpdateDog(c *gin.Context)
	DeleteDog(c *gin.Context)
}
