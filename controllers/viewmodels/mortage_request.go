package viewmodels

import "errors"

type MortageRequest struct {
	PropertyPrice      float64 `json:"property_rice" validate:"required"`
	DownPayment        float64 `json:"down_payment" validate:"required"`
	AnnualInterestRate float64 `json:"annual_interest_rate" validate:"required"`
	AmortizationPeriod int     `json:"amortization_period" validate:"required"`
	PaymentSchedule    string  `json:"payment_schedule" validate:"required"`
}

func (mr MortageRequest) Validate() error {

	if mr.PropertyPrice <= 0.0 {
		return errors.New("property_price cannot be less than 0")
	}

	if mr.DownPayment <= 0.0 || mr.DownPayment > 100.0 {
		return errors.New("down_payment must be between 0 and 100")
	}
	validAmortization := validateAmortizationPeriod(mr.AmortizationPeriod)
	if !validAmortization {
		return errors.New("amortization_period must be one of: 5, 10, 15, 20, 25, 30")
	}
	validPayment := validatePaymentSchedule(mr.PaymentSchedule)
	if !validPayment {
		return errors.New("payment_schedule must be one of: accelerated bi-weekly, bi-weekly, monthly")
	}

	if (mr.DownPayment/100) < 20.0 && mr.AmortizationPeriod > 25 {
		return errors.New("amortization_period could not be more than 25 when down_payment is less than 20%")
	}

	return nil
}

func validatePaymentSchedule(paymentSchedule string) bool {

	switch paymentSchedule {
	case
		"accelerated bi-weekly",
		"bi-weekly",
		"monthly":
		return true
	}
	return false
}

func validateAmortizationPeriod(amortizationPeriod int) bool {

	switch amortizationPeriod {
	case
		5, 10, 15, 20, 25, 30:
		return true
	}
	return false

}
