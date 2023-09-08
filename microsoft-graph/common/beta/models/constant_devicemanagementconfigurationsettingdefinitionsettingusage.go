package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSettingDefinitionSettingUsage string

const (
	DeviceManagementConfigurationSettingDefinitionSettingUsagecompliance    DeviceManagementConfigurationSettingDefinitionSettingUsage = "Compliance"
	DeviceManagementConfigurationSettingDefinitionSettingUsageconfiguration DeviceManagementConfigurationSettingDefinitionSettingUsage = "Configuration"
	DeviceManagementConfigurationSettingDefinitionSettingUsagenone          DeviceManagementConfigurationSettingDefinitionSettingUsage = "None"
)

func PossibleValuesForDeviceManagementConfigurationSettingDefinitionSettingUsage() []string {
	return []string{
		string(DeviceManagementConfigurationSettingDefinitionSettingUsagecompliance),
		string(DeviceManagementConfigurationSettingDefinitionSettingUsageconfiguration),
		string(DeviceManagementConfigurationSettingDefinitionSettingUsagenone),
	}
}

func parseDeviceManagementConfigurationSettingDefinitionSettingUsage(input string) (*DeviceManagementConfigurationSettingDefinitionSettingUsage, error) {
	vals := map[string]DeviceManagementConfigurationSettingDefinitionSettingUsage{
		"compliance":    DeviceManagementConfigurationSettingDefinitionSettingUsagecompliance,
		"configuration": DeviceManagementConfigurationSettingDefinitionSettingUsageconfiguration,
		"none":          DeviceManagementConfigurationSettingDefinitionSettingUsagenone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationSettingDefinitionSettingUsage(input)
	return &out, nil
}
