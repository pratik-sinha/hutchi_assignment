package service

import (
	"assignment/internal/dto"
	"assignment/internal/models"
	"assignment/internal/repository"
	"assignment/pkg/utils"
)

type service struct {
	r repository.Repository
	v utils.ValidatorInterface
}

func NewService(r repository.Repository, v utils.ValidatorInterface) Service {
	return &service{
		r: r,
		v: v}

}

func (s *service) GetDogs() ([]models.Dog, error) {
	return s.r.GetDogs()
}

func (s *service) CreateDog(req dto.CreateDogRequest) (*string, error) {
	err := s.v.Struct(req)
	if err != nil {
		return nil, err
	}
	return s.r.CreateDog(models.Dog{Breed: req.Breed, SubBreed: req.SubBreed})
}

func (s *service) UpdateDog(req dto.UpdateDogRequest) error {
	err := s.v.Struct(req)
	if err != nil {
		return err
	}
	return s.r.UpdateDog(models.Dog{ID: req.Id, Breed: req.Breed, SubBreed: req.SubBreed})
}

func (s *service) DeleteDog(req dto.DeleteDogRequest) error {
	err := s.v.Struct(req)
	if err != nil {
		return err
	}
	return s.r.DeleteDog(req.Id)
}
