package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PaymentTermOperationPredicate struct {
	CalculateDiscountOnCreditMemos *bool
	Code                           *string
	DiscountDateCalculation        *string
	DiscountPercent                *float64
	DisplayName                    *string
	DueDateCalculation             *string
	Id                             *string
	LastModifiedDateTime           *string
	ODataType                      *string
}

func (p PaymentTermOperationPredicate) Matches(input PaymentTerm) bool {

	if p.CalculateDiscountOnCreditMemos != nil && (input.CalculateDiscountOnCreditMemos == nil || *p.CalculateDiscountOnCreditMemos != *input.CalculateDiscountOnCreditMemos) {
		return false
	}

	if p.Code != nil && (input.Code == nil || *p.Code != *input.Code) {
		return false
	}

	if p.DiscountDateCalculation != nil && (input.DiscountDateCalculation == nil || *p.DiscountDateCalculation != *input.DiscountDateCalculation) {
		return false
	}

	if p.DiscountPercent != nil && (input.DiscountPercent == nil || *p.DiscountPercent != *input.DiscountPercent) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.DueDateCalculation != nil && (input.DueDateCalculation == nil || *p.DueDateCalculation != *input.DueDateCalculation) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
