package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationRedirectSettingDefinitionSettingUsage string

const (
	DeviceManagementConfigurationRedirectSettingDefinitionSettingUsage_Compliance    DeviceManagementConfigurationRedirectSettingDefinitionSettingUsage = "compliance"
	DeviceManagementConfigurationRedirectSettingDefinitionSettingUsage_Configuration DeviceManagementConfigurationRedirectSettingDefinitionSettingUsage = "configuration"
	DeviceManagementConfigurationRedirectSettingDefinitionSettingUsage_None          DeviceManagementConfigurationRedirectSettingDefinitionSettingUsage = "none"
)

func PossibleValuesForDeviceManagementConfigurationRedirectSettingDefinitionSettingUsage() []string {
	return []string{
		string(DeviceManagementConfigurationRedirectSettingDefinitionSettingUsage_Compliance),
		string(DeviceManagementConfigurationRedirectSettingDefinitionSettingUsage_Configuration),
		string(DeviceManagementConfigurationRedirectSettingDefinitionSettingUsage_None),
	}
}

func (s *DeviceManagementConfigurationRedirectSettingDefinitionSettingUsage) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationRedirectSettingDefinitionSettingUsage(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationRedirectSettingDefinitionSettingUsage(input string) (*DeviceManagementConfigurationRedirectSettingDefinitionSettingUsage, error) {
	vals := map[string]DeviceManagementConfigurationRedirectSettingDefinitionSettingUsage{
		"compliance":    DeviceManagementConfigurationRedirectSettingDefinitionSettingUsage_Compliance,
		"configuration": DeviceManagementConfigurationRedirectSettingDefinitionSettingUsage_Configuration,
		"none":          DeviceManagementConfigurationRedirectSettingDefinitionSettingUsage_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationRedirectSettingDefinitionSettingUsage(input)
	return &out, nil
}
