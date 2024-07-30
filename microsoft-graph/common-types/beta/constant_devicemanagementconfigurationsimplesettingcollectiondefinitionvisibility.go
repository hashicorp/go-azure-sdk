package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSimpleSettingCollectionDefinitionVisibility string

const (
	DeviceManagementConfigurationSimpleSettingCollectionDefinitionVisibility_None            DeviceManagementConfigurationSimpleSettingCollectionDefinitionVisibility = "none"
	DeviceManagementConfigurationSimpleSettingCollectionDefinitionVisibility_SettingsCatalog DeviceManagementConfigurationSimpleSettingCollectionDefinitionVisibility = "settingsCatalog"
	DeviceManagementConfigurationSimpleSettingCollectionDefinitionVisibility_Template        DeviceManagementConfigurationSimpleSettingCollectionDefinitionVisibility = "template"
)

func PossibleValuesForDeviceManagementConfigurationSimpleSettingCollectionDefinitionVisibility() []string {
	return []string{
		string(DeviceManagementConfigurationSimpleSettingCollectionDefinitionVisibility_None),
		string(DeviceManagementConfigurationSimpleSettingCollectionDefinitionVisibility_SettingsCatalog),
		string(DeviceManagementConfigurationSimpleSettingCollectionDefinitionVisibility_Template),
	}
}

func (s *DeviceManagementConfigurationSimpleSettingCollectionDefinitionVisibility) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationSimpleSettingCollectionDefinitionVisibility(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationSimpleSettingCollectionDefinitionVisibility(input string) (*DeviceManagementConfigurationSimpleSettingCollectionDefinitionVisibility, error) {
	vals := map[string]DeviceManagementConfigurationSimpleSettingCollectionDefinitionVisibility{
		"none":            DeviceManagementConfigurationSimpleSettingCollectionDefinitionVisibility_None,
		"settingscatalog": DeviceManagementConfigurationSimpleSettingCollectionDefinitionVisibility_SettingsCatalog,
		"template":        DeviceManagementConfigurationSimpleSettingCollectionDefinitionVisibility_Template,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationSimpleSettingCollectionDefinitionVisibility(input)
	return &out, nil
}
