package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSettingInstanceOperationPredicate struct {
	ODataType           *string
	SettingDefinitionId *string
}

func (p DeviceManagementConfigurationSettingInstanceOperationPredicate) Matches(input DeviceManagementConfigurationSettingInstance) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.SettingDefinitionId != nil && (input.SettingDefinitionId == nil || *p.SettingDefinitionId != *input.SettingDefinitionId) {
		return false
	}

	return true
}
