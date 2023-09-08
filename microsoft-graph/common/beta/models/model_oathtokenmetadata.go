package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OathTokenMetadata struct {
	Enabled                *bool       `json:"enabled,omitempty"`
	Manufacturer           *string     `json:"manufacturer,omitempty"`
	ManufacturerProperties *[]KeyValue `json:"manufacturerProperties,omitempty"`
	ODataType              *string     `json:"@odata.type,omitempty"`
	SerialNumber           *string     `json:"serialNumber,omitempty"`
	TokenType              *string     `json:"tokenType,omitempty"`
}
