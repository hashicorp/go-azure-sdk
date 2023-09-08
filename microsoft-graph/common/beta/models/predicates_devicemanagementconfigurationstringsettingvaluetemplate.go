package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationStringSettingValueTemplateOperationPredicate struct {
	ODataType              *string
	SettingValueTemplateId *string
}

func (p DeviceManagementConfigurationStringSettingValueTemplateOperationPredicate) Matches(input DeviceManagementConfigurationStringSettingValueTemplate) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.SettingValueTemplateId != nil && (input.SettingValueTemplateId == nil || *p.SettingValueTemplateId != *input.SettingValueTemplateId) {
		return false
	}

	return true
}
