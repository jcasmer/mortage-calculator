package api

import (
	"log"
	"mortage-calculator/client"
	"mortage-calculator/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func Routes(log log.Logger) http.Handler {

	mortageController := generateController(log)

	router := mux.NewRouter()

	router.HandleFunc("/", mortageController.GetMortagePayment).Methods(http.MethodPost)

	return router
}

func generateController(logg log.Logger) controllers.MortageController {

	mortageClientFactory := func(log log.Logger) client.MortageClient {
		return client.NewMortageClient(log)
	}
	return controllers.MortageController{
		MortageClientFactory: mortageClientFactory,
	}
}
