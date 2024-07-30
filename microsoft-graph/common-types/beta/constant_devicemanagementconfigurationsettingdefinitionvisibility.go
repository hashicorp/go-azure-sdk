package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSettingDefinitionVisibility string

const (
	DeviceManagementConfigurationSettingDefinitionVisibility_None            DeviceManagementConfigurationSettingDefinitionVisibility = "none"
	DeviceManagementConfigurationSettingDefinitionVisibility_SettingsCatalog DeviceManagementConfigurationSettingDefinitionVisibility = "settingsCatalog"
	DeviceManagementConfigurationSettingDefinitionVisibility_Template        DeviceManagementConfigurationSettingDefinitionVisibility = "template"
)

func PossibleValuesForDeviceManagementConfigurationSettingDefinitionVisibility() []string {
	return []string{
		string(DeviceManagementConfigurationSettingDefinitionVisibility_None),
		string(DeviceManagementConfigurationSettingDefinitionVisibility_SettingsCatalog),
		string(DeviceManagementConfigurationSettingDefinitionVisibility_Template),
	}
}

func (s *DeviceManagementConfigurationSettingDefinitionVisibility) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationSettingDefinitionVisibility(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationSettingDefinitionVisibility(input string) (*DeviceManagementConfigurationSettingDefinitionVisibility, error) {
	vals := map[string]DeviceManagementConfigurationSettingDefinitionVisibility{
		"none":            DeviceManagementConfigurationSettingDefinitionVisibility_None,
		"settingscatalog": DeviceManagementConfigurationSettingDefinitionVisibility_SettingsCatalog,
		"template":        DeviceManagementConfigurationSettingDefinitionVisibility_Template,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationSettingDefinitionVisibility(input)
	return &out, nil
}
