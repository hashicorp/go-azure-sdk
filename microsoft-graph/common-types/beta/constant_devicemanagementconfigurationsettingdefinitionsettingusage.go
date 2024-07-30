package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSettingDefinitionSettingUsage string

const (
	DeviceManagementConfigurationSettingDefinitionSettingUsage_Compliance    DeviceManagementConfigurationSettingDefinitionSettingUsage = "compliance"
	DeviceManagementConfigurationSettingDefinitionSettingUsage_Configuration DeviceManagementConfigurationSettingDefinitionSettingUsage = "configuration"
	DeviceManagementConfigurationSettingDefinitionSettingUsage_None          DeviceManagementConfigurationSettingDefinitionSettingUsage = "none"
)

func PossibleValuesForDeviceManagementConfigurationSettingDefinitionSettingUsage() []string {
	return []string{
		string(DeviceManagementConfigurationSettingDefinitionSettingUsage_Compliance),
		string(DeviceManagementConfigurationSettingDefinitionSettingUsage_Configuration),
		string(DeviceManagementConfigurationSettingDefinitionSettingUsage_None),
	}
}

func (s *DeviceManagementConfigurationSettingDefinitionSettingUsage) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationSettingDefinitionSettingUsage(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationSettingDefinitionSettingUsage(input string) (*DeviceManagementConfigurationSettingDefinitionSettingUsage, error) {
	vals := map[string]DeviceManagementConfigurationSettingDefinitionSettingUsage{
		"compliance":    DeviceManagementConfigurationSettingDefinitionSettingUsage_Compliance,
		"configuration": DeviceManagementConfigurationSettingDefinitionSettingUsage_Configuration,
		"none":          DeviceManagementConfigurationSettingDefinitionSettingUsage_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationSettingDefinitionSettingUsage(input)
	return &out, nil
}
