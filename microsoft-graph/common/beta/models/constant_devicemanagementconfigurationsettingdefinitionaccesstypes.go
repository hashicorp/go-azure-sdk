package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSettingDefinitionAccessTypes string

const (
	DeviceManagementConfigurationSettingDefinitionAccessTypesadd     DeviceManagementConfigurationSettingDefinitionAccessTypes = "Add"
	DeviceManagementConfigurationSettingDefinitionAccessTypescopy    DeviceManagementConfigurationSettingDefinitionAccessTypes = "Copy"
	DeviceManagementConfigurationSettingDefinitionAccessTypesdelete  DeviceManagementConfigurationSettingDefinitionAccessTypes = "Delete"
	DeviceManagementConfigurationSettingDefinitionAccessTypesexecute DeviceManagementConfigurationSettingDefinitionAccessTypes = "Execute"
	DeviceManagementConfigurationSettingDefinitionAccessTypesget     DeviceManagementConfigurationSettingDefinitionAccessTypes = "Get"
	DeviceManagementConfigurationSettingDefinitionAccessTypesnone    DeviceManagementConfigurationSettingDefinitionAccessTypes = "None"
	DeviceManagementConfigurationSettingDefinitionAccessTypesreplace DeviceManagementConfigurationSettingDefinitionAccessTypes = "Replace"
)

func PossibleValuesForDeviceManagementConfigurationSettingDefinitionAccessTypes() []string {
	return []string{
		string(DeviceManagementConfigurationSettingDefinitionAccessTypesadd),
		string(DeviceManagementConfigurationSettingDefinitionAccessTypescopy),
		string(DeviceManagementConfigurationSettingDefinitionAccessTypesdelete),
		string(DeviceManagementConfigurationSettingDefinitionAccessTypesexecute),
		string(DeviceManagementConfigurationSettingDefinitionAccessTypesget),
		string(DeviceManagementConfigurationSettingDefinitionAccessTypesnone),
		string(DeviceManagementConfigurationSettingDefinitionAccessTypesreplace),
	}
}

func parseDeviceManagementConfigurationSettingDefinitionAccessTypes(input string) (*DeviceManagementConfigurationSettingDefinitionAccessTypes, error) {
	vals := map[string]DeviceManagementConfigurationSettingDefinitionAccessTypes{
		"add":     DeviceManagementConfigurationSettingDefinitionAccessTypesadd,
		"copy":    DeviceManagementConfigurationSettingDefinitionAccessTypescopy,
		"delete":  DeviceManagementConfigurationSettingDefinitionAccessTypesdelete,
		"execute": DeviceManagementConfigurationSettingDefinitionAccessTypesexecute,
		"get":     DeviceManagementConfigurationSettingDefinitionAccessTypesget,
		"none":    DeviceManagementConfigurationSettingDefinitionAccessTypesnone,
		"replace": DeviceManagementConfigurationSettingDefinitionAccessTypesreplace,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationSettingDefinitionAccessTypes(input)
	return &out, nil
}
