package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationChoiceSettingValueTemplate struct {
	DefaultValue               *DeviceManagementConfigurationChoiceSettingValueDefaultTemplate    `json:"defaultValue,omitempty"`
	ODataType                  *string                                                            `json:"@odata.type,omitempty"`
	RecommendedValueDefinition *DeviceManagementConfigurationChoiceSettingValueDefinitionTemplate `json:"recommendedValueDefinition,omitempty"`
	RequiredValueDefinition    *DeviceManagementConfigurationChoiceSettingValueDefinitionTemplate `json:"requiredValueDefinition,omitempty"`
	SettingValueTemplateId     *string                                                            `json:"settingValueTemplateId,omitempty"`
}
