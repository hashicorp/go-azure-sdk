package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationRedirectSettingDefinitionAccessTypes string

const (
	DeviceManagementConfigurationRedirectSettingDefinitionAccessTypes_Add     DeviceManagementConfigurationRedirectSettingDefinitionAccessTypes = "add"
	DeviceManagementConfigurationRedirectSettingDefinitionAccessTypes_Copy    DeviceManagementConfigurationRedirectSettingDefinitionAccessTypes = "copy"
	DeviceManagementConfigurationRedirectSettingDefinitionAccessTypes_Delete  DeviceManagementConfigurationRedirectSettingDefinitionAccessTypes = "delete"
	DeviceManagementConfigurationRedirectSettingDefinitionAccessTypes_Execute DeviceManagementConfigurationRedirectSettingDefinitionAccessTypes = "execute"
	DeviceManagementConfigurationRedirectSettingDefinitionAccessTypes_Get     DeviceManagementConfigurationRedirectSettingDefinitionAccessTypes = "get"
	DeviceManagementConfigurationRedirectSettingDefinitionAccessTypes_None    DeviceManagementConfigurationRedirectSettingDefinitionAccessTypes = "none"
	DeviceManagementConfigurationRedirectSettingDefinitionAccessTypes_Replace DeviceManagementConfigurationRedirectSettingDefinitionAccessTypes = "replace"
)

func PossibleValuesForDeviceManagementConfigurationRedirectSettingDefinitionAccessTypes() []string {
	return []string{
		string(DeviceManagementConfigurationRedirectSettingDefinitionAccessTypes_Add),
		string(DeviceManagementConfigurationRedirectSettingDefinitionAccessTypes_Copy),
		string(DeviceManagementConfigurationRedirectSettingDefinitionAccessTypes_Delete),
		string(DeviceManagementConfigurationRedirectSettingDefinitionAccessTypes_Execute),
		string(DeviceManagementConfigurationRedirectSettingDefinitionAccessTypes_Get),
		string(DeviceManagementConfigurationRedirectSettingDefinitionAccessTypes_None),
		string(DeviceManagementConfigurationRedirectSettingDefinitionAccessTypes_Replace),
	}
}

func (s *DeviceManagementConfigurationRedirectSettingDefinitionAccessTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationRedirectSettingDefinitionAccessTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationRedirectSettingDefinitionAccessTypes(input string) (*DeviceManagementConfigurationRedirectSettingDefinitionAccessTypes, error) {
	vals := map[string]DeviceManagementConfigurationRedirectSettingDefinitionAccessTypes{
		"add":     DeviceManagementConfigurationRedirectSettingDefinitionAccessTypes_Add,
		"copy":    DeviceManagementConfigurationRedirectSettingDefinitionAccessTypes_Copy,
		"delete":  DeviceManagementConfigurationRedirectSettingDefinitionAccessTypes_Delete,
		"execute": DeviceManagementConfigurationRedirectSettingDefinitionAccessTypes_Execute,
		"get":     DeviceManagementConfigurationRedirectSettingDefinitionAccessTypes_Get,
		"none":    DeviceManagementConfigurationRedirectSettingDefinitionAccessTypes_None,
		"replace": DeviceManagementConfigurationRedirectSettingDefinitionAccessTypes_Replace,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationRedirectSettingDefinitionAccessTypes(input)
	return &out, nil
}
