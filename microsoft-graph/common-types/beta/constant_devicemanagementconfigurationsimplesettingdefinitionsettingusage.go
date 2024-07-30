package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSimpleSettingDefinitionSettingUsage string

const (
	DeviceManagementConfigurationSimpleSettingDefinitionSettingUsage_Compliance    DeviceManagementConfigurationSimpleSettingDefinitionSettingUsage = "compliance"
	DeviceManagementConfigurationSimpleSettingDefinitionSettingUsage_Configuration DeviceManagementConfigurationSimpleSettingDefinitionSettingUsage = "configuration"
	DeviceManagementConfigurationSimpleSettingDefinitionSettingUsage_None          DeviceManagementConfigurationSimpleSettingDefinitionSettingUsage = "none"
)

func PossibleValuesForDeviceManagementConfigurationSimpleSettingDefinitionSettingUsage() []string {
	return []string{
		string(DeviceManagementConfigurationSimpleSettingDefinitionSettingUsage_Compliance),
		string(DeviceManagementConfigurationSimpleSettingDefinitionSettingUsage_Configuration),
		string(DeviceManagementConfigurationSimpleSettingDefinitionSettingUsage_None),
	}
}

func (s *DeviceManagementConfigurationSimpleSettingDefinitionSettingUsage) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationSimpleSettingDefinitionSettingUsage(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationSimpleSettingDefinitionSettingUsage(input string) (*DeviceManagementConfigurationSimpleSettingDefinitionSettingUsage, error) {
	vals := map[string]DeviceManagementConfigurationSimpleSettingDefinitionSettingUsage{
		"compliance":    DeviceManagementConfigurationSimpleSettingDefinitionSettingUsage_Compliance,
		"configuration": DeviceManagementConfigurationSimpleSettingDefinitionSettingUsage_Configuration,
		"none":          DeviceManagementConfigurationSimpleSettingDefinitionSettingUsage_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationSimpleSettingDefinitionSettingUsage(input)
	return &out, nil
}
