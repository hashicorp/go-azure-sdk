package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Vendor struct {
	Address               *PostalAddressType `json:"address,omitempty"`
	Balance               *float64           `json:"balance,omitempty"`
	Blocked               *string            `json:"blocked,omitempty"`
	Currency              *Currency          `json:"currency,omitempty"`
	CurrencyCode          *string            `json:"currencyCode,omitempty"`
	CurrencyId            *string            `json:"currencyId,omitempty"`
	DisplayName           *string            `json:"displayName,omitempty"`
	Email                 *string            `json:"email,omitempty"`
	Id                    *string            `json:"id,omitempty"`
	LastModifiedDateTime  *string            `json:"lastModifiedDateTime,omitempty"`
	Number                *string            `json:"number,omitempty"`
	ODataType             *string            `json:"@odata.type,omitempty"`
	PaymentMethod         *PaymentMethod     `json:"paymentMethod,omitempty"`
	PaymentMethodId       *string            `json:"paymentMethodId,omitempty"`
	PaymentTerm           *PaymentTerm       `json:"paymentTerm,omitempty"`
	PaymentTermsId        *string            `json:"paymentTermsId,omitempty"`
	PhoneNumber           *string            `json:"phoneNumber,omitempty"`
	Picture               *[]Picture         `json:"picture,omitempty"`
	TaxLiable             *bool              `json:"taxLiable,omitempty"`
	TaxRegistrationNumber *string            `json:"taxRegistrationNumber,omitempty"`
	Website               *string            `json:"website,omitempty"`
}
