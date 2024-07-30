package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSettingApplicabilityPlatform string

const (
	DeviceManagementConfigurationSettingApplicabilityPlatform_Android    DeviceManagementConfigurationSettingApplicabilityPlatform = "android"
	DeviceManagementConfigurationSettingApplicabilityPlatform_IOS        DeviceManagementConfigurationSettingApplicabilityPlatform = "iOS"
	DeviceManagementConfigurationSettingApplicabilityPlatform_Linux      DeviceManagementConfigurationSettingApplicabilityPlatform = "linux"
	DeviceManagementConfigurationSettingApplicabilityPlatform_MacOS      DeviceManagementConfigurationSettingApplicabilityPlatform = "macOS"
	DeviceManagementConfigurationSettingApplicabilityPlatform_None       DeviceManagementConfigurationSettingApplicabilityPlatform = "none"
	DeviceManagementConfigurationSettingApplicabilityPlatform_Windows10  DeviceManagementConfigurationSettingApplicabilityPlatform = "windows10"
	DeviceManagementConfigurationSettingApplicabilityPlatform_Windows10X DeviceManagementConfigurationSettingApplicabilityPlatform = "windows10X"
)

func PossibleValuesForDeviceManagementConfigurationSettingApplicabilityPlatform() []string {
	return []string{
		string(DeviceManagementConfigurationSettingApplicabilityPlatform_Android),
		string(DeviceManagementConfigurationSettingApplicabilityPlatform_IOS),
		string(DeviceManagementConfigurationSettingApplicabilityPlatform_Linux),
		string(DeviceManagementConfigurationSettingApplicabilityPlatform_MacOS),
		string(DeviceManagementConfigurationSettingApplicabilityPlatform_None),
		string(DeviceManagementConfigurationSettingApplicabilityPlatform_Windows10),
		string(DeviceManagementConfigurationSettingApplicabilityPlatform_Windows10X),
	}
}

func (s *DeviceManagementConfigurationSettingApplicabilityPlatform) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationSettingApplicabilityPlatform(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationSettingApplicabilityPlatform(input string) (*DeviceManagementConfigurationSettingApplicabilityPlatform, error) {
	vals := map[string]DeviceManagementConfigurationSettingApplicabilityPlatform{
		"android":    DeviceManagementConfigurationSettingApplicabilityPlatform_Android,
		"ios":        DeviceManagementConfigurationSettingApplicabilityPlatform_IOS,
		"linux":      DeviceManagementConfigurationSettingApplicabilityPlatform_Linux,
		"macos":      DeviceManagementConfigurationSettingApplicabilityPlatform_MacOS,
		"none":       DeviceManagementConfigurationSettingApplicabilityPlatform_None,
		"windows10":  DeviceManagementConfigurationSettingApplicabilityPlatform_Windows10,
		"windows10x": DeviceManagementConfigurationSettingApplicabilityPlatform_Windows10X,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationSettingApplicabilityPlatform(input)
	return &out, nil
}
