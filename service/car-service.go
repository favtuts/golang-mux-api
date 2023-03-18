package service

import (
	"fmt"
	"net/http"
)

type CarService interface {
	FetchData()
}

const (
	carServiceUrl = "https://myfakeapi.com/api/cars/1"
)

type fetchCarDataService struct{}

func NewCarService() CarService {
	return &fetchCarDataService{}
}

func (*fetchCarDataService) FetchData() {
	client := http.Client{}
	fmt.Printf("Fetching the url %s", carServiceUrl)

	// Call the external API
	resp, _ := client.Get(carServiceUrl)

	// Write response to the channel (TODO)
	carDataChannel <- resp
}
