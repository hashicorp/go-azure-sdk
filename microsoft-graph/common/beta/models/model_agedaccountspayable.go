package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AgedAccountsPayable struct {
	AgedAsOfDate       *string  `json:"agedAsOfDate,omitempty"`
	BalanceDue         *float64 `json:"balanceDue,omitempty"`
	CurrencyCode       *string  `json:"currencyCode,omitempty"`
	CurrentAmount      *float64 `json:"currentAmount,omitempty"`
	Id                 *string  `json:"id,omitempty"`
	Name               *string  `json:"name,omitempty"`
	ODataType          *string  `json:"@odata.type,omitempty"`
	Period1Amount      *float64 `json:"period1Amount,omitempty"`
	Period2Amount      *float64 `json:"period2Amount,omitempty"`
	Period3Amount      *float64 `json:"period3Amount,omitempty"`
	PeriodLengthFilter *string  `json:"periodLengthFilter,omitempty"`
	VendorId           *string  `json:"vendorId,omitempty"`
	VendorNumber       *string  `json:"vendorNumber,omitempty"`
}
