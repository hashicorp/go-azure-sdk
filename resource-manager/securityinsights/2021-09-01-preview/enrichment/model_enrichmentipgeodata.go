package enrichment

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnrichmentIPGeodata struct {
	Asn              *string `json:"asn,omitempty"`
	Carrier          *string `json:"carrier,omitempty"`
	City             *string `json:"city,omitempty"`
	CityCf           *int64  `json:"cityCf,omitempty"`
	Continent        *string `json:"continent,omitempty"`
	Country          *string `json:"country,omitempty"`
	CountryCf        *int64  `json:"countryCf,omitempty"`
	IPAddr           *string `json:"ipAddr,omitempty"`
	IPRoutingType    *string `json:"ipRoutingType,omitempty"`
	Latitude         *string `json:"latitude,omitempty"`
	Longitude        *string `json:"longitude,omitempty"`
	Organization     *string `json:"organization,omitempty"`
	OrganizationType *string `json:"organizationType,omitempty"`
	Region           *string `json:"region,omitempty"`
	State            *string `json:"state,omitempty"`
	StateCf          *int64  `json:"stateCf,omitempty"`
	StateCode        *string `json:"stateCode,omitempty"`
}
