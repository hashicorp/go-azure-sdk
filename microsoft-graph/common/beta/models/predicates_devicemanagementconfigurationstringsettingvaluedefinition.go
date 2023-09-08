package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationStringSettingValueDefinitionOperationPredicate struct {
	InputValidationSchema *string
	IsSecret              *bool
	MaximumLength         *int64
	MinimumLength         *int64
	ODataType             *string
}

func (p DeviceManagementConfigurationStringSettingValueDefinitionOperationPredicate) Matches(input DeviceManagementConfigurationStringSettingValueDefinition) bool {

	if p.InputValidationSchema != nil && (input.InputValidationSchema == nil || *p.InputValidationSchema != *input.InputValidationSchema) {
		return false
	}

	if p.IsSecret != nil && (input.IsSecret == nil || *p.IsSecret != *input.IsSecret) {
		return false
	}

	if p.MaximumLength != nil && (input.MaximumLength == nil || *p.MaximumLength != *input.MaximumLength) {
		return false
	}

	if p.MinimumLength != nil && (input.MinimumLength == nil || *p.MinimumLength != *input.MinimumLength) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
