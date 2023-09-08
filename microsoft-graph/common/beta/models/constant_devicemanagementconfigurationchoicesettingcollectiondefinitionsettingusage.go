package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationChoiceSettingCollectionDefinitionSettingUsage string

const (
	DeviceManagementConfigurationChoiceSettingCollectionDefinitionSettingUsagecompliance    DeviceManagementConfigurationChoiceSettingCollectionDefinitionSettingUsage = "Compliance"
	DeviceManagementConfigurationChoiceSettingCollectionDefinitionSettingUsageconfiguration DeviceManagementConfigurationChoiceSettingCollectionDefinitionSettingUsage = "Configuration"
	DeviceManagementConfigurationChoiceSettingCollectionDefinitionSettingUsagenone          DeviceManagementConfigurationChoiceSettingCollectionDefinitionSettingUsage = "None"
)

func PossibleValuesForDeviceManagementConfigurationChoiceSettingCollectionDefinitionSettingUsage() []string {
	return []string{
		string(DeviceManagementConfigurationChoiceSettingCollectionDefinitionSettingUsagecompliance),
		string(DeviceManagementConfigurationChoiceSettingCollectionDefinitionSettingUsageconfiguration),
		string(DeviceManagementConfigurationChoiceSettingCollectionDefinitionSettingUsagenone),
	}
}

func parseDeviceManagementConfigurationChoiceSettingCollectionDefinitionSettingUsage(input string) (*DeviceManagementConfigurationChoiceSettingCollectionDefinitionSettingUsage, error) {
	vals := map[string]DeviceManagementConfigurationChoiceSettingCollectionDefinitionSettingUsage{
		"compliance":    DeviceManagementConfigurationChoiceSettingCollectionDefinitionSettingUsagecompliance,
		"configuration": DeviceManagementConfigurationChoiceSettingCollectionDefinitionSettingUsageconfiguration,
		"none":          DeviceManagementConfigurationChoiceSettingCollectionDefinitionSettingUsagenone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationChoiceSettingCollectionDefinitionSettingUsage(input)
	return &out, nil
}
