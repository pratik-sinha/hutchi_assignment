package repository

import (
	"assignment/internal/models"
	"assignment/pkg/utils"
	"errors"

	"github.com/jmoiron/sqlx"
)

type respository struct {
	dbConn *sqlx.DB
}

func NewRepository(dbConn *sqlx.DB) Repository {
	return &respository{
		dbConn: dbConn,
	}
}

func (r *respository) CreateDog(dog models.Dog) (*string, error) {
	var id string
	query := "INSERT INTO DOGS (breed,sub_breed) VALUES ($1,$2) RETURNING id"
	err := r.dbConn.Get(&id, query, dog.Breed, dog.SubBreed)
	if err != nil {
		return nil, err
	}
	return &id, nil
}

func (r *respository) GetDogs() ([]models.Dog, error) {
	var res []models.Dog
	err := r.dbConn.Select(&res, "select * from dogs")
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *respository) UpdateDog(dog models.Dog) error {
	query := `UPDATE dogs SET breed=$1, sub_breed=$2, updated_at=$3 WHERE id = $4`
	res, err := r.dbConn.Exec(query, dog.Breed, dog.SubBreed, utils.GetUTCTimeStamp(), dog.ID)
	if err != nil {
		return err
	}
	rows, _ := res.RowsAffected()

	if rows != 1 {
		return errors.New("No dog breed found with id " + dog.ID)
	}

	return nil
}

func (r *respository) DeleteDog(Id string) error {
	query := `DELETE FROM dogs WHERE id = $1`
	res, err := r.dbConn.Exec(query, Id)
	if err != nil {
		return err
	}
	rows, _ := res.RowsAffected()

	if rows != 1 {
		return errors.New("No dog breed found with id " + Id)
	}

	return nil
}
