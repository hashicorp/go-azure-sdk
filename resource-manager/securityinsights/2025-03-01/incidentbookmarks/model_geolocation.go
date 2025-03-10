package incidentbookmarks

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GeoLocation struct {
	Asn         *int64   `json:"asn,omitempty"`
	City        *string  `json:"city,omitempty"`
	CountryCode *string  `json:"countryCode,omitempty"`
	CountryName *string  `json:"countryName,omitempty"`
	Latitude    *float64 `json:"latitude,omitempty"`
	Longitude   *float64 `json:"longitude,omitempty"`
	State       *string  `json:"state,omitempty"`
}
