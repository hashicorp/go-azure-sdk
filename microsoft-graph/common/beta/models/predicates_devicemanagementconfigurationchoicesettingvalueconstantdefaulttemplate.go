package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationChoiceSettingValueConstantDefaultTemplateOperationPredicate struct {
	ODataType                 *string
	SettingDefinitionOptionId *string
}

func (p DeviceManagementConfigurationChoiceSettingValueConstantDefaultTemplateOperationPredicate) Matches(input DeviceManagementConfigurationChoiceSettingValueConstantDefaultTemplate) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.SettingDefinitionOptionId != nil && (input.SettingDefinitionOptionId == nil || *p.SettingDefinitionOptionId != *input.SettingDefinitionOptionId) {
		return false
	}

	return true
}
