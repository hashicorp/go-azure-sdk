package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type B2cIdentityUserFlowOperationPredicate struct {
	DefaultLanguageTag             *string
	Id                             *string
	IsLanguageCustomizationEnabled *bool
	ODataType                      *string
}

func (p B2cIdentityUserFlowOperationPredicate) Matches(input B2cIdentityUserFlow) bool {

	if p.DefaultLanguageTag != nil && (input.DefaultLanguageTag == nil || *p.DefaultLanguageTag != *input.DefaultLanguageTag) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsLanguageCustomizationEnabled != nil && (input.IsLanguageCustomizationEnabled == nil || *p.IsLanguageCustomizationEnabled != *input.IsLanguageCustomizationEnabled) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
