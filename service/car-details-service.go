package service

import (
	"golang-mux-api/entity"
)

type CarDetailsService interface {
	GetDetails() entity.CarDetails
}

type service struct{}

func NewCarDetailsService() CarDetailsService {
	return &service{}
}

func (*service) GetDetails() entity.CarDetails {
	return entity.CarDetails{
		ID:        0,
		Brand:     "",
		Model:     "",
		Year:      0,
		FirstName: "",
		LastName:  "",
		Email:     "",
	}
}
