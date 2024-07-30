package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypes string

const (
	DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypes_Add     DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypes = "add"
	DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypes_Copy    DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypes = "copy"
	DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypes_Delete  DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypes = "delete"
	DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypes_Execute DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypes = "execute"
	DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypes_Get     DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypes = "get"
	DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypes_None    DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypes = "none"
	DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypes_Replace DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypes = "replace"
)

func PossibleValuesForDeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypes() []string {
	return []string{
		string(DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypes_Add),
		string(DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypes_Copy),
		string(DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypes_Delete),
		string(DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypes_Execute),
		string(DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypes_Get),
		string(DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypes_None),
		string(DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypes_Replace),
	}
}

func (s *DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypes(input string) (*DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypes, error) {
	vals := map[string]DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypes{
		"add":     DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypes_Add,
		"copy":    DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypes_Copy,
		"delete":  DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypes_Delete,
		"execute": DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypes_Execute,
		"get":     DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypes_Get,
		"none":    DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypes_None,
		"replace": DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypes_Replace,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationChoiceSettingCollectionDefinitionAccessTypes(input)
	return &out, nil
}
