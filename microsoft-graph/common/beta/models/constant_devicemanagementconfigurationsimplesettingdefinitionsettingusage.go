package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSimpleSettingDefinitionSettingUsage string

const (
	DeviceManagementConfigurationSimpleSettingDefinitionSettingUsagecompliance    DeviceManagementConfigurationSimpleSettingDefinitionSettingUsage = "Compliance"
	DeviceManagementConfigurationSimpleSettingDefinitionSettingUsageconfiguration DeviceManagementConfigurationSimpleSettingDefinitionSettingUsage = "Configuration"
	DeviceManagementConfigurationSimpleSettingDefinitionSettingUsagenone          DeviceManagementConfigurationSimpleSettingDefinitionSettingUsage = "None"
)

func PossibleValuesForDeviceManagementConfigurationSimpleSettingDefinitionSettingUsage() []string {
	return []string{
		string(DeviceManagementConfigurationSimpleSettingDefinitionSettingUsagecompliance),
		string(DeviceManagementConfigurationSimpleSettingDefinitionSettingUsageconfiguration),
		string(DeviceManagementConfigurationSimpleSettingDefinitionSettingUsagenone),
	}
}

func parseDeviceManagementConfigurationSimpleSettingDefinitionSettingUsage(input string) (*DeviceManagementConfigurationSimpleSettingDefinitionSettingUsage, error) {
	vals := map[string]DeviceManagementConfigurationSimpleSettingDefinitionSettingUsage{
		"compliance":    DeviceManagementConfigurationSimpleSettingDefinitionSettingUsagecompliance,
		"configuration": DeviceManagementConfigurationSimpleSettingDefinitionSettingUsageconfiguration,
		"none":          DeviceManagementConfigurationSimpleSettingDefinitionSettingUsagenone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationSimpleSettingDefinitionSettingUsage(input)
	return &out, nil
}
