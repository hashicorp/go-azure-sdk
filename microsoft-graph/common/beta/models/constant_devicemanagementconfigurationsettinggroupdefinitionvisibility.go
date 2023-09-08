package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSettingGroupDefinitionVisibility string

const (
	DeviceManagementConfigurationSettingGroupDefinitionVisibilitynone            DeviceManagementConfigurationSettingGroupDefinitionVisibility = "None"
	DeviceManagementConfigurationSettingGroupDefinitionVisibilitysettingsCatalog DeviceManagementConfigurationSettingGroupDefinitionVisibility = "SettingsCatalog"
	DeviceManagementConfigurationSettingGroupDefinitionVisibilitytemplate        DeviceManagementConfigurationSettingGroupDefinitionVisibility = "Template"
)

func PossibleValuesForDeviceManagementConfigurationSettingGroupDefinitionVisibility() []string {
	return []string{
		string(DeviceManagementConfigurationSettingGroupDefinitionVisibilitynone),
		string(DeviceManagementConfigurationSettingGroupDefinitionVisibilitysettingsCatalog),
		string(DeviceManagementConfigurationSettingGroupDefinitionVisibilitytemplate),
	}
}

func parseDeviceManagementConfigurationSettingGroupDefinitionVisibility(input string) (*DeviceManagementConfigurationSettingGroupDefinitionVisibility, error) {
	vals := map[string]DeviceManagementConfigurationSettingGroupDefinitionVisibility{
		"none":            DeviceManagementConfigurationSettingGroupDefinitionVisibilitynone,
		"settingscatalog": DeviceManagementConfigurationSettingGroupDefinitionVisibilitysettingsCatalog,
		"template":        DeviceManagementConfigurationSettingGroupDefinitionVisibilitytemplate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationSettingGroupDefinitionVisibility(input)
	return &out, nil
}
