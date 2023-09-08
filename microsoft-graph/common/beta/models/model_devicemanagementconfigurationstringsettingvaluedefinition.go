package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationStringSettingValueDefinition struct {
	FileTypes             *[]string                                                        `json:"fileTypes,omitempty"`
	Format                *DeviceManagementConfigurationStringSettingValueDefinitionFormat `json:"format,omitempty"`
	InputValidationSchema *string                                                          `json:"inputValidationSchema,omitempty"`
	IsSecret              *bool                                                            `json:"isSecret,omitempty"`
	MaximumLength         *int64                                                           `json:"maximumLength,omitempty"`
	MinimumLength         *int64                                                           `json:"minimumLength,omitempty"`
	ODataType             *string                                                          `json:"@odata.type,omitempty"`
}
