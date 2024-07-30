package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityBaselineTemplate struct {
	Categories                   *[]DeviceManagementTemplateSettingCategory `json:"categories,omitempty"`
	CategoryDeviceStateSummaries *[]SecurityBaselineCategoryStateSummary    `json:"categoryDeviceStateSummaries,omitempty"`
	Description                  *string                                    `json:"description,omitempty"`
	DeviceStateSummary           *SecurityBaselineStateSummary              `json:"deviceStateSummary,omitempty"`
	DeviceStates                 *[]SecurityBaselineDeviceState             `json:"deviceStates,omitempty"`
	DisplayName                  *string                                    `json:"displayName,omitempty"`
	Id                           *string                                    `json:"id,omitempty"`
	IntentCount                  *int64                                     `json:"intentCount,omitempty"`
	IsDeprecated                 *bool                                      `json:"isDeprecated,omitempty"`
	MigratableTo                 *[]DeviceManagementTemplate                `json:"migratableTo,omitempty"`
	ODataType                    *string                                    `json:"@odata.type,omitempty"`
	PlatformType                 *SecurityBaselineTemplatePlatformType      `json:"platformType,omitempty"`
	PublishedDateTime            *string                                    `json:"publishedDateTime,omitempty"`
	Settings                     *[]DeviceManagementSettingInstance         `json:"settings,omitempty"`
	TemplateSubtype              *SecurityBaselineTemplateTemplateSubtype   `json:"templateSubtype,omitempty"`
	TemplateType                 *SecurityBaselineTemplateTemplateType      `json:"templateType,omitempty"`
	VersionInfo                  *string                                    `json:"versionInfo,omitempty"`
}
