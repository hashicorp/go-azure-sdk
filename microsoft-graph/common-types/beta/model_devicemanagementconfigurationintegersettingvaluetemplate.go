package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationIntegerSettingValueTemplate struct {
	DefaultValue               *DeviceManagementConfigurationIntegerSettingValueDefaultTemplate    `json:"defaultValue,omitempty"`
	ODataType                  *string                                                             `json:"@odata.type,omitempty"`
	RecommendedValueDefinition *DeviceManagementConfigurationIntegerSettingValueDefinitionTemplate `json:"recommendedValueDefinition,omitempty"`
	RequiredValueDefinition    *DeviceManagementConfigurationIntegerSettingValueDefinitionTemplate `json:"requiredValueDefinition,omitempty"`
	SettingValueTemplateId     *string                                                             `json:"settingValueTemplateId,omitempty"`
}
