package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PhysicalOfficeAddress struct {
	City            *string `json:"city,omitempty"`
	CountryOrRegion *string `json:"countryOrRegion,omitempty"`
	ODataType       *string `json:"@odata.type,omitempty"`
	OfficeLocation  *string `json:"officeLocation,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	State           *string `json:"state,omitempty"`
	Street          *string `json:"street,omitempty"`
}
