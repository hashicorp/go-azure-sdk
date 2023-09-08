package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AgedAccountsPayableOperationPredicate struct {
	AgedAsOfDate       *string
	BalanceDue         *float64
	CurrencyCode       *string
	CurrentAmount      *float64
	Id                 *string
	Name               *string
	ODataType          *string
	Period1Amount      *float64
	Period2Amount      *float64
	Period3Amount      *float64
	PeriodLengthFilter *string
	VendorId           *string
	VendorNumber       *string
}

func (p AgedAccountsPayableOperationPredicate) Matches(input AgedAccountsPayable) bool {

	if p.AgedAsOfDate != nil && (input.AgedAsOfDate == nil || *p.AgedAsOfDate != *input.AgedAsOfDate) {
		return false
	}

	if p.BalanceDue != nil && (input.BalanceDue == nil || *p.BalanceDue != *input.BalanceDue) {
		return false
	}

	if p.CurrencyCode != nil && (input.CurrencyCode == nil || *p.CurrencyCode != *input.CurrencyCode) {
		return false
	}

	if p.CurrentAmount != nil && (input.CurrentAmount == nil || *p.CurrentAmount != *input.CurrentAmount) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Period1Amount != nil && (input.Period1Amount == nil || *p.Period1Amount != *input.Period1Amount) {
		return false
	}

	if p.Period2Amount != nil && (input.Period2Amount == nil || *p.Period2Amount != *input.Period2Amount) {
		return false
	}

	if p.Period3Amount != nil && (input.Period3Amount == nil || *p.Period3Amount != *input.Period3Amount) {
		return false
	}

	if p.PeriodLengthFilter != nil && (input.PeriodLengthFilter == nil || *p.PeriodLengthFilter != *input.PeriodLengthFilter) {
		return false
	}

	if p.VendorId != nil && (input.VendorId == nil || *p.VendorId != *input.VendorId) {
		return false
	}

	if p.VendorNumber != nil && (input.VendorNumber == nil || *p.VendorNumber != *input.VendorNumber) {
		return false
	}

	return true
}
