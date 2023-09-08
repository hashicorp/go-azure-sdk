package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OmaSettingFloatingPointOperationPredicate struct {
	Description *string
	DisplayName *string
	ODataType   *string
	OmaUri      *string
}

func (p OmaSettingFloatingPointOperationPredicate) Matches(input OmaSettingFloatingPoint) bool {

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

	return true
}
