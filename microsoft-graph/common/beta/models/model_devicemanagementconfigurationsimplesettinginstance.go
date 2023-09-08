package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSimpleSettingInstance struct {
	ODataType                        *string                                                        `json:"@odata.type,omitempty"`
	SettingDefinitionId              *string                                                        `json:"settingDefinitionId,omitempty"`
	SettingInstanceTemplateReference *DeviceManagementConfigurationSettingInstanceTemplateReference `json:"settingInstanceTemplateReference,omitempty"`
	SimpleSettingValue               *DeviceManagementConfigurationSimpleSettingValue               `json:"simpleSettingValue,omitempty"`
}
