package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationRedirectSettingDefinitionVisibility string

const (
	DeviceManagementConfigurationRedirectSettingDefinitionVisibility_None            DeviceManagementConfigurationRedirectSettingDefinitionVisibility = "none"
	DeviceManagementConfigurationRedirectSettingDefinitionVisibility_SettingsCatalog DeviceManagementConfigurationRedirectSettingDefinitionVisibility = "settingsCatalog"
	DeviceManagementConfigurationRedirectSettingDefinitionVisibility_Template        DeviceManagementConfigurationRedirectSettingDefinitionVisibility = "template"
)

func PossibleValuesForDeviceManagementConfigurationRedirectSettingDefinitionVisibility() []string {
	return []string{
		string(DeviceManagementConfigurationRedirectSettingDefinitionVisibility_None),
		string(DeviceManagementConfigurationRedirectSettingDefinitionVisibility_SettingsCatalog),
		string(DeviceManagementConfigurationRedirectSettingDefinitionVisibility_Template),
	}
}

func (s *DeviceManagementConfigurationRedirectSettingDefinitionVisibility) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationRedirectSettingDefinitionVisibility(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationRedirectSettingDefinitionVisibility(input string) (*DeviceManagementConfigurationRedirectSettingDefinitionVisibility, error) {
	vals := map[string]DeviceManagementConfigurationRedirectSettingDefinitionVisibility{
		"none":            DeviceManagementConfigurationRedirectSettingDefinitionVisibility_None,
		"settingscatalog": DeviceManagementConfigurationRedirectSettingDefinitionVisibility_SettingsCatalog,
		"template":        DeviceManagementConfigurationRedirectSettingDefinitionVisibility_Template,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationRedirectSettingDefinitionVisibility(input)
	return &out, nil
}
