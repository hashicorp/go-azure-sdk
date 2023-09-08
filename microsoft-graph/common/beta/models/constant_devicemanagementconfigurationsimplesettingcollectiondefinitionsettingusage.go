package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSimpleSettingCollectionDefinitionSettingUsage string

const (
	DeviceManagementConfigurationSimpleSettingCollectionDefinitionSettingUsagecompliance    DeviceManagementConfigurationSimpleSettingCollectionDefinitionSettingUsage = "Compliance"
	DeviceManagementConfigurationSimpleSettingCollectionDefinitionSettingUsageconfiguration DeviceManagementConfigurationSimpleSettingCollectionDefinitionSettingUsage = "Configuration"
	DeviceManagementConfigurationSimpleSettingCollectionDefinitionSettingUsagenone          DeviceManagementConfigurationSimpleSettingCollectionDefinitionSettingUsage = "None"
)

func PossibleValuesForDeviceManagementConfigurationSimpleSettingCollectionDefinitionSettingUsage() []string {
	return []string{
		string(DeviceManagementConfigurationSimpleSettingCollectionDefinitionSettingUsagecompliance),
		string(DeviceManagementConfigurationSimpleSettingCollectionDefinitionSettingUsageconfiguration),
		string(DeviceManagementConfigurationSimpleSettingCollectionDefinitionSettingUsagenone),
	}
}

func parseDeviceManagementConfigurationSimpleSettingCollectionDefinitionSettingUsage(input string) (*DeviceManagementConfigurationSimpleSettingCollectionDefinitionSettingUsage, error) {
	vals := map[string]DeviceManagementConfigurationSimpleSettingCollectionDefinitionSettingUsage{
		"compliance":    DeviceManagementConfigurationSimpleSettingCollectionDefinitionSettingUsagecompliance,
		"configuration": DeviceManagementConfigurationSimpleSettingCollectionDefinitionSettingUsageconfiguration,
		"none":          DeviceManagementConfigurationSimpleSettingCollectionDefinitionSettingUsagenone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationSimpleSettingCollectionDefinitionSettingUsage(input)
	return &out, nil
}
