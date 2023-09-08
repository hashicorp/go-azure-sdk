package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSettingGroupCollectionDefinitionSettingUsage string

const (
	DeviceManagementConfigurationSettingGroupCollectionDefinitionSettingUsagecompliance    DeviceManagementConfigurationSettingGroupCollectionDefinitionSettingUsage = "Compliance"
	DeviceManagementConfigurationSettingGroupCollectionDefinitionSettingUsageconfiguration DeviceManagementConfigurationSettingGroupCollectionDefinitionSettingUsage = "Configuration"
	DeviceManagementConfigurationSettingGroupCollectionDefinitionSettingUsagenone          DeviceManagementConfigurationSettingGroupCollectionDefinitionSettingUsage = "None"
)

func PossibleValuesForDeviceManagementConfigurationSettingGroupCollectionDefinitionSettingUsage() []string {
	return []string{
		string(DeviceManagementConfigurationSettingGroupCollectionDefinitionSettingUsagecompliance),
		string(DeviceManagementConfigurationSettingGroupCollectionDefinitionSettingUsageconfiguration),
		string(DeviceManagementConfigurationSettingGroupCollectionDefinitionSettingUsagenone),
	}
}

func parseDeviceManagementConfigurationSettingGroupCollectionDefinitionSettingUsage(input string) (*DeviceManagementConfigurationSettingGroupCollectionDefinitionSettingUsage, error) {
	vals := map[string]DeviceManagementConfigurationSettingGroupCollectionDefinitionSettingUsage{
		"compliance":    DeviceManagementConfigurationSettingGroupCollectionDefinitionSettingUsagecompliance,
		"configuration": DeviceManagementConfigurationSettingGroupCollectionDefinitionSettingUsageconfiguration,
		"none":          DeviceManagementConfigurationSettingGroupCollectionDefinitionSettingUsagenone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationSettingGroupCollectionDefinitionSettingUsage(input)
	return &out, nil
}
