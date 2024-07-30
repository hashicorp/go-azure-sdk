package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationChoiceSettingInstance struct {
	ChoiceSettingValue               *DeviceManagementConfigurationChoiceSettingValue               `json:"choiceSettingValue,omitempty"`
	ODataType                        *string                                                        `json:"@odata.type,omitempty"`
	SettingDefinitionId              *string                                                        `json:"settingDefinitionId,omitempty"`
	SettingInstanceTemplateReference *DeviceManagementConfigurationSettingInstanceTemplateReference `json:"settingInstanceTemplateReference,omitempty"`
}
