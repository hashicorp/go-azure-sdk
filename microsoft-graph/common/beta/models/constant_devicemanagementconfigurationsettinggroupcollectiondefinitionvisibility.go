package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSettingGroupCollectionDefinitionVisibility string

const (
	DeviceManagementConfigurationSettingGroupCollectionDefinitionVisibilitynone            DeviceManagementConfigurationSettingGroupCollectionDefinitionVisibility = "None"
	DeviceManagementConfigurationSettingGroupCollectionDefinitionVisibilitysettingsCatalog DeviceManagementConfigurationSettingGroupCollectionDefinitionVisibility = "SettingsCatalog"
	DeviceManagementConfigurationSettingGroupCollectionDefinitionVisibilitytemplate        DeviceManagementConfigurationSettingGroupCollectionDefinitionVisibility = "Template"
)

func PossibleValuesForDeviceManagementConfigurationSettingGroupCollectionDefinitionVisibility() []string {
	return []string{
		string(DeviceManagementConfigurationSettingGroupCollectionDefinitionVisibilitynone),
		string(DeviceManagementConfigurationSettingGroupCollectionDefinitionVisibilitysettingsCatalog),
		string(DeviceManagementConfigurationSettingGroupCollectionDefinitionVisibilitytemplate),
	}
}

func parseDeviceManagementConfigurationSettingGroupCollectionDefinitionVisibility(input string) (*DeviceManagementConfigurationSettingGroupCollectionDefinitionVisibility, error) {
	vals := map[string]DeviceManagementConfigurationSettingGroupCollectionDefinitionVisibility{
		"none":            DeviceManagementConfigurationSettingGroupCollectionDefinitionVisibilitynone,
		"settingscatalog": DeviceManagementConfigurationSettingGroupCollectionDefinitionVisibilitysettingsCatalog,
		"template":        DeviceManagementConfigurationSettingGroupCollectionDefinitionVisibilitytemplate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationSettingGroupCollectionDefinitionVisibility(input)
	return &out, nil
}
