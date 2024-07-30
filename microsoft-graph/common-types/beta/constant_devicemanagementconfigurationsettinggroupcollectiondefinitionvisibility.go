package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSettingGroupCollectionDefinitionVisibility string

const (
	DeviceManagementConfigurationSettingGroupCollectionDefinitionVisibility_None            DeviceManagementConfigurationSettingGroupCollectionDefinitionVisibility = "none"
	DeviceManagementConfigurationSettingGroupCollectionDefinitionVisibility_SettingsCatalog DeviceManagementConfigurationSettingGroupCollectionDefinitionVisibility = "settingsCatalog"
	DeviceManagementConfigurationSettingGroupCollectionDefinitionVisibility_Template        DeviceManagementConfigurationSettingGroupCollectionDefinitionVisibility = "template"
)

func PossibleValuesForDeviceManagementConfigurationSettingGroupCollectionDefinitionVisibility() []string {
	return []string{
		string(DeviceManagementConfigurationSettingGroupCollectionDefinitionVisibility_None),
		string(DeviceManagementConfigurationSettingGroupCollectionDefinitionVisibility_SettingsCatalog),
		string(DeviceManagementConfigurationSettingGroupCollectionDefinitionVisibility_Template),
	}
}

func (s *DeviceManagementConfigurationSettingGroupCollectionDefinitionVisibility) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationSettingGroupCollectionDefinitionVisibility(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationSettingGroupCollectionDefinitionVisibility(input string) (*DeviceManagementConfigurationSettingGroupCollectionDefinitionVisibility, error) {
	vals := map[string]DeviceManagementConfigurationSettingGroupCollectionDefinitionVisibility{
		"none":            DeviceManagementConfigurationSettingGroupCollectionDefinitionVisibility_None,
		"settingscatalog": DeviceManagementConfigurationSettingGroupCollectionDefinitionVisibility_SettingsCatalog,
		"template":        DeviceManagementConfigurationSettingGroupCollectionDefinitionVisibility_Template,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationSettingGroupCollectionDefinitionVisibility(input)
	return &out, nil
}
