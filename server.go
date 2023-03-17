package main

import (
	"fmt"
	"golang-mux-api/controller"
	router "golang-mux-api/http"
	"golang-mux-api/service"
	"net/http"
)

var (
	carDetailsService    service.CarDetailsService       = service.NewCarDetailsService()
	carDetailsController controller.CarDetailsController = controller.NewCarDetailsController(carDetailsService)
	httpRouter           router.Router                   = router.NewChiRouter()
)

func main() {
	const port string = ":8000"

	httpRouter.GET("/", func(resp http.ResponseWriter, request *http.Request) {
		fmt.Fprint(resp, "Up and running...")
	})

	httpRouter.GET("/carDetails", carDetailsController.GetCarDetails)
	httpRouter.SERVE(port)
}
