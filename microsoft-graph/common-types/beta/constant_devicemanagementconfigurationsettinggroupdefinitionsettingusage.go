package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSettingGroupDefinitionSettingUsage string

const (
	DeviceManagementConfigurationSettingGroupDefinitionSettingUsage_Compliance    DeviceManagementConfigurationSettingGroupDefinitionSettingUsage = "compliance"
	DeviceManagementConfigurationSettingGroupDefinitionSettingUsage_Configuration DeviceManagementConfigurationSettingGroupDefinitionSettingUsage = "configuration"
	DeviceManagementConfigurationSettingGroupDefinitionSettingUsage_None          DeviceManagementConfigurationSettingGroupDefinitionSettingUsage = "none"
)

func PossibleValuesForDeviceManagementConfigurationSettingGroupDefinitionSettingUsage() []string {
	return []string{
		string(DeviceManagementConfigurationSettingGroupDefinitionSettingUsage_Compliance),
		string(DeviceManagementConfigurationSettingGroupDefinitionSettingUsage_Configuration),
		string(DeviceManagementConfigurationSettingGroupDefinitionSettingUsage_None),
	}
}

func (s *DeviceManagementConfigurationSettingGroupDefinitionSettingUsage) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationSettingGroupDefinitionSettingUsage(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationSettingGroupDefinitionSettingUsage(input string) (*DeviceManagementConfigurationSettingGroupDefinitionSettingUsage, error) {
	vals := map[string]DeviceManagementConfigurationSettingGroupDefinitionSettingUsage{
		"compliance":    DeviceManagementConfigurationSettingGroupDefinitionSettingUsage_Compliance,
		"configuration": DeviceManagementConfigurationSettingGroupDefinitionSettingUsage_Configuration,
		"none":          DeviceManagementConfigurationSettingGroupDefinitionSettingUsage_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationSettingGroupDefinitionSettingUsage(input)
	return &out, nil
}
