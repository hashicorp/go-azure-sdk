package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationGroupSettingInstanceTemplate struct {
	GroupSettingValueTemplate *DeviceManagementConfigurationGroupSettingValueTemplate `json:"groupSettingValueTemplate,omitempty"`
	IsRequired                *bool                                                   `json:"isRequired,omitempty"`
	ODataType                 *string                                                 `json:"@odata.type,omitempty"`
	SettingDefinitionId       *string                                                 `json:"settingDefinitionId,omitempty"`
	SettingInstanceTemplateId *string                                                 `json:"settingInstanceTemplateId,omitempty"`
}
