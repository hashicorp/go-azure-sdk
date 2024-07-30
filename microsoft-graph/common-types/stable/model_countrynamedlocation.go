package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CountryNamedLocation struct {
	CountriesAndRegions               *[]string                                `json:"countriesAndRegions,omitempty"`
	CountryLookupMethod               *CountryNamedLocationCountryLookupMethod `json:"countryLookupMethod,omitempty"`
	CreatedDateTime                   *string                                  `json:"createdDateTime,omitempty"`
	DisplayName                       *string                                  `json:"displayName,omitempty"`
	Id                                *string                                  `json:"id,omitempty"`
	IncludeUnknownCountriesAndRegions *bool                                    `json:"includeUnknownCountriesAndRegions,omitempty"`
	ModifiedDateTime                  *string                                  `json:"modifiedDateTime,omitempty"`
	ODataType                         *string                                  `json:"@odata.type,omitempty"`
}
