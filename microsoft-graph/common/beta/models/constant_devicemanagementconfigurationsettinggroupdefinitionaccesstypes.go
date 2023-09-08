package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSettingGroupDefinitionAccessTypes string

const (
	DeviceManagementConfigurationSettingGroupDefinitionAccessTypesadd     DeviceManagementConfigurationSettingGroupDefinitionAccessTypes = "Add"
	DeviceManagementConfigurationSettingGroupDefinitionAccessTypescopy    DeviceManagementConfigurationSettingGroupDefinitionAccessTypes = "Copy"
	DeviceManagementConfigurationSettingGroupDefinitionAccessTypesdelete  DeviceManagementConfigurationSettingGroupDefinitionAccessTypes = "Delete"
	DeviceManagementConfigurationSettingGroupDefinitionAccessTypesexecute DeviceManagementConfigurationSettingGroupDefinitionAccessTypes = "Execute"
	DeviceManagementConfigurationSettingGroupDefinitionAccessTypesget     DeviceManagementConfigurationSettingGroupDefinitionAccessTypes = "Get"
	DeviceManagementConfigurationSettingGroupDefinitionAccessTypesnone    DeviceManagementConfigurationSettingGroupDefinitionAccessTypes = "None"
	DeviceManagementConfigurationSettingGroupDefinitionAccessTypesreplace DeviceManagementConfigurationSettingGroupDefinitionAccessTypes = "Replace"
)

func PossibleValuesForDeviceManagementConfigurationSettingGroupDefinitionAccessTypes() []string {
	return []string{
		string(DeviceManagementConfigurationSettingGroupDefinitionAccessTypesadd),
		string(DeviceManagementConfigurationSettingGroupDefinitionAccessTypescopy),
		string(DeviceManagementConfigurationSettingGroupDefinitionAccessTypesdelete),
		string(DeviceManagementConfigurationSettingGroupDefinitionAccessTypesexecute),
		string(DeviceManagementConfigurationSettingGroupDefinitionAccessTypesget),
		string(DeviceManagementConfigurationSettingGroupDefinitionAccessTypesnone),
		string(DeviceManagementConfigurationSettingGroupDefinitionAccessTypesreplace),
	}
}

func parseDeviceManagementConfigurationSettingGroupDefinitionAccessTypes(input string) (*DeviceManagementConfigurationSettingGroupDefinitionAccessTypes, error) {
	vals := map[string]DeviceManagementConfigurationSettingGroupDefinitionAccessTypes{
		"add":     DeviceManagementConfigurationSettingGroupDefinitionAccessTypesadd,
		"copy":    DeviceManagementConfigurationSettingGroupDefinitionAccessTypescopy,
		"delete":  DeviceManagementConfigurationSettingGroupDefinitionAccessTypesdelete,
		"execute": DeviceManagementConfigurationSettingGroupDefinitionAccessTypesexecute,
		"get":     DeviceManagementConfigurationSettingGroupDefinitionAccessTypesget,
		"none":    DeviceManagementConfigurationSettingGroupDefinitionAccessTypesnone,
		"replace": DeviceManagementConfigurationSettingGroupDefinitionAccessTypesreplace,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationSettingGroupDefinitionAccessTypes(input)
	return &out, nil
}
