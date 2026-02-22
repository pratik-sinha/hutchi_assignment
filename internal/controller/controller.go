package controller

import (
	"assignment/internal/dto"
	"assignment/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type controller struct {
	s service.Service
}

func NewController(s service.Service) Controller {
	return &controller{
		s: s,
	}
}

func (cont *controller) GetDogs(c *gin.Context) {
	res, err := cont.s.GetDogs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, gin.H{"status": true, "body": res})
}

func (cont *controller) CreateDog(c *gin.Context) {
	var req dto.CreateDogRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	Id, err := cont.s.CreateDog(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, gin.H{"status": true, "body": dto.CreateDogResponse{Id: *Id}})
}

func (cont *controller) UpdateDog(c *gin.Context) {
	var req dto.UpdateDogRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	err = cont.s.UpdateDog(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, gin.H{"status": true})
}

func (cont *controller) DeleteDog(c *gin.Context) {
	var req dto.DeleteDogRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	err = cont.s.DeleteDog(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, gin.H{"status": true})
}
