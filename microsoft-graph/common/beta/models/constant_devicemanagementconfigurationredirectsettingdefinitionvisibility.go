package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationRedirectSettingDefinitionVisibility string

const (
	DeviceManagementConfigurationRedirectSettingDefinitionVisibilitynone            DeviceManagementConfigurationRedirectSettingDefinitionVisibility = "None"
	DeviceManagementConfigurationRedirectSettingDefinitionVisibilitysettingsCatalog DeviceManagementConfigurationRedirectSettingDefinitionVisibility = "SettingsCatalog"
	DeviceManagementConfigurationRedirectSettingDefinitionVisibilitytemplate        DeviceManagementConfigurationRedirectSettingDefinitionVisibility = "Template"
)

func PossibleValuesForDeviceManagementConfigurationRedirectSettingDefinitionVisibility() []string {
	return []string{
		string(DeviceManagementConfigurationRedirectSettingDefinitionVisibilitynone),
		string(DeviceManagementConfigurationRedirectSettingDefinitionVisibilitysettingsCatalog),
		string(DeviceManagementConfigurationRedirectSettingDefinitionVisibilitytemplate),
	}
}

func parseDeviceManagementConfigurationRedirectSettingDefinitionVisibility(input string) (*DeviceManagementConfigurationRedirectSettingDefinitionVisibility, error) {
	vals := map[string]DeviceManagementConfigurationRedirectSettingDefinitionVisibility{
		"none":            DeviceManagementConfigurationRedirectSettingDefinitionVisibilitynone,
		"settingscatalog": DeviceManagementConfigurationRedirectSettingDefinitionVisibilitysettingsCatalog,
		"template":        DeviceManagementConfigurationRedirectSettingDefinitionVisibilitytemplate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationRedirectSettingDefinitionVisibility(input)
	return &out, nil
}
