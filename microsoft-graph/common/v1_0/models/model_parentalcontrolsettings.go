package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ParentalControlSettings struct {
	CountriesBlockedForMinors *[]string `json:"countriesBlockedForMinors,omitempty"`
	LegalAgeGroupRule         *string   `json:"legalAgeGroupRule,omitempty"`
	ODataType                 *string   `json:"@odata.type,omitempty"`
}
