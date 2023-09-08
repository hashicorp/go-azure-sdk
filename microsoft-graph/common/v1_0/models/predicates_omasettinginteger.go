package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OmaSettingIntegerOperationPredicate struct {
	Description *string
	DisplayName *string
	ODataType   *string
	OmaUri      *string
	Value       *int64
}

func (p OmaSettingIntegerOperationPredicate) Matches(input OmaSettingInteger) bool {

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OmaUri != nil && (input.OmaUri == nil || *p.OmaUri != *input.OmaUri) {
		return false
	}

	if p.Value != nil && (input.Value == nil || *p.Value != *input.Value) {
		return false
	}

	return true
}
