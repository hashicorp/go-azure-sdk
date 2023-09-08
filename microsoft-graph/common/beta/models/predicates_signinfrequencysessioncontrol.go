package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInFrequencySessionControlOperationPredicate struct {
	IsEnabled *bool
	ODataType *string
	Value     *int64
}

func (p SignInFrequencySessionControlOperationPredicate) Matches(input SignInFrequencySessionControl) bool {

	if p.IsEnabled != nil && (input.IsEnabled == nil || *p.IsEnabled != *input.IsEnabled) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Value != nil && (input.Value == nil || *p.Value != *input.Value) {
		return false
	}

	return true
}
