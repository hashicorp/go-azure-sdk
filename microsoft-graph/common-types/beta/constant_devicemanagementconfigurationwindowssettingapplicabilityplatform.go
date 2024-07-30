package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationWindowsSettingApplicabilityPlatform string

const (
	DeviceManagementConfigurationWindowsSettingApplicabilityPlatform_Android    DeviceManagementConfigurationWindowsSettingApplicabilityPlatform = "android"
	DeviceManagementConfigurationWindowsSettingApplicabilityPlatform_IOS        DeviceManagementConfigurationWindowsSettingApplicabilityPlatform = "iOS"
	DeviceManagementConfigurationWindowsSettingApplicabilityPlatform_Linux      DeviceManagementConfigurationWindowsSettingApplicabilityPlatform = "linux"
	DeviceManagementConfigurationWindowsSettingApplicabilityPlatform_MacOS      DeviceManagementConfigurationWindowsSettingApplicabilityPlatform = "macOS"
	DeviceManagementConfigurationWindowsSettingApplicabilityPlatform_None       DeviceManagementConfigurationWindowsSettingApplicabilityPlatform = "none"
	DeviceManagementConfigurationWindowsSettingApplicabilityPlatform_Windows10  DeviceManagementConfigurationWindowsSettingApplicabilityPlatform = "windows10"
	DeviceManagementConfigurationWindowsSettingApplicabilityPlatform_Windows10X DeviceManagementConfigurationWindowsSettingApplicabilityPlatform = "windows10X"
)

func PossibleValuesForDeviceManagementConfigurationWindowsSettingApplicabilityPlatform() []string {
	return []string{
		string(DeviceManagementConfigurationWindowsSettingApplicabilityPlatform_Android),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityPlatform_IOS),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityPlatform_Linux),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityPlatform_MacOS),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityPlatform_None),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityPlatform_Windows10),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityPlatform_Windows10X),
	}
}

func (s *DeviceManagementConfigurationWindowsSettingApplicabilityPlatform) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationWindowsSettingApplicabilityPlatform(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationWindowsSettingApplicabilityPlatform(input string) (*DeviceManagementConfigurationWindowsSettingApplicabilityPlatform, error) {
	vals := map[string]DeviceManagementConfigurationWindowsSettingApplicabilityPlatform{
		"android":    DeviceManagementConfigurationWindowsSettingApplicabilityPlatform_Android,
		"ios":        DeviceManagementConfigurationWindowsSettingApplicabilityPlatform_IOS,
		"linux":      DeviceManagementConfigurationWindowsSettingApplicabilityPlatform_Linux,
		"macos":      DeviceManagementConfigurationWindowsSettingApplicabilityPlatform_MacOS,
		"none":       DeviceManagementConfigurationWindowsSettingApplicabilityPlatform_None,
		"windows10":  DeviceManagementConfigurationWindowsSettingApplicabilityPlatform_Windows10,
		"windows10x": DeviceManagementConfigurationWindowsSettingApplicabilityPlatform_Windows10X,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationWindowsSettingApplicabilityPlatform(input)
	return &out, nil
}
