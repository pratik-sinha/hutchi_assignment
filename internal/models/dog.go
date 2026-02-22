package models

import "time"

type Dog struct {
	ID        string    `db:"id"`
	Breed     string    `db:"breed"`
	SubBreed  string    `db:"sub_breed"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
