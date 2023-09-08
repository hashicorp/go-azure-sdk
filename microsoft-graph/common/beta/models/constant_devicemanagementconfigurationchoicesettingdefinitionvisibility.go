package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationChoiceSettingDefinitionVisibility string

const (
	DeviceManagementConfigurationChoiceSettingDefinitionVisibilitynone            DeviceManagementConfigurationChoiceSettingDefinitionVisibility = "None"
	DeviceManagementConfigurationChoiceSettingDefinitionVisibilitysettingsCatalog DeviceManagementConfigurationChoiceSettingDefinitionVisibility = "SettingsCatalog"
	DeviceManagementConfigurationChoiceSettingDefinitionVisibilitytemplate        DeviceManagementConfigurationChoiceSettingDefinitionVisibility = "Template"
)

func PossibleValuesForDeviceManagementConfigurationChoiceSettingDefinitionVisibility() []string {
	return []string{
		string(DeviceManagementConfigurationChoiceSettingDefinitionVisibilitynone),
		string(DeviceManagementConfigurationChoiceSettingDefinitionVisibilitysettingsCatalog),
		string(DeviceManagementConfigurationChoiceSettingDefinitionVisibilitytemplate),
	}
}

func parseDeviceManagementConfigurationChoiceSettingDefinitionVisibility(input string) (*DeviceManagementConfigurationChoiceSettingDefinitionVisibility, error) {
	vals := map[string]DeviceManagementConfigurationChoiceSettingDefinitionVisibility{
		"none":            DeviceManagementConfigurationChoiceSettingDefinitionVisibilitynone,
		"settingscatalog": DeviceManagementConfigurationChoiceSettingDefinitionVisibilitysettingsCatalog,
		"template":        DeviceManagementConfigurationChoiceSettingDefinitionVisibilitytemplate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationChoiceSettingDefinitionVisibility(input)
	return &out, nil
}
