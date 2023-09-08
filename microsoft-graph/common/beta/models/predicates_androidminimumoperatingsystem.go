package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidMinimumOperatingSystemOperationPredicate struct {
	ODataType *string
	V100      *bool
	V110      *bool
	V40       *bool
	V403      *bool
	V41       *bool
	V42       *bool
	V43       *bool
	V44       *bool
	V50       *bool
	V51       *bool
	V60       *bool
	V70       *bool
	V71       *bool
	V80       *bool
	V81       *bool
	V90       *bool
}

func (p AndroidMinimumOperatingSystemOperationPredicate) Matches(input AndroidMinimumOperatingSystem) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.V100 != nil && (input.V100 == nil || *p.V100 != *input.V100) {
		return false
	}

	if p.V110 != nil && (input.V110 == nil || *p.V110 != *input.V110) {
		return false
	}

	if p.V40 != nil && (input.V40 == nil || *p.V40 != *input.V40) {
		return false
	}

	if p.V403 != nil && (input.V403 == nil || *p.V403 != *input.V403) {
		return false
	}

	if p.V41 != nil && (input.V41 == nil || *p.V41 != *input.V41) {
		return false
	}

	if p.V42 != nil && (input.V42 == nil || *p.V42 != *input.V42) {
		return false
	}

	if p.V43 != nil && (input.V43 == nil || *p.V43 != *input.V43) {
		return false
	}

	if p.V44 != nil && (input.V44 == nil || *p.V44 != *input.V44) {
		return false
	}

	if p.V50 != nil && (input.V50 == nil || *p.V50 != *input.V50) {
		return false
	}

	if p.V51 != nil && (input.V51 == nil || *p.V51 != *input.V51) {
		return false
	}

	if p.V60 != nil && (input.V60 == nil || *p.V60 != *input.V60) {
		return false
	}

	if p.V70 != nil && (input.V70 == nil || *p.V70 != *input.V70) {
		return false
	}

	if p.V71 != nil && (input.V71 == nil || *p.V71 != *input.V71) {
		return false
	}

	if p.V80 != nil && (input.V80 == nil || *p.V80 != *input.V80) {
		return false
	}

	if p.V81 != nil && (input.V81 == nil || *p.V81 != *input.V81) {
		return false
	}

	if p.V90 != nil && (input.V90 == nil || *p.V90 != *input.V90) {
		return false
	}

	return true
}
