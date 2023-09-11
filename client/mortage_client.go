package client

import (
	"log"
	"math"
	"mortage-calculator/controllers/viewmodels"
	"mortage-calculator/dto"
)

type MortageClient interface {
	CalculateMortagePayment(viewmodels.MortageRequest) dto.MortagePayment
}

type mortageClient struct {
	log log.Logger
}

func NewMortageClient(log log.Logger) MortageClient {
	return &mortageClient{
		log: log,
	}
}

func (mc *mortageClient) CalculateMortagePayment(req viewmodels.MortageRequest) dto.MortagePayment {

	totalM := float64(req.PropertyPrice - ((req.DownPayment / 100) * req.PropertyPrice))
	chmc := totalM * getCHMC(req.DownPayment)

	totalM += chmc

	n := float64(12 * req.AmortizationPeriod)
	r := ((req.AnnualInterestRate / 100) / 12)
	interest := math.Pow(1+r, n)
	m := totalM * ((r * interest) / (interest - 1))

	if req.PaymentSchedule == "bi-weekly" {
		m = (m * 12) / 26
	}

	if req.PaymentSchedule == "accelerated bi-weekly" {
		m = (m / 2)
	}
	return dto.MortagePayment{
		TotalMortage: math.Round(m),
		CMHC:         chmc,
	}
}

func getCHMC(downPayment float64) float64 {

	downPayment = 100 - downPayment

	if downPayment > 90 && downPayment <= 95 {
		return 0.04
	}
	if downPayment > 85 && downPayment <= 90 {
		return 0.031
	}
	if downPayment > 80 && downPayment <= 85 {
		return 0.028
	}
	return 0.0
}
