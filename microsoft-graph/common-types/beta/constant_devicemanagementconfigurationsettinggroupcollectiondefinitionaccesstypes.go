package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypes string

const (
	DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypes_Add     DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypes = "add"
	DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypes_Copy    DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypes = "copy"
	DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypes_Delete  DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypes = "delete"
	DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypes_Execute DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypes = "execute"
	DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypes_Get     DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypes = "get"
	DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypes_None    DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypes = "none"
	DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypes_Replace DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypes = "replace"
)

func PossibleValuesForDeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypes() []string {
	return []string{
		string(DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypes_Add),
		string(DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypes_Copy),
		string(DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypes_Delete),
		string(DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypes_Execute),
		string(DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypes_Get),
		string(DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypes_None),
		string(DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypes_Replace),
	}
}

func (s *DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypes(input string) (*DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypes, error) {
	vals := map[string]DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypes{
		"add":     DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypes_Add,
		"copy":    DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypes_Copy,
		"delete":  DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypes_Delete,
		"execute": DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypes_Execute,
		"get":     DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypes_Get,
		"none":    DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypes_None,
		"replace": DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypes_Replace,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationSettingGroupCollectionDefinitionAccessTypes(input)
	return &out, nil
}
