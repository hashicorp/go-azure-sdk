package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationRedirectSettingDefinitionAccessTypes string

const (
	DeviceManagementConfigurationRedirectSettingDefinitionAccessTypesadd     DeviceManagementConfigurationRedirectSettingDefinitionAccessTypes = "Add"
	DeviceManagementConfigurationRedirectSettingDefinitionAccessTypescopy    DeviceManagementConfigurationRedirectSettingDefinitionAccessTypes = "Copy"
	DeviceManagementConfigurationRedirectSettingDefinitionAccessTypesdelete  DeviceManagementConfigurationRedirectSettingDefinitionAccessTypes = "Delete"
	DeviceManagementConfigurationRedirectSettingDefinitionAccessTypesexecute DeviceManagementConfigurationRedirectSettingDefinitionAccessTypes = "Execute"
	DeviceManagementConfigurationRedirectSettingDefinitionAccessTypesget     DeviceManagementConfigurationRedirectSettingDefinitionAccessTypes = "Get"
	DeviceManagementConfigurationRedirectSettingDefinitionAccessTypesnone    DeviceManagementConfigurationRedirectSettingDefinitionAccessTypes = "None"
	DeviceManagementConfigurationRedirectSettingDefinitionAccessTypesreplace DeviceManagementConfigurationRedirectSettingDefinitionAccessTypes = "Replace"
)

func PossibleValuesForDeviceManagementConfigurationRedirectSettingDefinitionAccessTypes() []string {
	return []string{
		string(DeviceManagementConfigurationRedirectSettingDefinitionAccessTypesadd),
		string(DeviceManagementConfigurationRedirectSettingDefinitionAccessTypescopy),
		string(DeviceManagementConfigurationRedirectSettingDefinitionAccessTypesdelete),
		string(DeviceManagementConfigurationRedirectSettingDefinitionAccessTypesexecute),
		string(DeviceManagementConfigurationRedirectSettingDefinitionAccessTypesget),
		string(DeviceManagementConfigurationRedirectSettingDefinitionAccessTypesnone),
		string(DeviceManagementConfigurationRedirectSettingDefinitionAccessTypesreplace),
	}
}

func parseDeviceManagementConfigurationRedirectSettingDefinitionAccessTypes(input string) (*DeviceManagementConfigurationRedirectSettingDefinitionAccessTypes, error) {
	vals := map[string]DeviceManagementConfigurationRedirectSettingDefinitionAccessTypes{
		"add":     DeviceManagementConfigurationRedirectSettingDefinitionAccessTypesadd,
		"copy":    DeviceManagementConfigurationRedirectSettingDefinitionAccessTypescopy,
		"delete":  DeviceManagementConfigurationRedirectSettingDefinitionAccessTypesdelete,
		"execute": DeviceManagementConfigurationRedirectSettingDefinitionAccessTypesexecute,
		"get":     DeviceManagementConfigurationRedirectSettingDefinitionAccessTypesget,
		"none":    DeviceManagementConfigurationRedirectSettingDefinitionAccessTypesnone,
		"replace": DeviceManagementConfigurationRedirectSettingDefinitionAccessTypesreplace,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationRedirectSettingDefinitionAccessTypes(input)
	return &out, nil
}
