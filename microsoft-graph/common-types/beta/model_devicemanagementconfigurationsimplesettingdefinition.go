package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSimpleSettingDefinition struct {
	AccessTypes                    *DeviceManagementConfigurationSimpleSettingDefinitionAccessTypes  `json:"accessTypes,omitempty"`
	Applicability                  *DeviceManagementConfigurationSettingApplicability                `json:"applicability,omitempty"`
	BaseUri                        *string                                                           `json:"baseUri,omitempty"`
	CategoryId                     *string                                                           `json:"categoryId,omitempty"`
	DefaultValue                   *DeviceManagementConfigurationSettingValue                        `json:"defaultValue,omitempty"`
	DependedOnBy                   *[]DeviceManagementConfigurationSettingDependedOnBy               `json:"dependedOnBy,omitempty"`
	DependentOn                    *[]DeviceManagementConfigurationDependentOn                       `json:"dependentOn,omitempty"`
	Description                    *string                                                           `json:"description,omitempty"`
	DisplayName                    *string                                                           `json:"displayName,omitempty"`
	HelpText                       *string                                                           `json:"helpText,omitempty"`
	Id                             *string                                                           `json:"id,omitempty"`
	InfoUrls                       *[]string                                                         `json:"infoUrls,omitempty"`
	Keywords                       *[]string                                                         `json:"keywords,omitempty"`
	Name                           *string                                                           `json:"name,omitempty"`
	ODataType                      *string                                                           `json:"@odata.type,omitempty"`
	Occurrence                     *DeviceManagementConfigurationSettingOccurrence                   `json:"occurrence,omitempty"`
	OffsetUri                      *string                                                           `json:"offsetUri,omitempty"`
	ReferredSettingInformationList *[]DeviceManagementConfigurationReferredSettingInformation        `json:"referredSettingInformationList,omitempty"`
	RootDefinitionId               *string                                                           `json:"rootDefinitionId,omitempty"`
	SettingUsage                   *DeviceManagementConfigurationSimpleSettingDefinitionSettingUsage `json:"settingUsage,omitempty"`
	UxBehavior                     *DeviceManagementConfigurationSimpleSettingDefinitionUxBehavior   `json:"uxBehavior,omitempty"`
	ValueDefinition                *DeviceManagementConfigurationSettingValueDefinition              `json:"valueDefinition,omitempty"`
	Version                        *string                                                           `json:"version,omitempty"`
	Visibility                     *DeviceManagementConfigurationSimpleSettingDefinitionVisibility   `json:"visibility,omitempty"`
}
