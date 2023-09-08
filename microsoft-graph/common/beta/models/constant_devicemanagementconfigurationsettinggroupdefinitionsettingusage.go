package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSettingGroupDefinitionSettingUsage string

const (
	DeviceManagementConfigurationSettingGroupDefinitionSettingUsagecompliance    DeviceManagementConfigurationSettingGroupDefinitionSettingUsage = "Compliance"
	DeviceManagementConfigurationSettingGroupDefinitionSettingUsageconfiguration DeviceManagementConfigurationSettingGroupDefinitionSettingUsage = "Configuration"
	DeviceManagementConfigurationSettingGroupDefinitionSettingUsagenone          DeviceManagementConfigurationSettingGroupDefinitionSettingUsage = "None"
)

func PossibleValuesForDeviceManagementConfigurationSettingGroupDefinitionSettingUsage() []string {
	return []string{
		string(DeviceManagementConfigurationSettingGroupDefinitionSettingUsagecompliance),
		string(DeviceManagementConfigurationSettingGroupDefinitionSettingUsageconfiguration),
		string(DeviceManagementConfigurationSettingGroupDefinitionSettingUsagenone),
	}
}

func parseDeviceManagementConfigurationSettingGroupDefinitionSettingUsage(input string) (*DeviceManagementConfigurationSettingGroupDefinitionSettingUsage, error) {
	vals := map[string]DeviceManagementConfigurationSettingGroupDefinitionSettingUsage{
		"compliance":    DeviceManagementConfigurationSettingGroupDefinitionSettingUsagecompliance,
		"configuration": DeviceManagementConfigurationSettingGroupDefinitionSettingUsageconfiguration,
		"none":          DeviceManagementConfigurationSettingGroupDefinitionSettingUsagenone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationSettingGroupDefinitionSettingUsage(input)
	return &out, nil
}
