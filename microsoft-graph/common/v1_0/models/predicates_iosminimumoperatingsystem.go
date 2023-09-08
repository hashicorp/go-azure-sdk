package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosMinimumOperatingSystemOperationPredicate struct {
	ODataType *string
	V100      *bool
	V110      *bool
	V120      *bool
	V130      *bool
	V140      *bool
	V150      *bool
	V80       *bool
	V90       *bool
}

func (p IosMinimumOperatingSystemOperationPredicate) Matches(input IosMinimumOperatingSystem) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.V100 != nil && (input.V100 == nil || *p.V100 != *input.V100) {
		return false
	}

	if p.V110 != nil && (input.V110 == nil || *p.V110 != *input.V110) {
		return false
	}

	if p.V120 != nil && (input.V120 == nil || *p.V120 != *input.V120) {
		return false
	}

	if p.V130 != nil && (input.V130 == nil || *p.V130 != *input.V130) {
		return false
	}

	if p.V140 != nil && (input.V140 == nil || *p.V140 != *input.V140) {
		return false
	}

	if p.V150 != nil && (input.V150 == nil || *p.V150 != *input.V150) {
		return false
	}

	if p.V80 != nil && (input.V80 == nil || *p.V80 != *input.V80) {
		return false
	}

	if p.V90 != nil && (input.V90 == nil || *p.V90 != *input.V90) {
		return false
	}

	return true
}
