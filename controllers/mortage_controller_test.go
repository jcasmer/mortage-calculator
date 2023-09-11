package controllers_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"mortage-calculator/client"
	"mortage-calculator/controllers"
	"mortage-calculator/controllers/viewmodels"
	"mortage-calculator/dto"
	"mortage-calculator/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type funcCalculateMortagePayment func(viewmodels.MortageRequest) dto.MortagePayment
type MockMortageClient struct {
	CalculateMortage funcCalculateMortagePayment
}

func (p *MockMortageClient) CalculateMortagePayment(req viewmodels.MortageRequest) dto.MortagePayment {
	return p.CalculateMortage(req)
}

func TestMortageController(t *testing.T) {

	cases := []struct {
		name      string
		request   interface{}
		assert    func(resp viewmodels.BaseResponse)
		morClient client.MortageClient
	}{
		{
			name:    "succes getting mortage payment",
			request: mocks.SuccesRequestMock,
			morClient: &MockMortageClient{
				CalculateMortage: func(req viewmodels.MortageRequest) dto.MortagePayment {
					return mocks.DtoMock
				},
			},
			assert: func(resp viewmodels.BaseResponse) {
				assert.Equal(t, mocks.DtoMock, resp.Data)
			},
		},
		{
			name: "Error  getting mortage payment",
			request: `{
			}`,
			morClient: &MockMortageClient{
				CalculateMortage: func(req viewmodels.MortageRequest) dto.MortagePayment {
					return mocks.DtoMock
				},
			},
			assert: func(resp viewmodels.BaseResponse) {
				assert.Equal(t, "json: cannot unmarshal string into Go value of type viewmodels.MortageRequest", resp.Error)
			},
		},
		{
			name:    "Error  getting mortage payment",
			request: viewmodels.MortageRequest{},
			morClient: &MockMortageClient{
				CalculateMortage: func(req viewmodels.MortageRequest) dto.MortagePayment {
					return mocks.DtoMock
				},
			},
			assert: func(resp viewmodels.BaseResponse) {
				assert.Equal(t, "property_price cannot be less than 0", resp.Error)
			},
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			mortageClient := func(log log.Logger) client.MortageClient {
				return tt.morClient
			}

			bc := &controllers.MortageController{
				MortageClientFactory: mortageClient,
			}

			jsonStr, _ := json.Marshal(tt.request)

			recorder := httptest.NewRecorder()
			request := httptest.NewRequest(http.MethodPost, "http://example.com/", bytes.NewReader(jsonStr))

			router := mux.NewRouter()
			router.HandleFunc("/", bc.GetMortagePayment).Methods(http.MethodPost)

			router.ServeHTTP(recorder, request)

			resp := recorder.Result()

			body, err := ioutil.ReadAll(resp.Body)
			require.NoError(t, err, "should return a readable response body")
			responseBody := viewmodels.BaseResponse{}
			err = json.Unmarshal(body, &responseBody)

			require.NoError(t, err, "should unmarshal the response wrapper without error")

			tt.assert(responseBody)
		})
	}
}
