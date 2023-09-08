package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserFlowLanguageConfigurationOperationPredicate struct {
	DisplayName *string
	Id          *string
	IsEnabled   *bool
	ODataType   *string
}

func (p UserFlowLanguageConfigurationOperationPredicate) Matches(input UserFlowLanguageConfiguration) bool {

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsEnabled != nil && (input.IsEnabled == nil || *p.IsEnabled != *input.IsEnabled) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
