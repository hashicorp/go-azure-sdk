package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSimpleSettingDefinitionAccessTypes string

const (
	DeviceManagementConfigurationSimpleSettingDefinitionAccessTypes_Add     DeviceManagementConfigurationSimpleSettingDefinitionAccessTypes = "add"
	DeviceManagementConfigurationSimpleSettingDefinitionAccessTypes_Copy    DeviceManagementConfigurationSimpleSettingDefinitionAccessTypes = "copy"
	DeviceManagementConfigurationSimpleSettingDefinitionAccessTypes_Delete  DeviceManagementConfigurationSimpleSettingDefinitionAccessTypes = "delete"
	DeviceManagementConfigurationSimpleSettingDefinitionAccessTypes_Execute DeviceManagementConfigurationSimpleSettingDefinitionAccessTypes = "execute"
	DeviceManagementConfigurationSimpleSettingDefinitionAccessTypes_Get     DeviceManagementConfigurationSimpleSettingDefinitionAccessTypes = "get"
	DeviceManagementConfigurationSimpleSettingDefinitionAccessTypes_None    DeviceManagementConfigurationSimpleSettingDefinitionAccessTypes = "none"
	DeviceManagementConfigurationSimpleSettingDefinitionAccessTypes_Replace DeviceManagementConfigurationSimpleSettingDefinitionAccessTypes = "replace"
)

func PossibleValuesForDeviceManagementConfigurationSimpleSettingDefinitionAccessTypes() []string {
	return []string{
		string(DeviceManagementConfigurationSimpleSettingDefinitionAccessTypes_Add),
		string(DeviceManagementConfigurationSimpleSettingDefinitionAccessTypes_Copy),
		string(DeviceManagementConfigurationSimpleSettingDefinitionAccessTypes_Delete),
		string(DeviceManagementConfigurationSimpleSettingDefinitionAccessTypes_Execute),
		string(DeviceManagementConfigurationSimpleSettingDefinitionAccessTypes_Get),
		string(DeviceManagementConfigurationSimpleSettingDefinitionAccessTypes_None),
		string(DeviceManagementConfigurationSimpleSettingDefinitionAccessTypes_Replace),
	}
}

func (s *DeviceManagementConfigurationSimpleSettingDefinitionAccessTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationSimpleSettingDefinitionAccessTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationSimpleSettingDefinitionAccessTypes(input string) (*DeviceManagementConfigurationSimpleSettingDefinitionAccessTypes, error) {
	vals := map[string]DeviceManagementConfigurationSimpleSettingDefinitionAccessTypes{
		"add":     DeviceManagementConfigurationSimpleSettingDefinitionAccessTypes_Add,
		"copy":    DeviceManagementConfigurationSimpleSettingDefinitionAccessTypes_Copy,
		"delete":  DeviceManagementConfigurationSimpleSettingDefinitionAccessTypes_Delete,
		"execute": DeviceManagementConfigurationSimpleSettingDefinitionAccessTypes_Execute,
		"get":     DeviceManagementConfigurationSimpleSettingDefinitionAccessTypes_Get,
		"none":    DeviceManagementConfigurationSimpleSettingDefinitionAccessTypes_None,
		"replace": DeviceManagementConfigurationSimpleSettingDefinitionAccessTypes_Replace,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationSimpleSettingDefinitionAccessTypes(input)
	return &out, nil
}
