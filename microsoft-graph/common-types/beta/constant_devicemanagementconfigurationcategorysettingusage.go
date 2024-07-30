package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationCategorySettingUsage string

const (
	DeviceManagementConfigurationCategorySettingUsage_Compliance    DeviceManagementConfigurationCategorySettingUsage = "compliance"
	DeviceManagementConfigurationCategorySettingUsage_Configuration DeviceManagementConfigurationCategorySettingUsage = "configuration"
	DeviceManagementConfigurationCategorySettingUsage_None          DeviceManagementConfigurationCategorySettingUsage = "none"
)

func PossibleValuesForDeviceManagementConfigurationCategorySettingUsage() []string {
	return []string{
		string(DeviceManagementConfigurationCategorySettingUsage_Compliance),
		string(DeviceManagementConfigurationCategorySettingUsage_Configuration),
		string(DeviceManagementConfigurationCategorySettingUsage_None),
	}
}

func (s *DeviceManagementConfigurationCategorySettingUsage) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationCategorySettingUsage(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationCategorySettingUsage(input string) (*DeviceManagementConfigurationCategorySettingUsage, error) {
	vals := map[string]DeviceManagementConfigurationCategorySettingUsage{
		"compliance":    DeviceManagementConfigurationCategorySettingUsage_Compliance,
		"configuration": DeviceManagementConfigurationCategorySettingUsage_Configuration,
		"none":          DeviceManagementConfigurationCategorySettingUsage_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationCategorySettingUsage(input)
	return &out, nil
}
