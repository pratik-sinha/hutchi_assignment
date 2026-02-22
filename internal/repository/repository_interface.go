package repository

import "assignment/internal/models"

type Repository interface {
	CreateDog(models.Dog) (*string, error)
	GetDogs() ([]models.Dog, error)
	DeleteDog(Id string) error
	UpdateDog(dog models.Dog) error
}
