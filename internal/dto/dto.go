package dto

type CreateDogRequest struct {
	Breed    string `json:"breed" validate:"required"`
	SubBreed string `json:"sub_breed" validate:"required"`
}

type UpdateDogRequest struct {
	Id       string `json:"id" validate:"required,uuid4"`
	Breed    string `json:"breed" validate:"required"`
	SubBreed string `json:"sub_breed" validate:"required"`
}

type DeleteDogRequest struct {
	Id string `json:"id" validate:"required,uuid4"`
}

type CreateDogResponse struct {
	Id string `json:"id"`
}
