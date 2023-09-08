package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSimpleSettingInstanceTemplateOperationPredicate struct {
	IsRequired                *bool
	ODataType                 *string
	SettingDefinitionId       *string
	SettingInstanceTemplateId *string
}

func (p DeviceManagementConfigurationSimpleSettingInstanceTemplateOperationPredicate) Matches(input DeviceManagementConfigurationSimpleSettingInstanceTemplate) bool {

	if p.IsRequired != nil && (input.IsRequired == nil || *p.IsRequired != *input.IsRequired) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.SettingDefinitionId != nil && (input.SettingDefinitionId == nil || *p.SettingDefinitionId != *input.SettingDefinitionId) {
		return false
	}

	if p.SettingInstanceTemplateId != nil && (input.SettingInstanceTemplateId == nil || *p.SettingInstanceTemplateId != *input.SettingInstanceTemplateId) {
		return false
	}

	return true
}
