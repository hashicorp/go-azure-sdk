package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationOptionDefinition struct {
	DependedOnBy *[]DeviceManagementConfigurationSettingDependedOnBy `json:"dependedOnBy,omitempty"`
	DependentOn  *[]DeviceManagementConfigurationDependentOn         `json:"dependentOn,omitempty"`
	Description  *string                                             `json:"description,omitempty"`
	DisplayName  *string                                             `json:"displayName,omitempty"`
	HelpText     *string                                             `json:"helpText,omitempty"`
	ItemId       *string                                             `json:"itemId,omitempty"`
	Name         *string                                             `json:"name,omitempty"`
	ODataType    *string                                             `json:"@odata.type,omitempty"`
	OptionValue  *DeviceManagementConfigurationSettingValue          `json:"optionValue,omitempty"`
}
