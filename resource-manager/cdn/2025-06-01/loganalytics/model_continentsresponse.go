package loganalytics

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContinentsResponse struct {
	Continents       *[]ContinentsResponseContinentsInlined       `json:"continents,omitempty"`
	CountryOrRegions *[]ContinentsResponseCountryOrRegionsInlined `json:"countryOrRegions,omitempty"`
}
