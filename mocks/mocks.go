package mocks

import (
	"mortage-calculator/controllers/viewmodels"
	"mortage-calculator/dto"
)

var (
	SuccesRequestMock = viewmodels.MortageRequest{
		PropertyPrice:      100000,
		DownPayment:        5,
		AnnualInterestRate: 4.69,
		AmortizationPeriod: 25,
		PaymentSchedule:    "monthly",
	}
	SuccesRequestBiWeeklyMock = viewmodels.MortageRequest{
		PropertyPrice:      100000,
		DownPayment:        5,
		AnnualInterestRate: 4.69,
		AmortizationPeriod: 25,
		PaymentSchedule:    "bi-weekly",
	}
	SuccesRequestAccBiWeeklyMock = viewmodels.MortageRequest{
		PropertyPrice:      100000,
		DownPayment:        5,
		AnnualInterestRate: 4.69,
		AmortizationPeriod: 25,
		PaymentSchedule:    "accelerated bi-weekly",
	}
	RequestZeroProceMock = viewmodels.MortageRequest{
		PropertyPrice:      0,
		DownPayment:        5,
		AnnualInterestRate: 4.69,
		AmortizationPeriod: 25,
		PaymentSchedule:    "monthly",
	}

	RequestHighDownMock = viewmodels.MortageRequest{
		PropertyPrice:      100000,
		DownPayment:        21,
		AnnualInterestRate: 4.69,
		AmortizationPeriod: 25,
		PaymentSchedule:    "monthly",
	}

	DtoMock = dto.MortagePayment{
		TotalMortage: 560,
		CMHC:         3800,
	}
	DtoZeroMock = dto.MortagePayment{
		TotalMortage: 0,
		CMHC:         0,
	}
	DtoBiWeeklyMock = dto.MortagePayment{
		TotalMortage: 258,
		CMHC:         3800,
	}
	DtoAccBiWeeklyMock = dto.MortagePayment{
		TotalMortage: 280,
		CMHC:         3800,
	}
)
