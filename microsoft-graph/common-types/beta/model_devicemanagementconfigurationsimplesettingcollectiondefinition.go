package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSimpleSettingCollectionDefinition struct {
	AccessTypes                    *DeviceManagementConfigurationSimpleSettingCollectionDefinitionAccessTypes  `json:"accessTypes,omitempty"`
	Applicability                  *DeviceManagementConfigurationSettingApplicability                          `json:"applicability,omitempty"`
	BaseUri                        *string                                                                     `json:"baseUri,omitempty"`
	CategoryId                     *string                                                                     `json:"categoryId,omitempty"`
	DefaultValue                   *DeviceManagementConfigurationSettingValue                                  `json:"defaultValue,omitempty"`
	DependedOnBy                   *[]DeviceManagementConfigurationSettingDependedOnBy                         `json:"dependedOnBy,omitempty"`
	DependentOn                    *[]DeviceManagementConfigurationDependentOn                                 `json:"dependentOn,omitempty"`
	Description                    *string                                                                     `json:"description,omitempty"`
	DisplayName                    *string                                                                     `json:"displayName,omitempty"`
	HelpText                       *string                                                                     `json:"helpText,omitempty"`
	Id                             *string                                                                     `json:"id,omitempty"`
	InfoUrls                       *[]string                                                                   `json:"infoUrls,omitempty"`
	Keywords                       *[]string                                                                   `json:"keywords,omitempty"`
	MaximumCount                   *int64                                                                      `json:"maximumCount,omitempty"`
	MinimumCount                   *int64                                                                      `json:"minimumCount,omitempty"`
	Name                           *string                                                                     `json:"name,omitempty"`
	ODataType                      *string                                                                     `json:"@odata.type,omitempty"`
	Occurrence                     *DeviceManagementConfigurationSettingOccurrence                             `json:"occurrence,omitempty"`
	OffsetUri                      *string                                                                     `json:"offsetUri,omitempty"`
	ReferredSettingInformationList *[]DeviceManagementConfigurationReferredSettingInformation                  `json:"referredSettingInformationList,omitempty"`
	RootDefinitionId               *string                                                                     `json:"rootDefinitionId,omitempty"`
	SettingUsage                   *DeviceManagementConfigurationSimpleSettingCollectionDefinitionSettingUsage `json:"settingUsage,omitempty"`
	UxBehavior                     *DeviceManagementConfigurationSimpleSettingCollectionDefinitionUxBehavior   `json:"uxBehavior,omitempty"`
	ValueDefinition                *DeviceManagementConfigurationSettingValueDefinition                        `json:"valueDefinition,omitempty"`
	Version                        *string                                                                     `json:"version,omitempty"`
	Visibility                     *DeviceManagementConfigurationSimpleSettingCollectionDefinitionVisibility   `json:"visibility,omitempty"`
}
