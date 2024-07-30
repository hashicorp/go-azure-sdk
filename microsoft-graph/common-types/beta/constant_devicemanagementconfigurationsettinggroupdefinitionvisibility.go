package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSettingGroupDefinitionVisibility string

const (
	DeviceManagementConfigurationSettingGroupDefinitionVisibility_None            DeviceManagementConfigurationSettingGroupDefinitionVisibility = "none"
	DeviceManagementConfigurationSettingGroupDefinitionVisibility_SettingsCatalog DeviceManagementConfigurationSettingGroupDefinitionVisibility = "settingsCatalog"
	DeviceManagementConfigurationSettingGroupDefinitionVisibility_Template        DeviceManagementConfigurationSettingGroupDefinitionVisibility = "template"
)

func PossibleValuesForDeviceManagementConfigurationSettingGroupDefinitionVisibility() []string {
	return []string{
		string(DeviceManagementConfigurationSettingGroupDefinitionVisibility_None),
		string(DeviceManagementConfigurationSettingGroupDefinitionVisibility_SettingsCatalog),
		string(DeviceManagementConfigurationSettingGroupDefinitionVisibility_Template),
	}
}

func (s *DeviceManagementConfigurationSettingGroupDefinitionVisibility) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationSettingGroupDefinitionVisibility(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationSettingGroupDefinitionVisibility(input string) (*DeviceManagementConfigurationSettingGroupDefinitionVisibility, error) {
	vals := map[string]DeviceManagementConfigurationSettingGroupDefinitionVisibility{
		"none":            DeviceManagementConfigurationSettingGroupDefinitionVisibility_None,
		"settingscatalog": DeviceManagementConfigurationSettingGroupDefinitionVisibility_SettingsCatalog,
		"template":        DeviceManagementConfigurationSettingGroupDefinitionVisibility_Template,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationSettingGroupDefinitionVisibility(input)
	return &out, nil
}
