package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSimpleSettingCollectionDefinitionAccessTypes string

const (
	DeviceManagementConfigurationSimpleSettingCollectionDefinitionAccessTypes_Add     DeviceManagementConfigurationSimpleSettingCollectionDefinitionAccessTypes = "add"
	DeviceManagementConfigurationSimpleSettingCollectionDefinitionAccessTypes_Copy    DeviceManagementConfigurationSimpleSettingCollectionDefinitionAccessTypes = "copy"
	DeviceManagementConfigurationSimpleSettingCollectionDefinitionAccessTypes_Delete  DeviceManagementConfigurationSimpleSettingCollectionDefinitionAccessTypes = "delete"
	DeviceManagementConfigurationSimpleSettingCollectionDefinitionAccessTypes_Execute DeviceManagementConfigurationSimpleSettingCollectionDefinitionAccessTypes = "execute"
	DeviceManagementConfigurationSimpleSettingCollectionDefinitionAccessTypes_Get     DeviceManagementConfigurationSimpleSettingCollectionDefinitionAccessTypes = "get"
	DeviceManagementConfigurationSimpleSettingCollectionDefinitionAccessTypes_None    DeviceManagementConfigurationSimpleSettingCollectionDefinitionAccessTypes = "none"
	DeviceManagementConfigurationSimpleSettingCollectionDefinitionAccessTypes_Replace DeviceManagementConfigurationSimpleSettingCollectionDefinitionAccessTypes = "replace"
)

func PossibleValuesForDeviceManagementConfigurationSimpleSettingCollectionDefinitionAccessTypes() []string {
	return []string{
		string(DeviceManagementConfigurationSimpleSettingCollectionDefinitionAccessTypes_Add),
		string(DeviceManagementConfigurationSimpleSettingCollectionDefinitionAccessTypes_Copy),
		string(DeviceManagementConfigurationSimpleSettingCollectionDefinitionAccessTypes_Delete),
		string(DeviceManagementConfigurationSimpleSettingCollectionDefinitionAccessTypes_Execute),
		string(DeviceManagementConfigurationSimpleSettingCollectionDefinitionAccessTypes_Get),
		string(DeviceManagementConfigurationSimpleSettingCollectionDefinitionAccessTypes_None),
		string(DeviceManagementConfigurationSimpleSettingCollectionDefinitionAccessTypes_Replace),
	}
}

func (s *DeviceManagementConfigurationSimpleSettingCollectionDefinitionAccessTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationSimpleSettingCollectionDefinitionAccessTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationSimpleSettingCollectionDefinitionAccessTypes(input string) (*DeviceManagementConfigurationSimpleSettingCollectionDefinitionAccessTypes, error) {
	vals := map[string]DeviceManagementConfigurationSimpleSettingCollectionDefinitionAccessTypes{
		"add":     DeviceManagementConfigurationSimpleSettingCollectionDefinitionAccessTypes_Add,
		"copy":    DeviceManagementConfigurationSimpleSettingCollectionDefinitionAccessTypes_Copy,
		"delete":  DeviceManagementConfigurationSimpleSettingCollectionDefinitionAccessTypes_Delete,
		"execute": DeviceManagementConfigurationSimpleSettingCollectionDefinitionAccessTypes_Execute,
		"get":     DeviceManagementConfigurationSimpleSettingCollectionDefinitionAccessTypes_Get,
		"none":    DeviceManagementConfigurationSimpleSettingCollectionDefinitionAccessTypes_None,
		"replace": DeviceManagementConfigurationSimpleSettingCollectionDefinitionAccessTypes_Replace,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationSimpleSettingCollectionDefinitionAccessTypes(input)
	return &out, nil
}
