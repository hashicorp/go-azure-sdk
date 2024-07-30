package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PhysicalAddress struct {
	City            *string              `json:"city,omitempty"`
	CountryOrRegion *string              `json:"countryOrRegion,omitempty"`
	ODataType       *string              `json:"@odata.type,omitempty"`
	PostOfficeBox   *string              `json:"postOfficeBox,omitempty"`
	PostalCode      *string              `json:"postalCode,omitempty"`
	State           *string              `json:"state,omitempty"`
	Street          *string              `json:"street,omitempty"`
	Type            *PhysicalAddressType `json:"type,omitempty"`
}
