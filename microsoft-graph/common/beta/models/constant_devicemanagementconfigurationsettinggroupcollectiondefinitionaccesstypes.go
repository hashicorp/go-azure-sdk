package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypes string

const (
	DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypesadd     DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypes = "Add"
	DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypescopy    DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypes = "Copy"
	DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypesdelete  DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypes = "Delete"
	DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypesexecute DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypes = "Execute"
	DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypesget     DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypes = "Get"
	DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypesnone    DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypes = "None"
	DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypesreplace DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypes = "Replace"
)

func PossibleValuesForDeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypes() []string {
	return []string{
		string(DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypesadd),
		string(DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypescopy),
		string(DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypesdelete),
		string(DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypesexecute),
		string(DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypesget),
		string(DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypesnone),
		string(DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypesreplace),
	}
}

func parseDeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypes(input string) (*DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypes, error) {
	vals := map[string]DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypes{
		"add":     DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypesadd,
		"copy":    DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypescopy,
		"delete":  DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypesdelete,
		"execute": DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypesexecute,
		"get":     DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypesget,
		"none":    DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypesnone,
		"replace": DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypesreplace,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypes(input)
	return &out, nil
}
