package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationChoiceSettingDefinitionSettingUsage string

const (
	DeviceManagementConfigurationChoiceSettingDefinitionSettingUsagecompliance    DeviceManagementConfigurationChoiceSettingDefinitionSettingUsage = "Compliance"
	DeviceManagementConfigurationChoiceSettingDefinitionSettingUsageconfiguration DeviceManagementConfigurationChoiceSettingDefinitionSettingUsage = "Configuration"
	DeviceManagementConfigurationChoiceSettingDefinitionSettingUsagenone          DeviceManagementConfigurationChoiceSettingDefinitionSettingUsage = "None"
)

func PossibleValuesForDeviceManagementConfigurationChoiceSettingDefinitionSettingUsage() []string {
	return []string{
		string(DeviceManagementConfigurationChoiceSettingDefinitionSettingUsagecompliance),
		string(DeviceManagementConfigurationChoiceSettingDefinitionSettingUsageconfiguration),
		string(DeviceManagementConfigurationChoiceSettingDefinitionSettingUsagenone),
	}
}

func parseDeviceManagementConfigurationChoiceSettingDefinitionSettingUsage(input string) (*DeviceManagementConfigurationChoiceSettingDefinitionSettingUsage, error) {
	vals := map[string]DeviceManagementConfigurationChoiceSettingDefinitionSettingUsage{
		"compliance":    DeviceManagementConfigurationChoiceSettingDefinitionSettingUsagecompliance,
		"configuration": DeviceManagementConfigurationChoiceSettingDefinitionSettingUsageconfiguration,
		"none":          DeviceManagementConfigurationChoiceSettingDefinitionSettingUsagenone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationChoiceSettingDefinitionSettingUsage(input)
	return &out, nil
}
