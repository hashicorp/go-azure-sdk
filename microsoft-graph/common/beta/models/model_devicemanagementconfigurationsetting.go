package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSetting struct {
	Id                 *string                                           `json:"id,omitempty"`
	ODataType          *string                                           `json:"@odata.type,omitempty"`
	SettingDefinitions *[]DeviceManagementConfigurationSettingDefinition `json:"settingDefinitions,omitempty"`
	SettingInstance    *DeviceManagementConfigurationSettingInstance     `json:"settingInstance,omitempty"`
}
