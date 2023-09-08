package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TimeZoneBaseOperationPredicate struct {
	Name      *string
	ODataType *string
}

func (p TimeZoneBaseOperationPredicate) Matches(input TimeZoneBase) bool {

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
