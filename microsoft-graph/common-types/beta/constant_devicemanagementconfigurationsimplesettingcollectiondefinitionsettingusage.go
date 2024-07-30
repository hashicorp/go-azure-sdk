package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSimpleSettingCollectionDefinitionSettingUsage string

const (
	DeviceManagementConfigurationSimpleSettingCollectionDefinitionSettingUsage_Compliance    DeviceManagementConfigurationSimpleSettingCollectionDefinitionSettingUsage = "compliance"
	DeviceManagementConfigurationSimpleSettingCollectionDefinitionSettingUsage_Configuration DeviceManagementConfigurationSimpleSettingCollectionDefinitionSettingUsage = "configuration"
	DeviceManagementConfigurationSimpleSettingCollectionDefinitionSettingUsage_None          DeviceManagementConfigurationSimpleSettingCollectionDefinitionSettingUsage = "none"
)

func PossibleValuesForDeviceManagementConfigurationSimpleSettingCollectionDefinitionSettingUsage() []string {
	return []string{
		string(DeviceManagementConfigurationSimpleSettingCollectionDefinitionSettingUsage_Compliance),
		string(DeviceManagementConfigurationSimpleSettingCollectionDefinitionSettingUsage_Configuration),
		string(DeviceManagementConfigurationSimpleSettingCollectionDefinitionSettingUsage_None),
	}
}

func (s *DeviceManagementConfigurationSimpleSettingCollectionDefinitionSettingUsage) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationSimpleSettingCollectionDefinitionSettingUsage(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationSimpleSettingCollectionDefinitionSettingUsage(input string) (*DeviceManagementConfigurationSimpleSettingCollectionDefinitionSettingUsage, error) {
	vals := map[string]DeviceManagementConfigurationSimpleSettingCollectionDefinitionSettingUsage{
		"compliance":    DeviceManagementConfigurationSimpleSettingCollectionDefinitionSettingUsage_Compliance,
		"configuration": DeviceManagementConfigurationSimpleSettingCollectionDefinitionSettingUsage_Configuration,
		"none":          DeviceManagementConfigurationSimpleSettingCollectionDefinitionSettingUsage_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationSimpleSettingCollectionDefinitionSettingUsage(input)
	return &out, nil
}
