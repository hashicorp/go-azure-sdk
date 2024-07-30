package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationChoiceSettingDefinitionAccessTypes string

const (
	DeviceManagementConfigurationChoiceSettingDefinitionAccessTypes_Add     DeviceManagementConfigurationChoiceSettingDefinitionAccessTypes = "add"
	DeviceManagementConfigurationChoiceSettingDefinitionAccessTypes_Copy    DeviceManagementConfigurationChoiceSettingDefinitionAccessTypes = "copy"
	DeviceManagementConfigurationChoiceSettingDefinitionAccessTypes_Delete  DeviceManagementConfigurationChoiceSettingDefinitionAccessTypes = "delete"
	DeviceManagementConfigurationChoiceSettingDefinitionAccessTypes_Execute DeviceManagementConfigurationChoiceSettingDefinitionAccessTypes = "execute"
	DeviceManagementConfigurationChoiceSettingDefinitionAccessTypes_Get     DeviceManagementConfigurationChoiceSettingDefinitionAccessTypes = "get"
	DeviceManagementConfigurationChoiceSettingDefinitionAccessTypes_None    DeviceManagementConfigurationChoiceSettingDefinitionAccessTypes = "none"
	DeviceManagementConfigurationChoiceSettingDefinitionAccessTypes_Replace DeviceManagementConfigurationChoiceSettingDefinitionAccessTypes = "replace"
)

func PossibleValuesForDeviceManagementConfigurationChoiceSettingDefinitionAccessTypes() []string {
	return []string{
		string(DeviceManagementConfigurationChoiceSettingDefinitionAccessTypes_Add),
		string(DeviceManagementConfigurationChoiceSettingDefinitionAccessTypes_Copy),
		string(DeviceManagementConfigurationChoiceSettingDefinitionAccessTypes_Delete),
		string(DeviceManagementConfigurationChoiceSettingDefinitionAccessTypes_Execute),
		string(DeviceManagementConfigurationChoiceSettingDefinitionAccessTypes_Get),
		string(DeviceManagementConfigurationChoiceSettingDefinitionAccessTypes_None),
		string(DeviceManagementConfigurationChoiceSettingDefinitionAccessTypes_Replace),
	}
}

func (s *DeviceManagementConfigurationChoiceSettingDefinitionAccessTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationChoiceSettingDefinitionAccessTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationChoiceSettingDefinitionAccessTypes(input string) (*DeviceManagementConfigurationChoiceSettingDefinitionAccessTypes, error) {
	vals := map[string]DeviceManagementConfigurationChoiceSettingDefinitionAccessTypes{
		"add":     DeviceManagementConfigurationChoiceSettingDefinitionAccessTypes_Add,
		"copy":    DeviceManagementConfigurationChoiceSettingDefinitionAccessTypes_Copy,
		"delete":  DeviceManagementConfigurationChoiceSettingDefinitionAccessTypes_Delete,
		"execute": DeviceManagementConfigurationChoiceSettingDefinitionAccessTypes_Execute,
		"get":     DeviceManagementConfigurationChoiceSettingDefinitionAccessTypes_Get,
		"none":    DeviceManagementConfigurationChoiceSettingDefinitionAccessTypes_None,
		"replace": DeviceManagementConfigurationChoiceSettingDefinitionAccessTypes_Replace,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationChoiceSettingDefinitionAccessTypes(input)
	return &out, nil
}
