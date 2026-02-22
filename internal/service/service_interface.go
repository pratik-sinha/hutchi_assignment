package service

import (
	"assignment/internal/dto"
	"assignment/internal/models"
)

type Service interface {
	GetDogs() ([]models.Dog, error)
	CreateDog(req dto.CreateDogRequest) (*string, error)
	UpdateDog(req dto.UpdateDogRequest) error
	DeleteDog(req dto.DeleteDogRequest) error
}
