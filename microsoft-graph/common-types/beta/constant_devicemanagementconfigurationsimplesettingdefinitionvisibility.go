package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSimpleSettingDefinitionVisibility string

const (
	DeviceManagementConfigurationSimpleSettingDefinitionVisibility_None            DeviceManagementConfigurationSimpleSettingDefinitionVisibility = "none"
	DeviceManagementConfigurationSimpleSettingDefinitionVisibility_SettingsCatalog DeviceManagementConfigurationSimpleSettingDefinitionVisibility = "settingsCatalog"
	DeviceManagementConfigurationSimpleSettingDefinitionVisibility_Template        DeviceManagementConfigurationSimpleSettingDefinitionVisibility = "template"
)

func PossibleValuesForDeviceManagementConfigurationSimpleSettingDefinitionVisibility() []string {
	return []string{
		string(DeviceManagementConfigurationSimpleSettingDefinitionVisibility_None),
		string(DeviceManagementConfigurationSimpleSettingDefinitionVisibility_SettingsCatalog),
		string(DeviceManagementConfigurationSimpleSettingDefinitionVisibility_Template),
	}
}

func (s *DeviceManagementConfigurationSimpleSettingDefinitionVisibility) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationSimpleSettingDefinitionVisibility(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationSimpleSettingDefinitionVisibility(input string) (*DeviceManagementConfigurationSimpleSettingDefinitionVisibility, error) {
	vals := map[string]DeviceManagementConfigurationSimpleSettingDefinitionVisibility{
		"none":            DeviceManagementConfigurationSimpleSettingDefinitionVisibility_None,
		"settingscatalog": DeviceManagementConfigurationSimpleSettingDefinitionVisibility_SettingsCatalog,
		"template":        DeviceManagementConfigurationSimpleSettingDefinitionVisibility_Template,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationSimpleSettingDefinitionVisibility(input)
	return &out, nil
}
