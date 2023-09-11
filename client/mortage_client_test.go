package client_test

import (
	"bytes"
	"log"
	"testing"

	"mortage-calculator/client"
	"mortage-calculator/mocks"

	"mortage-calculator/controllers/viewmodels"

	"mortage-calculator/dto"

	"github.com/stretchr/testify/assert"
)

var (
	buf    bytes.Buffer
	logger = log.New(&buf, "logger: ", log.Lshortfile)
)

func TestNewMortageClient(t *testing.T) {

	mortClient := client.NewMortageClient(*logger)
	assert.NotNil(t, mortClient)
}

func TestNewBrowseContext(t *testing.T) {
	cases := []struct {
		name     string
		request  viewmodels.MortageRequest
		response dto.MortagePayment
		assert   func(expected dto.MortagePayment, output dto.MortagePayment)
	}{
		{
			name:     "Success",
			request:  mocks.SuccesRequestMock,
			response: mocks.DtoMock,
			assert: func(expected dto.MortagePayment, output dto.MortagePayment) {

				assert.Equal(t, expected, output)

			},
		},
		{
			name:     "Success Bi-weekly",
			request:  mocks.SuccesRequestBiWeeklyMock,
			response: mocks.DtoBiWeeklyMock,
			assert: func(expected dto.MortagePayment, output dto.MortagePayment) {

				assert.Equal(t, expected, output)
			},
		},
		{
			name:     "Success accelerated bi-weekly",
			request:  mocks.SuccesRequestAccBiWeeklyMock,
			response: mocks.DtoAccBiWeeklyMock,
			assert: func(expected dto.MortagePayment, output dto.MortagePayment) {

				assert.Equal(t, expected, output)
			},
		},
		{
			name:     "Success with no values",
			request:  mocks.RequestZeroProceMock,
			response: mocks.DtoMock,
			assert: func(expected dto.MortagePayment, output dto.MortagePayment) {

				assert.NotEqual(t, expected, output)
				assert.Equal(t, 0.0, output.TotalMortage)
				assert.Equal(t, 0.0, output.CMHC)
			},
		},

		{
			name:     "Success with down payment is higher",
			request:  mocks.RequestHighDownMock,
			response: mocks.DtoMock,
			assert: func(expected dto.MortagePayment, output dto.MortagePayment) {

				assert.NotEqual(t, expected, output)
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			data := client.NewMortageClient(*logger)
			response := data.CalculateMortagePayment(c.request)
			c.assert(c.response, response)
		})

	}
}
