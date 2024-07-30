package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSettingDefinitionAccessTypes string

const (
	DeviceManagementConfigurationSettingDefinitionAccessTypes_Add     DeviceManagementConfigurationSettingDefinitionAccessTypes = "add"
	DeviceManagementConfigurationSettingDefinitionAccessTypes_Copy    DeviceManagementConfigurationSettingDefinitionAccessTypes = "copy"
	DeviceManagementConfigurationSettingDefinitionAccessTypes_Delete  DeviceManagementConfigurationSettingDefinitionAccessTypes = "delete"
	DeviceManagementConfigurationSettingDefinitionAccessTypes_Execute DeviceManagementConfigurationSettingDefinitionAccessTypes = "execute"
	DeviceManagementConfigurationSettingDefinitionAccessTypes_Get     DeviceManagementConfigurationSettingDefinitionAccessTypes = "get"
	DeviceManagementConfigurationSettingDefinitionAccessTypes_None    DeviceManagementConfigurationSettingDefinitionAccessTypes = "none"
	DeviceManagementConfigurationSettingDefinitionAccessTypes_Replace DeviceManagementConfigurationSettingDefinitionAccessTypes = "replace"
)

func PossibleValuesForDeviceManagementConfigurationSettingDefinitionAccessTypes() []string {
	return []string{
		string(DeviceManagementConfigurationSettingDefinitionAccessTypes_Add),
		string(DeviceManagementConfigurationSettingDefinitionAccessTypes_Copy),
		string(DeviceManagementConfigurationSettingDefinitionAccessTypes_Delete),
		string(DeviceManagementConfigurationSettingDefinitionAccessTypes_Execute),
		string(DeviceManagementConfigurationSettingDefinitionAccessTypes_Get),
		string(DeviceManagementConfigurationSettingDefinitionAccessTypes_None),
		string(DeviceManagementConfigurationSettingDefinitionAccessTypes_Replace),
	}
}

func (s *DeviceManagementConfigurationSettingDefinitionAccessTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationSettingDefinitionAccessTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationSettingDefinitionAccessTypes(input string) (*DeviceManagementConfigurationSettingDefinitionAccessTypes, error) {
	vals := map[string]DeviceManagementConfigurationSettingDefinitionAccessTypes{
		"add":     DeviceManagementConfigurationSettingDefinitionAccessTypes_Add,
		"copy":    DeviceManagementConfigurationSettingDefinitionAccessTypes_Copy,
		"delete":  DeviceManagementConfigurationSettingDefinitionAccessTypes_Delete,
		"execute": DeviceManagementConfigurationSettingDefinitionAccessTypes_Execute,
		"get":     DeviceManagementConfigurationSettingDefinitionAccessTypes_Get,
		"none":    DeviceManagementConfigurationSettingDefinitionAccessTypes_None,
		"replace": DeviceManagementConfigurationSettingDefinitionAccessTypes_Replace,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationSettingDefinitionAccessTypes(input)
	return &out, nil
}
