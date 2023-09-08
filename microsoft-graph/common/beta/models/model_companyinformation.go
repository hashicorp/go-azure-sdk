package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CompanyInformation struct {
	Address                    *PostalAddressType `json:"address,omitempty"`
	CurrencyCode               *string            `json:"currencyCode,omitempty"`
	CurrentFiscalYearStartDate *string            `json:"currentFiscalYearStartDate,omitempty"`
	DisplayName                *string            `json:"displayName,omitempty"`
	Email                      *string            `json:"email,omitempty"`
	FaxNumber                  *string            `json:"faxNumber,omitempty"`
	Id                         *string            `json:"id,omitempty"`
	Industry                   *string            `json:"industry,omitempty"`
	LastModifiedDateTime       *string            `json:"lastModifiedDateTime,omitempty"`
	ODataType                  *string            `json:"@odata.type,omitempty"`
	PhoneNumber                *string            `json:"phoneNumber,omitempty"`
	Picture                    *string            `json:"picture,omitempty"`
	TaxRegistrationNumber      *string            `json:"taxRegistrationNumber,omitempty"`
	Website                    *string            `json:"website,omitempty"`
}
