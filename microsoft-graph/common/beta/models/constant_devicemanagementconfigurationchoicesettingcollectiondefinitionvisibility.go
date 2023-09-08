package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationChoiceSettingCollectionDefinitionVisibility string

const (
	DeviceManagementConfigurationChoiceSettingCollectionDefinitionVisibilitynone            DeviceManagementConfigurationChoiceSettingCollectionDefinitionVisibility = "None"
	DeviceManagementConfigurationChoiceSettingCollectionDefinitionVisibilitysettingsCatalog DeviceManagementConfigurationChoiceSettingCollectionDefinitionVisibility = "SettingsCatalog"
	DeviceManagementConfigurationChoiceSettingCollectionDefinitionVisibilitytemplate        DeviceManagementConfigurationChoiceSettingCollectionDefinitionVisibility = "Template"
)

func PossibleValuesForDeviceManagementConfigurationChoiceSettingCollectionDefinitionVisibility() []string {
	return []string{
		string(DeviceManagementConfigurationChoiceSettingCollectionDefinitionVisibilitynone),
		string(DeviceManagementConfigurationChoiceSettingCollectionDefinitionVisibilitysettingsCatalog),
		string(DeviceManagementConfigurationChoiceSettingCollectionDefinitionVisibilitytemplate),
	}
}

func parseDeviceManagementConfigurationChoiceSettingCollectionDefinitionVisibility(input string) (*DeviceManagementConfigurationChoiceSettingCollectionDefinitionVisibility, error) {
	vals := map[string]DeviceManagementConfigurationChoiceSettingCollectionDefinitionVisibility{
		"none":            DeviceManagementConfigurationChoiceSettingCollectionDefinitionVisibilitynone,
		"settingscatalog": DeviceManagementConfigurationChoiceSettingCollectionDefinitionVisibilitysettingsCatalog,
		"template":        DeviceManagementConfigurationChoiceSettingCollectionDefinitionVisibilitytemplate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationChoiceSettingCollectionDefinitionVisibility(input)
	return &out, nil
}
