package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypes string

const (
	DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypesadd     DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypes = "Add"
	DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypescopy    DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypes = "Copy"
	DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypesdelete  DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypes = "Delete"
	DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypesexecute DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypes = "Execute"
	DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypesget     DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypes = "Get"
	DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypesnone    DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypes = "None"
	DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypesreplace DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypes = "Replace"
)

func PossibleValuesForDeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypes() []string {
	return []string{
		string(DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypesadd),
		string(DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypescopy),
		string(DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypesdelete),
		string(DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypesexecute),
		string(DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypesget),
		string(DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypesnone),
		string(DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypesreplace),
	}
}

func parseDeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypes(input string) (*DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypes, error) {
	vals := map[string]DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypes{
		"add":     DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypesadd,
		"copy":    DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypescopy,
		"delete":  DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypesdelete,
		"execute": DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypesexecute,
		"get":     DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypesget,
		"none":    DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypesnone,
		"replace": DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypesreplace,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypes(input)
	return &out, nil
}
