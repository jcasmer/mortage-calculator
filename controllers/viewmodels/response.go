package viewmodels

import "mortage-calculator/dto"

type BaseResponse struct {
	Data  dto.MortagePayment `json:"data,omitempty"`
	Error string             `json:"error,omitempty"`
}
