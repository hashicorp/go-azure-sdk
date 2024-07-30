package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInLocation struct {
	City            *string         `json:"city,omitempty"`
	CountryOrRegion *string         `json:"countryOrRegion,omitempty"`
	GeoCoordinates  *GeoCoordinates `json:"geoCoordinates,omitempty"`
	ODataType       *string         `json:"@odata.type,omitempty"`
	State           *string         `json:"state,omitempty"`
}
