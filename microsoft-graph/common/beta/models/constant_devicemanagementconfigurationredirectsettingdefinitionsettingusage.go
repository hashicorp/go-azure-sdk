package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationRedirectSettingDefinitionSettingUsage string

const (
	DeviceManagementConfigurationRedirectSettingDefinitionSettingUsagecompliance    DeviceManagementConfigurationRedirectSettingDefinitionSettingUsage = "Compliance"
	DeviceManagementConfigurationRedirectSettingDefinitionSettingUsageconfiguration DeviceManagementConfigurationRedirectSettingDefinitionSettingUsage = "Configuration"
	DeviceManagementConfigurationRedirectSettingDefinitionSettingUsagenone          DeviceManagementConfigurationRedirectSettingDefinitionSettingUsage = "None"
)

func PossibleValuesForDeviceManagementConfigurationRedirectSettingDefinitionSettingUsage() []string {
	return []string{
		string(DeviceManagementConfigurationRedirectSettingDefinitionSettingUsagecompliance),
		string(DeviceManagementConfigurationRedirectSettingDefinitionSettingUsageconfiguration),
		string(DeviceManagementConfigurationRedirectSettingDefinitionSettingUsagenone),
	}
}

func parseDeviceManagementConfigurationRedirectSettingDefinitionSettingUsage(input string) (*DeviceManagementConfigurationRedirectSettingDefinitionSettingUsage, error) {
	vals := map[string]DeviceManagementConfigurationRedirectSettingDefinitionSettingUsage{
		"compliance":    DeviceManagementConfigurationRedirectSettingDefinitionSettingUsagecompliance,
		"configuration": DeviceManagementConfigurationRedirectSettingDefinitionSettingUsageconfiguration,
		"none":          DeviceManagementConfigurationRedirectSettingDefinitionSettingUsagenone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationRedirectSettingDefinitionSettingUsage(input)
	return &out, nil
}
