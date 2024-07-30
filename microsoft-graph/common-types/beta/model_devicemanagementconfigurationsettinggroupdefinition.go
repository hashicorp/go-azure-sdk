package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSettingGroupDefinition struct {
	AccessTypes                    *DeviceManagementConfigurationSettingGroupDefinitionAccessTypes  `json:"accessTypes,omitempty"`
	Applicability                  *DeviceManagementConfigurationSettingApplicability               `json:"applicability,omitempty"`
	BaseUri                        *string                                                          `json:"baseUri,omitempty"`
	CategoryId                     *string                                                          `json:"categoryId,omitempty"`
	ChildIds                       *[]string                                                        `json:"childIds,omitempty"`
	DependedOnBy                   *[]DeviceManagementConfigurationSettingDependedOnBy              `json:"dependedOnBy,omitempty"`
	DependentOn                    *[]DeviceManagementConfigurationDependentOn                      `json:"dependentOn,omitempty"`
	Description                    *string                                                          `json:"description,omitempty"`
	DisplayName                    *string                                                          `json:"displayName,omitempty"`
	HelpText                       *string                                                          `json:"helpText,omitempty"`
	Id                             *string                                                          `json:"id,omitempty"`
	InfoUrls                       *[]string                                                        `json:"infoUrls,omitempty"`
	Keywords                       *[]string                                                        `json:"keywords,omitempty"`
	Name                           *string                                                          `json:"name,omitempty"`
	ODataType                      *string                                                          `json:"@odata.type,omitempty"`
	Occurrence                     *DeviceManagementConfigurationSettingOccurrence                  `json:"occurrence,omitempty"`
	OffsetUri                      *string                                                          `json:"offsetUri,omitempty"`
	ReferredSettingInformationList *[]DeviceManagementConfigurationReferredSettingInformation       `json:"referredSettingInformationList,omitempty"`
	RootDefinitionId               *string                                                          `json:"rootDefinitionId,omitempty"`
	SettingUsage                   *DeviceManagementConfigurationSettingGroupDefinitionSettingUsage `json:"settingUsage,omitempty"`
	UxBehavior                     *DeviceManagementConfigurationSettingGroupDefinitionUxBehavior   `json:"uxBehavior,omitempty"`
	Version                        *string                                                          `json:"version,omitempty"`
	Visibility                     *DeviceManagementConfigurationSettingGroupDefinitionVisibility   `json:"visibility,omitempty"`
}
