package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationCategorySettingUsage string

const (
	DeviceManagementConfigurationCategorySettingUsagecompliance    DeviceManagementConfigurationCategorySettingUsage = "Compliance"
	DeviceManagementConfigurationCategorySettingUsageconfiguration DeviceManagementConfigurationCategorySettingUsage = "Configuration"
	DeviceManagementConfigurationCategorySettingUsagenone          DeviceManagementConfigurationCategorySettingUsage = "None"
)

func PossibleValuesForDeviceManagementConfigurationCategorySettingUsage() []string {
	return []string{
		string(DeviceManagementConfigurationCategorySettingUsagecompliance),
		string(DeviceManagementConfigurationCategorySettingUsageconfiguration),
		string(DeviceManagementConfigurationCategorySettingUsagenone),
	}
}

func parseDeviceManagementConfigurationCategorySettingUsage(input string) (*DeviceManagementConfigurationCategorySettingUsage, error) {
	vals := map[string]DeviceManagementConfigurationCategorySettingUsage{
		"compliance":    DeviceManagementConfigurationCategorySettingUsagecompliance,
		"configuration": DeviceManagementConfigurationCategorySettingUsageconfiguration,
		"none":          DeviceManagementConfigurationCategorySettingUsagenone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationCategorySettingUsage(input)
	return &out, nil
}
