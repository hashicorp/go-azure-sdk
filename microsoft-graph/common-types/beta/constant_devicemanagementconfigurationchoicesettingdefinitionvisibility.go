package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationChoiceSettingDefinitionVisibility string

const (
	DeviceManagementConfigurationChoiceSettingDefinitionVisibility_None            DeviceManagementConfigurationChoiceSettingDefinitionVisibility = "none"
	DeviceManagementConfigurationChoiceSettingDefinitionVisibility_SettingsCatalog DeviceManagementConfigurationChoiceSettingDefinitionVisibility = "settingsCatalog"
	DeviceManagementConfigurationChoiceSettingDefinitionVisibility_Template        DeviceManagementConfigurationChoiceSettingDefinitionVisibility = "template"
)

func PossibleValuesForDeviceManagementConfigurationChoiceSettingDefinitionVisibility() []string {
	return []string{
		string(DeviceManagementConfigurationChoiceSettingDefinitionVisibility_None),
		string(DeviceManagementConfigurationChoiceSettingDefinitionVisibility_SettingsCatalog),
		string(DeviceManagementConfigurationChoiceSettingDefinitionVisibility_Template),
	}
}

func (s *DeviceManagementConfigurationChoiceSettingDefinitionVisibility) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationChoiceSettingDefinitionVisibility(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationChoiceSettingDefinitionVisibility(input string) (*DeviceManagementConfigurationChoiceSettingDefinitionVisibility, error) {
	vals := map[string]DeviceManagementConfigurationChoiceSettingDefinitionVisibility{
		"none":            DeviceManagementConfigurationChoiceSettingDefinitionVisibility_None,
		"settingscatalog": DeviceManagementConfigurationChoiceSettingDefinitionVisibility_SettingsCatalog,
		"template":        DeviceManagementConfigurationChoiceSettingDefinitionVisibility_Template,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationChoiceSettingDefinitionVisibility(input)
	return &out, nil
}
