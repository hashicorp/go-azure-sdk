package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSimpleSettingDefinitionAccessTypes string

const (
	DeviceManagementConfigurationSimpleSettingDefinitionAccessTypesadd     DeviceManagementConfigurationSimpleSettingDefinitionAccessTypes = "Add"
	DeviceManagementConfigurationSimpleSettingDefinitionAccessTypescopy    DeviceManagementConfigurationSimpleSettingDefinitionAccessTypes = "Copy"
	DeviceManagementConfigurationSimpleSettingDefinitionAccessTypesdelete  DeviceManagementConfigurationSimpleSettingDefinitionAccessTypes = "Delete"
	DeviceManagementConfigurationSimpleSettingDefinitionAccessTypesexecute DeviceManagementConfigurationSimpleSettingDefinitionAccessTypes = "Execute"
	DeviceManagementConfigurationSimpleSettingDefinitionAccessTypesget     DeviceManagementConfigurationSimpleSettingDefinitionAccessTypes = "Get"
	DeviceManagementConfigurationSimpleSettingDefinitionAccessTypesnone    DeviceManagementConfigurationSimpleSettingDefinitionAccessTypes = "None"
	DeviceManagementConfigurationSimpleSettingDefinitionAccessTypesreplace DeviceManagementConfigurationSimpleSettingDefinitionAccessTypes = "Replace"
)

func PossibleValuesForDeviceManagementConfigurationSimpleSettingDefinitionAccessTypes() []string {
	return []string{
		string(DeviceManagementConfigurationSimpleSettingDefinitionAccessTypesadd),
		string(DeviceManagementConfigurationSimpleSettingDefinitionAccessTypescopy),
		string(DeviceManagementConfigurationSimpleSettingDefinitionAccessTypesdelete),
		string(DeviceManagementConfigurationSimpleSettingDefinitionAccessTypesexecute),
		string(DeviceManagementConfigurationSimpleSettingDefinitionAccessTypesget),
		string(DeviceManagementConfigurationSimpleSettingDefinitionAccessTypesnone),
		string(DeviceManagementConfigurationSimpleSettingDefinitionAccessTypesreplace),
	}
}

func parseDeviceManagementConfigurationSimpleSettingDefinitionAccessTypes(input string) (*DeviceManagementConfigurationSimpleSettingDefinitionAccessTypes, error) {
	vals := map[string]DeviceManagementConfigurationSimpleSettingDefinitionAccessTypes{
		"add":     DeviceManagementConfigurationSimpleSettingDefinitionAccessTypesadd,
		"copy":    DeviceManagementConfigurationSimpleSettingDefinitionAccessTypescopy,
		"delete":  DeviceManagementConfigurationSimpleSettingDefinitionAccessTypesdelete,
		"execute": DeviceManagementConfigurationSimpleSettingDefinitionAccessTypesexecute,
		"get":     DeviceManagementConfigurationSimpleSettingDefinitionAccessTypesget,
		"none":    DeviceManagementConfigurationSimpleSettingDefinitionAccessTypesnone,
		"replace": DeviceManagementConfigurationSimpleSettingDefinitionAccessTypesreplace,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationSimpleSettingDefinitionAccessTypes(input)
	return &out, nil
}
