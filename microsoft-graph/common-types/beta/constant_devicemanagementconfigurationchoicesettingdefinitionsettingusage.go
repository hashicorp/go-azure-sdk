package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationChoiceSettingDefinitionSettingUsage string

const (
	DeviceManagementConfigurationChoiceSettingDefinitionSettingUsage_Compliance    DeviceManagementConfigurationChoiceSettingDefinitionSettingUsage = "compliance"
	DeviceManagementConfigurationChoiceSettingDefinitionSettingUsage_Configuration DeviceManagementConfigurationChoiceSettingDefinitionSettingUsage = "configuration"
	DeviceManagementConfigurationChoiceSettingDefinitionSettingUsage_None          DeviceManagementConfigurationChoiceSettingDefinitionSettingUsage = "none"
)

func PossibleValuesForDeviceManagementConfigurationChoiceSettingDefinitionSettingUsage() []string {
	return []string{
		string(DeviceManagementConfigurationChoiceSettingDefinitionSettingUsage_Compliance),
		string(DeviceManagementConfigurationChoiceSettingDefinitionSettingUsage_Configuration),
		string(DeviceManagementConfigurationChoiceSettingDefinitionSettingUsage_None),
	}
}

func (s *DeviceManagementConfigurationChoiceSettingDefinitionSettingUsage) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationChoiceSettingDefinitionSettingUsage(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationChoiceSettingDefinitionSettingUsage(input string) (*DeviceManagementConfigurationChoiceSettingDefinitionSettingUsage, error) {
	vals := map[string]DeviceManagementConfigurationChoiceSettingDefinitionSettingUsage{
		"compliance":    DeviceManagementConfigurationChoiceSettingDefinitionSettingUsage_Compliance,
		"configuration": DeviceManagementConfigurationChoiceSettingDefinitionSettingUsage_Configuration,
		"none":          DeviceManagementConfigurationChoiceSettingDefinitionSettingUsage_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationChoiceSettingDefinitionSettingUsage(input)
	return &out, nil
}
