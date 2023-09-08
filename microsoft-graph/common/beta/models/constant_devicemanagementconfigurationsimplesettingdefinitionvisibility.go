package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSimpleSettingDefinitionVisibility string

const (
	DeviceManagementConfigurationSimpleSettingDefinitionVisibilitynone            DeviceManagementConfigurationSimpleSettingDefinitionVisibility = "None"
	DeviceManagementConfigurationSimpleSettingDefinitionVisibilitysettingsCatalog DeviceManagementConfigurationSimpleSettingDefinitionVisibility = "SettingsCatalog"
	DeviceManagementConfigurationSimpleSettingDefinitionVisibilitytemplate        DeviceManagementConfigurationSimpleSettingDefinitionVisibility = "Template"
)

func PossibleValuesForDeviceManagementConfigurationSimpleSettingDefinitionVisibility() []string {
	return []string{
		string(DeviceManagementConfigurationSimpleSettingDefinitionVisibilitynone),
		string(DeviceManagementConfigurationSimpleSettingDefinitionVisibilitysettingsCatalog),
		string(DeviceManagementConfigurationSimpleSettingDefinitionVisibilitytemplate),
	}
}

func parseDeviceManagementConfigurationSimpleSettingDefinitionVisibility(input string) (*DeviceManagementConfigurationSimpleSettingDefinitionVisibility, error) {
	vals := map[string]DeviceManagementConfigurationSimpleSettingDefinitionVisibility{
		"none":            DeviceManagementConfigurationSimpleSettingDefinitionVisibilitynone,
		"settingscatalog": DeviceManagementConfigurationSimpleSettingDefinitionVisibilitysettingsCatalog,
		"template":        DeviceManagementConfigurationSimpleSettingDefinitionVisibilitytemplate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationSimpleSettingDefinitionVisibility(input)
	return &out, nil
}
