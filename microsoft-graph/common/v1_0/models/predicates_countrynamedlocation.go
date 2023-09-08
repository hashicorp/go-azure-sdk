package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CountryNamedLocationOperationPredicate struct {
	CreatedDateTime                   *string
	DisplayName                       *string
	Id                                *string
	IncludeUnknownCountriesAndRegions *bool
	ModifiedDateTime                  *string
	ODataType                         *string
}

func (p CountryNamedLocationOperationPredicate) Matches(input CountryNamedLocation) bool {

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IncludeUnknownCountriesAndRegions != nil && (input.IncludeUnknownCountriesAndRegions == nil || *p.IncludeUnknownCountriesAndRegions != *input.IncludeUnknownCountriesAndRegions) {
		return false
	}

	if p.ModifiedDateTime != nil && (input.ModifiedDateTime == nil || *p.ModifiedDateTime != *input.ModifiedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
