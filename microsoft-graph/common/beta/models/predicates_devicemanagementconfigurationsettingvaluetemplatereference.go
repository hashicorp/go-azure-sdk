package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSettingValueTemplateReferenceOperationPredicate struct {
	ODataType              *string
	SettingValueTemplateId *string
	UseTemplateDefault     *bool
}

func (p DeviceManagementConfigurationSettingValueTemplateReferenceOperationPredicate) Matches(input DeviceManagementConfigurationSettingValueTemplateReference) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.SettingValueTemplateId != nil && (input.SettingValueTemplateId == nil || *p.SettingValueTemplateId != *input.SettingValueTemplateId) {
		return false
	}

	if p.UseTemplateDefault != nil && (input.UseTemplateDefault == nil || *p.UseTemplateDefault != *input.UseTemplateDefault) {
		return false
	}

	return true
}
