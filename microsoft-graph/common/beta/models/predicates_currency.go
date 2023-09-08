package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CurrencyOperationPredicate struct {
	AmountDecimalPlaces     *string
	AmountRoundingPrecision *float64
	Code                    *string
	DisplayName             *string
	Id                      *string
	LastModifiedDateTime    *string
	ODataType               *string
	Symbol                  *string
}

func (p CurrencyOperationPredicate) Matches(input Currency) bool {

	if p.AmountDecimalPlaces != nil && (input.AmountDecimalPlaces == nil || *p.AmountDecimalPlaces != *input.AmountDecimalPlaces) {
		return false
	}

	if p.AmountRoundingPrecision != nil && (input.AmountRoundingPrecision == nil || *p.AmountRoundingPrecision != *input.AmountRoundingPrecision) {
		return false
	}

	if p.Code != nil && (input.Code == nil || *p.Code != *input.Code) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
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

	if p.Symbol != nil && (input.Symbol == nil || *p.Symbol != *input.Symbol) {
		return false
	}

	return true
}
