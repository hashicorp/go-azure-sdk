package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PowerliftIncidentMetadataOperationPredicate struct {
	Application       *string
	ClientVersion     *string
	CreatedAtDateTime *string
	EasyId            *string
	Locale            *string
	ODataType         *string
	Platform          *string
	PowerliftId       *string
}

func (p PowerliftIncidentMetadataOperationPredicate) Matches(input PowerliftIncidentMetadata) bool {

	if p.Application != nil && (input.Application == nil || *p.Application != *input.Application) {
		return false
	}

	if p.ClientVersion != nil && (input.ClientVersion == nil || *p.ClientVersion != *input.ClientVersion) {
		return false
	}

	if p.CreatedAtDateTime != nil && (input.CreatedAtDateTime == nil || *p.CreatedAtDateTime != *input.CreatedAtDateTime) {
		return false
	}

	if p.EasyId != nil && (input.EasyId == nil || *p.EasyId != *input.EasyId) {
		return false
	}

	if p.Locale != nil && (input.Locale == nil || *p.Locale != *input.Locale) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Platform != nil && (input.Platform == nil || *p.Platform != *input.Platform) {
		return false
	}

	if p.PowerliftId != nil && (input.PowerliftId == nil || *p.PowerliftId != *input.PowerliftId) {
		return false
	}

	return true
}
