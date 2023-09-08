package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationChoiceSettingCollectionInstance struct {
	ChoiceSettingCollectionValue     *[]DeviceManagementConfigurationChoiceSettingValue             `json:"choiceSettingCollectionValue,omitempty"`
	ODataType                        *string                                                        `json:"@odata.type,omitempty"`
	SettingDefinitionId              *string                                                        `json:"settingDefinitionId,omitempty"`
	SettingInstanceTemplateReference *DeviceManagementConfigurationSettingInstanceTemplateReference `json:"settingInstanceTemplateReference,omitempty"`
}
