package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSettingDefinitionVisibility string

const (
	DeviceManagementConfigurationSettingDefinitionVisibilitynone            DeviceManagementConfigurationSettingDefinitionVisibility = "None"
	DeviceManagementConfigurationSettingDefinitionVisibilitysettingsCatalog DeviceManagementConfigurationSettingDefinitionVisibility = "SettingsCatalog"
	DeviceManagementConfigurationSettingDefinitionVisibilitytemplate        DeviceManagementConfigurationSettingDefinitionVisibility = "Template"
)

func PossibleValuesForDeviceManagementConfigurationSettingDefinitionVisibility() []string {
	return []string{
		string(DeviceManagementConfigurationSettingDefinitionVisibilitynone),
		string(DeviceManagementConfigurationSettingDefinitionVisibilitysettingsCatalog),
		string(DeviceManagementConfigurationSettingDefinitionVisibilitytemplate),
	}
}

func parseDeviceManagementConfigurationSettingDefinitionVisibility(input string) (*DeviceManagementConfigurationSettingDefinitionVisibility, error) {
	vals := map[string]DeviceManagementConfigurationSettingDefinitionVisibility{
		"none":            DeviceManagementConfigurationSettingDefinitionVisibilitynone,
		"settingscatalog": DeviceManagementConfigurationSettingDefinitionVisibilitysettingsCatalog,
		"template":        DeviceManagementConfigurationSettingDefinitionVisibilitytemplate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationSettingDefinitionVisibility(input)
	return &out, nil
}
