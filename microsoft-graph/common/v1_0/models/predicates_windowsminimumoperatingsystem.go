package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsMinimumOperatingSystemOperationPredicate struct {
	ODataType *string
	V100      *bool
	V80       *bool
	V81       *bool
}

func (p WindowsMinimumOperatingSystemOperationPredicate) Matches(input WindowsMinimumOperatingSystem) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.V100 != nil && (input.V100 == nil || *p.V100 != *input.V100) {
		return false
	}

	if p.V80 != nil && (input.V80 == nil || *p.V80 != *input.V80) {
		return false
	}

	if p.V81 != nil && (input.V81 == nil || *p.V81 != *input.V81) {
		return false
	}

	return true
}
