package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSettingGroupCollectionDefinitionSettingUsage string

const (
	DeviceManagementConfigurationSettingGroupCollectionDefinitionSettingUsage_Compliance    DeviceManagementConfigurationSettingGroupCollectionDefinitionSettingUsage = "compliance"
	DeviceManagementConfigurationSettingGroupCollectionDefinitionSettingUsage_Configuration DeviceManagementConfigurationSettingGroupCollectionDefinitionSettingUsage = "configuration"
	DeviceManagementConfigurationSettingGroupCollectionDefinitionSettingUsage_None          DeviceManagementConfigurationSettingGroupCollectionDefinitionSettingUsage = "none"
)

func PossibleValuesForDeviceManagementConfigurationSettingGroupCollectionDefinitionSettingUsage() []string {
	return []string{
		string(DeviceManagementConfigurationSettingGroupCollectionDefinitionSettingUsage_Compliance),
		string(DeviceManagementConfigurationSettingGroupCollectionDefinitionSettingUsage_Configuration),
		string(DeviceManagementConfigurationSettingGroupCollectionDefinitionSettingUsage_None),
	}
}

func (s *DeviceManagementConfigurationSettingGroupCollectionDefinitionSettingUsage) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationSettingGroupCollectionDefinitionSettingUsage(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationSettingGroupCollectionDefinitionSettingUsage(input string) (*DeviceManagementConfigurationSettingGroupCollectionDefinitionSettingUsage, error) {
	vals := map[string]DeviceManagementConfigurationSettingGroupCollectionDefinitionSettingUsage{
		"compliance":    DeviceManagementConfigurationSettingGroupCollectionDefinitionSettingUsage_Compliance,
		"configuration": DeviceManagementConfigurationSettingGroupCollectionDefinitionSettingUsage_Configuration,
		"none":          DeviceManagementConfigurationSettingGroupCollectionDefinitionSettingUsage_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationSettingGroupCollectionDefinitionSettingUsage(input)
	return &out, nil
}
