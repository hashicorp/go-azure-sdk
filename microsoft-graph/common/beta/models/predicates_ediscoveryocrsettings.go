package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryOcrSettingsOperationPredicate struct {
	IsEnabled    *bool
	MaxImageSize *int64
	ODataType    *string
	Timeout      *string
}

func (p EdiscoveryOcrSettingsOperationPredicate) Matches(input EdiscoveryOcrSettings) bool {

	if p.IsEnabled != nil && (input.IsEnabled == nil || *p.IsEnabled != *input.IsEnabled) {
		return false
	}

	if p.MaxImageSize != nil && (input.MaxImageSize == nil || *p.MaxImageSize != *input.MaxImageSize) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Timeout != nil && (input.Timeout == nil || *p.Timeout != *input.Timeout) {
		return false
	}

	return true
}
