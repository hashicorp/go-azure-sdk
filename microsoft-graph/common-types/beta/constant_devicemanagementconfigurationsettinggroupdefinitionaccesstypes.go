package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSettingGroupDefinitionAccessTypes string

const (
	DeviceManagementConfigurationSettingGroupDefinitionAccessTypes_Add     DeviceManagementConfigurationSettingGroupDefinitionAccessTypes = "add"
	DeviceManagementConfigurationSettingGroupDefinitionAccessTypes_Copy    DeviceManagementConfigurationSettingGroupDefinitionAccessTypes = "copy"
	DeviceManagementConfigurationSettingGroupDefinitionAccessTypes_Delete  DeviceManagementConfigurationSettingGroupDefinitionAccessTypes = "delete"
	DeviceManagementConfigurationSettingGroupDefinitionAccessTypes_Execute DeviceManagementConfigurationSettingGroupDefinitionAccessTypes = "execute"
	DeviceManagementConfigurationSettingGroupDefinitionAccessTypes_Get     DeviceManagementConfigurationSettingGroupDefinitionAccessTypes = "get"
	DeviceManagementConfigurationSettingGroupDefinitionAccessTypes_None    DeviceManagementConfigurationSettingGroupDefinitionAccessTypes = "none"
	DeviceManagementConfigurationSettingGroupDefinitionAccessTypes_Replace DeviceManagementConfigurationSettingGroupDefinitionAccessTypes = "replace"
)

func PossibleValuesForDeviceManagementConfigurationSettingGroupDefinitionAccessTypes() []string {
	return []string{
		string(DeviceManagementConfigurationSettingGroupDefinitionAccessTypes_Add),
		string(DeviceManagementConfigurationSettingGroupDefinitionAccessTypes_Copy),
		string(DeviceManagementConfigurationSettingGroupDefinitionAccessTypes_Delete),
		string(DeviceManagementConfigurationSettingGroupDefinitionAccessTypes_Execute),
		string(DeviceManagementConfigurationSettingGroupDefinitionAccessTypes_Get),
		string(DeviceManagementConfigurationSettingGroupDefinitionAccessTypes_None),
		string(DeviceManagementConfigurationSettingGroupDefinitionAccessTypes_Replace),
	}
}

func (s *DeviceManagementConfigurationSettingGroupDefinitionAccessTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationSettingGroupDefinitionAccessTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationSettingGroupDefinitionAccessTypes(input string) (*DeviceManagementConfigurationSettingGroupDefinitionAccessTypes, error) {
	vals := map[string]DeviceManagementConfigurationSettingGroupDefinitionAccessTypes{
		"add":     DeviceManagementConfigurationSettingGroupDefinitionAccessTypes_Add,
		"copy":    DeviceManagementConfigurationSettingGroupDefinitionAccessTypes_Copy,
		"delete":  DeviceManagementConfigurationSettingGroupDefinitionAccessTypes_Delete,
		"execute": DeviceManagementConfigurationSettingGroupDefinitionAccessTypes_Execute,
		"get":     DeviceManagementConfigurationSettingGroupDefinitionAccessTypes_Get,
		"none":    DeviceManagementConfigurationSettingGroupDefinitionAccessTypes_None,
		"replace": DeviceManagementConfigurationSettingGroupDefinitionAccessTypes_Replace,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationSettingGroupDefinitionAccessTypes(input)
	return &out, nil
}
