package controllers

import (
	"encoding/json"
	"log"
	"mortage-calculator/client"
	"mortage-calculator/controllers/viewmodels"
	"net/http"
)

type MortageController struct {
	MortageClientFactory func(log log.Logger) client.MortageClient
}

func (mc *MortageController) GetMortagePayment(w http.ResponseWriter, r *http.Request) {

	response := viewmodels.BaseResponse{}
	mortageRequest := viewmodels.MortageRequest{}
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&mortageRequest)
	if err != nil {
		log.Println("error while attempting to parse (unmarshal) the request")
		response.Error = err.Error()
		SendResponse(w, http.StatusBadRequest, response)
		return
	}

	err = mortageRequest.Validate()
	if err != nil {
		log.Println("error validating request")
		response.Error = err.Error()
		SendResponse(w, http.StatusBadRequest, response)
		return
	}

	client := mc.MortageClientFactory(*log.Default())
	clientDto := client.CalculateMortagePayment(mortageRequest)

	response.Data = clientDto
	SendResponse(w, http.StatusOK, response)
}

func SendResponse(w http.ResponseWriter, httpStatus int, data viewmodels.BaseResponse) {
	w.Header().Add("Content-Type", "application/json")
	bytes, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(httpStatus)
	w.Write(bytes)
}
