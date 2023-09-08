package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationChoiceSettingDefinitionAccessTypes string

const (
	DeviceManagementConfigurationChoiceSettingDefinitionAccessTypesadd     DeviceManagementConfigurationChoiceSettingDefinitionAccessTypes = "Add"
	DeviceManagementConfigurationChoiceSettingDefinitionAccessTypescopy    DeviceManagementConfigurationChoiceSettingDefinitionAccessTypes = "Copy"
	DeviceManagementConfigurationChoiceSettingDefinitionAccessTypesdelete  DeviceManagementConfigurationChoiceSettingDefinitionAccessTypes = "Delete"
	DeviceManagementConfigurationChoiceSettingDefinitionAccessTypesexecute DeviceManagementConfigurationChoiceSettingDefinitionAccessTypes = "Execute"
	DeviceManagementConfigurationChoiceSettingDefinitionAccessTypesget     DeviceManagementConfigurationChoiceSettingDefinitionAccessTypes = "Get"
	DeviceManagementConfigurationChoiceSettingDefinitionAccessTypesnone    DeviceManagementConfigurationChoiceSettingDefinitionAccessTypes = "None"
	DeviceManagementConfigurationChoiceSettingDefinitionAccessTypesreplace DeviceManagementConfigurationChoiceSettingDefinitionAccessTypes = "Replace"
)

func PossibleValuesForDeviceManagementConfigurationChoiceSettingDefinitionAccessTypes() []string {
	return []string{
		string(DeviceManagementConfigurationChoiceSettingDefinitionAccessTypesadd),
		string(DeviceManagementConfigurationChoiceSettingDefinitionAccessTypescopy),
		string(DeviceManagementConfigurationChoiceSettingDefinitionAccessTypesdelete),
		string(DeviceManagementConfigurationChoiceSettingDefinitionAccessTypesexecute),
		string(DeviceManagementConfigurationChoiceSettingDefinitionAccessTypesget),
		string(DeviceManagementConfigurationChoiceSettingDefinitionAccessTypesnone),
		string(DeviceManagementConfigurationChoiceSettingDefinitionAccessTypesreplace),
	}
}

func parseDeviceManagementConfigurationChoiceSettingDefinitionAccessTypes(input string) (*DeviceManagementConfigurationChoiceSettingDefinitionAccessTypes, error) {
	vals := map[string]DeviceManagementConfigurationChoiceSettingDefinitionAccessTypes{
		"add":     DeviceManagementConfigurationChoiceSettingDefinitionAccessTypesadd,
		"copy":    DeviceManagementConfigurationChoiceSettingDefinitionAccessTypescopy,
		"delete":  DeviceManagementConfigurationChoiceSettingDefinitionAccessTypesdelete,
		"execute": DeviceManagementConfigurationChoiceSettingDefinitionAccessTypesexecute,
		"get":     DeviceManagementConfigurationChoiceSettingDefinitionAccessTypesget,
		"none":    DeviceManagementConfigurationChoiceSettingDefinitionAccessTypesnone,
		"replace": DeviceManagementConfigurationChoiceSettingDefinitionAccessTypesreplace,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationChoiceSettingDefinitionAccessTypes(input)
	return &out, nil
}
