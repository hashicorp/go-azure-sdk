package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DateTimeTimeZoneTypeOperationPredicate struct {
	DateTime  *string
	ODataType *string
}

func (p DateTimeTimeZoneTypeOperationPredicate) Matches(input DateTimeTimeZoneType) bool {

	if p.DateTime != nil && (input.DateTime == nil || *p.DateTime != *input.DateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
