package utils

import "github.com/go-playground/validator/v10"

type validatorStruct struct {
	v *validator.Validate
}

type ValidatorInterface interface {
	Struct(interface{}) error
}

func NewValidator() ValidatorInterface {
	return &validatorStruct{
		v: validator.New(),
	}
}

func (vs *validatorStruct) Struct(value interface{}) error {
	err := vs.v.Struct(value)
	if err != nil {
		return err
	}
	return nil
}
