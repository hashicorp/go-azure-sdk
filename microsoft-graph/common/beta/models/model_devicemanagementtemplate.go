package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementTemplate struct {
	Categories        *[]DeviceManagementTemplateSettingCategory `json:"categories,omitempty"`
	Description       *string                                    `json:"description,omitempty"`
	DisplayName       *string                                    `json:"displayName,omitempty"`
	Id                *string                                    `json:"id,omitempty"`
	IntentCount       *int64                                     `json:"intentCount,omitempty"`
	IsDeprecated      *bool                                      `json:"isDeprecated,omitempty"`
	MigratableTo      *[]DeviceManagementTemplate                `json:"migratableTo,omitempty"`
	ODataType         *string                                    `json:"@odata.type,omitempty"`
	PlatformType      *DeviceManagementTemplatePlatformType      `json:"platformType,omitempty"`
	PublishedDateTime *string                                    `json:"publishedDateTime,omitempty"`
	Settings          *[]DeviceManagementSettingInstance         `json:"settings,omitempty"`
	TemplateSubtype   *DeviceManagementTemplateTemplateSubtype   `json:"templateSubtype,omitempty"`
	TemplateType      *DeviceManagementTemplateTemplateType      `json:"templateType,omitempty"`
	VersionInfo       *string                                    `json:"versionInfo,omitempty"`
}
