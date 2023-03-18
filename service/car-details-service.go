package service

import (
	"encoding/json"
	"fmt"
	"golang-mux-api/entity"
	"net/http"
)

type CarDetailsService interface {
	GetDetails() entity.CarDetails
}

var (
	carService       CarService   = NewCarService()
	ownerService     OwnerService = NewOwnerService()
	carDataChannel                = make(chan *http.Response)
	ownerDataChannel              = make(chan *http.Response)
)

type serviceCarDetails struct{}

func NewCarDetailsService() CarDetailsService {
	return &serviceCarDetails{}
}

func (*serviceCarDetails) GetDetails() entity.CarDetails {
	// goroutine get data from endpoint 1: https://myfakeapi.com/api/cars/1
	go carService.FetchData()

	// goroutine get data from endpoint 2: https://myfakeapi.com/api/users/1
	go ownerService.FetchData()

	// create carChannel to get the data from endpoint 1
	car, _ := getCarData()

	// create ownerChannel to get the data from endpoint 2
	owner, _ := getOwnerData()

	return entity.CarDetails{
		ID:        car.ID,
		Brand:     car.Brand,
		Model:     car.Model,
		Year:      car.Year,
		FirstName: owner.FirstName,
		LastName:  owner.LastName,
		Email:     owner.Email,
	}

}

func getCarData() (entity.Car, error) {
	r1 := <-carDataChannel
	var car entity.Car
	err := json.NewDecoder(r1.Body).Decode(&car)
	if err != nil {
		fmt.Print(err.Error())
		return car, err
	}
	return car, nil
}

func getOwnerData() (entity.Owner, error) {
	r1 := <-ownerDataChannel
	var owner entity.Owner
	err := json.NewDecoder(r1.Body).Decode(&owner)
	if err != nil {
		fmt.Print(err.Error())
		return owner, err
	}
	return owner, nil
}
