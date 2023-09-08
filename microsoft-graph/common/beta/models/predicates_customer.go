package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomerOperationPredicate struct {
	Blocked               *string
	CurrencyCode          *string
	CurrencyId            *string
	DisplayName           *string
	Email                 *string
	Id                    *string
	LastModifiedDateTime  *string
	Number                *string
	ODataType             *string
	PaymentMethodId       *string
	PaymentTermsId        *string
	PhoneNumber           *string
	ShipmentMethodId      *string
	TaxAreaDisplayName    *string
	TaxAreaId             *string
	TaxLiable             *bool
	TaxRegistrationNumber *string
	Type                  *string
	Website               *string
}

func (p CustomerOperationPredicate) Matches(input Customer) bool {

	if p.Blocked != nil && (input.Blocked == nil || *p.Blocked != *input.Blocked) {
		return false
	}

	if p.CurrencyCode != nil && (input.CurrencyCode == nil || *p.CurrencyCode != *input.CurrencyCode) {
		return false
	}

	if p.CurrencyId != nil && (input.CurrencyId == nil || *p.CurrencyId != *input.CurrencyId) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Email != nil && (input.Email == nil || *p.Email != *input.Email) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.Number != nil && (input.Number == nil || *p.Number != *input.Number) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PaymentMethodId != nil && (input.PaymentMethodId == nil || *p.PaymentMethodId != *input.PaymentMethodId) {
		return false
	}

	if p.PaymentTermsId != nil && (input.PaymentTermsId == nil || *p.PaymentTermsId != *input.PaymentTermsId) {
		return false
	}

	if p.PhoneNumber != nil && (input.PhoneNumber == nil || *p.PhoneNumber != *input.PhoneNumber) {
		return false
	}

	if p.ShipmentMethodId != nil && (input.ShipmentMethodId == nil || *p.ShipmentMethodId != *input.ShipmentMethodId) {
		return false
	}

	if p.TaxAreaDisplayName != nil && (input.TaxAreaDisplayName == nil || *p.TaxAreaDisplayName != *input.TaxAreaDisplayName) {
		return false
	}

	if p.TaxAreaId != nil && (input.TaxAreaId == nil || *p.TaxAreaId != *input.TaxAreaId) {
		return false
	}

	if p.TaxLiable != nil && (input.TaxLiable == nil || *p.TaxLiable != *input.TaxLiable) {
		return false
	}

	if p.TaxRegistrationNumber != nil && (input.TaxRegistrationNumber == nil || *p.TaxRegistrationNumber != *input.TaxRegistrationNumber) {
		return false
	}

	if p.Type != nil && (input.Type == nil || *p.Type != *input.Type) {
		return false
	}

	if p.Website != nil && (input.Website == nil || *p.Website != *input.Website) {
		return false
	}

	return true
}
