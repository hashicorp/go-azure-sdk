package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomTimeZoneOperationPredicate struct {
	Bias      *int64
	Name      *string
	ODataType *string
}

func (p CustomTimeZoneOperationPredicate) Matches(input CustomTimeZone) bool {

	if p.Bias != nil && (input.Bias == nil || *p.Bias != *input.Bias) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
