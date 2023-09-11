package viewmodels_test

import (
	"testing"

	"mortage-calculator/controllers/viewmodels"

	"github.com/stretchr/testify/assert"
)

func TestMortageRequest(t *testing.T) {

	cases := []struct {
		name           string
		MortageRequest viewmodels.MortageRequest
		assert         func(err error)
	}{
		{
			name: "Success response when input is valid ",
			MortageRequest: viewmodels.MortageRequest{
				PropertyPrice:      100000,
				DownPayment:        5,
				AnnualInterestRate: 4.69,
				AmortizationPeriod: 25,
				PaymentSchedule:    "monthly",
			},
			assert: func(err error) {
				assert.Nil(t, err)
			},
		},
		{
			name: "Error PropertyPrice is zero",
			MortageRequest: viewmodels.MortageRequest{
				PropertyPrice:      0,
				DownPayment:        5,
				AnnualInterestRate: 4.69,
				AmortizationPeriod: 25,
				PaymentSchedule:    "monthly",
			},
			assert: func(err error) {
				assert.NotNil(t, err)
				assert.Equal(t, err.Error(), "property_price cannot be less than 0")
			},
		},
		{
			name: "Error down_payment is zero",
			MortageRequest: viewmodels.MortageRequest{
				PropertyPrice:      10000,
				DownPayment:        0,
				AnnualInterestRate: 4.69,
				AmortizationPeriod: 25,
				PaymentSchedule:    "monthly",
			},
			assert: func(err error) {
				assert.NotNil(t, err)
				assert.Equal(t, err.Error(), "down_payment must be between 0 and 100")
			},
		},
		{
			name: "Error amortization_period is not valid",
			MortageRequest: viewmodels.MortageRequest{
				PropertyPrice:      10000,
				DownPayment:        5,
				AnnualInterestRate: 4.69,
				AmortizationPeriod: 45,
				PaymentSchedule:    "monthly",
			},
			assert: func(err error) {
				assert.NotNil(t, err)
				assert.Equal(t, err.Error(), "amortization_period must be one of: 5, 10, 15, 20, 25, 30")
			},
		},
		{
			name: "Error payment_schedule is not valid",
			MortageRequest: viewmodels.MortageRequest{
				PropertyPrice:      10000,
				DownPayment:        5,
				AnnualInterestRate: 4.69,
				AmortizationPeriod: 25,
				PaymentSchedule:    "year",
			},
			assert: func(err error) {
				assert.NotNil(t, err)
				assert.Equal(t, err.Error(), "payment_schedule must be one of: accelerated bi-weekly, bi-weekly, monthly")
			},
		},
		{
			name: "Error amortization_period & down_payment are not valid",
			MortageRequest: viewmodels.MortageRequest{
				PropertyPrice:      10000,
				DownPayment:        5,
				AnnualInterestRate: 4.69,
				AmortizationPeriod: 30,
				PaymentSchedule:    "monthly",
			},
			assert: func(err error) {
				assert.NotNil(t, err)
				assert.Equal(t, err.Error(), "amortization_period could not be more than 25 when down_payment is less than 20%")
			},
		},
	}

	for _, c := range cases {
		err := c.MortageRequest.Validate()

		c.assert(err)
	}
}
