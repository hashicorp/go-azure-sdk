package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationChoiceSettingCollectionDefinitionSettingUsage string

const (
	DeviceManagementConfigurationChoiceSettingCollectionDefinitionSettingUsage_Compliance    DeviceManagementConfigurationChoiceSettingCollectionDefinitionSettingUsage = "compliance"
	DeviceManagementConfigurationChoiceSettingCollectionDefinitionSettingUsage_Configuration DeviceManagementConfigurationChoiceSettingCollectionDefinitionSettingUsage = "configuration"
	DeviceManagementConfigurationChoiceSettingCollectionDefinitionSettingUsage_None          DeviceManagementConfigurationChoiceSettingCollectionDefinitionSettingUsage = "none"
)

func PossibleValuesForDeviceManagementConfigurationChoiceSettingCollectionDefinitionSettingUsage() []string {
	return []string{
		string(DeviceManagementConfigurationChoiceSettingCollectionDefinitionSettingUsage_Compliance),
		string(DeviceManagementConfigurationChoiceSettingCollectionDefinitionSettingUsage_Configuration),
		string(DeviceManagementConfigurationChoiceSettingCollectionDefinitionSettingUsage_None),
	}
}

func (s *DeviceManagementConfigurationChoiceSettingCollectionDefinitionSettingUsage) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationChoiceSettingCollectionDefinitionSettingUsage(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationChoiceSettingCollectionDefinitionSettingUsage(input string) (*DeviceManagementConfigurationChoiceSettingCollectionDefinitionSettingUsage, error) {
	vals := map[string]DeviceManagementConfigurationChoiceSettingCollectionDefinitionSettingUsage{
		"compliance":    DeviceManagementConfigurationChoiceSettingCollectionDefinitionSettingUsage_Compliance,
		"configuration": DeviceManagementConfigurationChoiceSettingCollectionDefinitionSettingUsage_Configuration,
		"none":          DeviceManagementConfigurationChoiceSettingCollectionDefinitionSettingUsage_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationChoiceSettingCollectionDefinitionSettingUsage(input)
	return &out, nil
}
