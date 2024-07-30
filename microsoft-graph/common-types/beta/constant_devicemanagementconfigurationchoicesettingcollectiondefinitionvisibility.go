package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationChoiceSettingCollectionDefinitionVisibility string

const (
	DeviceManagementConfigurationChoiceSettingCollectionDefinitionVisibility_None            DeviceManagementConfigurationChoiceSettingCollectionDefinitionVisibility = "none"
	DeviceManagementConfigurationChoiceSettingCollectionDefinitionVisibility_SettingsCatalog DeviceManagementConfigurationChoiceSettingCollectionDefinitionVisibility = "settingsCatalog"
	DeviceManagementConfigurationChoiceSettingCollectionDefinitionVisibility_Template        DeviceManagementConfigurationChoiceSettingCollectionDefinitionVisibility = "template"
)

func PossibleValuesForDeviceManagementConfigurationChoiceSettingCollectionDefinitionVisibility() []string {
	return []string{
		string(DeviceManagementConfigurationChoiceSettingCollectionDefinitionVisibility_None),
		string(DeviceManagementConfigurationChoiceSettingCollectionDefinitionVisibility_SettingsCatalog),
		string(DeviceManagementConfigurationChoiceSettingCollectionDefinitionVisibility_Template),
	}
}

func (s *DeviceManagementConfigurationChoiceSettingCollectionDefinitionVisibility) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationChoiceSettingCollectionDefinitionVisibility(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationChoiceSettingCollectionDefinitionVisibility(input string) (*DeviceManagementConfigurationChoiceSettingCollectionDefinitionVisibility, error) {
	vals := map[string]DeviceManagementConfigurationChoiceSettingCollectionDefinitionVisibility{
		"none":            DeviceManagementConfigurationChoiceSettingCollectionDefinitionVisibility_None,
		"settingscatalog": DeviceManagementConfigurationChoiceSettingCollectionDefinitionVisibility_SettingsCatalog,
		"template":        DeviceManagementConfigurationChoiceSettingCollectionDefinitionVisibility_Template,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationChoiceSettingCollectionDefinitionVisibility(input)
	return &out, nil
}
