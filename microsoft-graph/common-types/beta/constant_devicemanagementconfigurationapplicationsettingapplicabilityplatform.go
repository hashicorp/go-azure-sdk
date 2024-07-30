package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationApplicationSettingApplicabilityPlatform string

const (
	DeviceManagementConfigurationApplicationSettingApplicabilityPlatform_Android    DeviceManagementConfigurationApplicationSettingApplicabilityPlatform = "android"
	DeviceManagementConfigurationApplicationSettingApplicabilityPlatform_IOS        DeviceManagementConfigurationApplicationSettingApplicabilityPlatform = "iOS"
	DeviceManagementConfigurationApplicationSettingApplicabilityPlatform_Linux      DeviceManagementConfigurationApplicationSettingApplicabilityPlatform = "linux"
	DeviceManagementConfigurationApplicationSettingApplicabilityPlatform_MacOS      DeviceManagementConfigurationApplicationSettingApplicabilityPlatform = "macOS"
	DeviceManagementConfigurationApplicationSettingApplicabilityPlatform_None       DeviceManagementConfigurationApplicationSettingApplicabilityPlatform = "none"
	DeviceManagementConfigurationApplicationSettingApplicabilityPlatform_Windows10  DeviceManagementConfigurationApplicationSettingApplicabilityPlatform = "windows10"
	DeviceManagementConfigurationApplicationSettingApplicabilityPlatform_Windows10X DeviceManagementConfigurationApplicationSettingApplicabilityPlatform = "windows10X"
)

func PossibleValuesForDeviceManagementConfigurationApplicationSettingApplicabilityPlatform() []string {
	return []string{
		string(DeviceManagementConfigurationApplicationSettingApplicabilityPlatform_Android),
		string(DeviceManagementConfigurationApplicationSettingApplicabilityPlatform_IOS),
		string(DeviceManagementConfigurationApplicationSettingApplicabilityPlatform_Linux),
		string(DeviceManagementConfigurationApplicationSettingApplicabilityPlatform_MacOS),
		string(DeviceManagementConfigurationApplicationSettingApplicabilityPlatform_None),
		string(DeviceManagementConfigurationApplicationSettingApplicabilityPlatform_Windows10),
		string(DeviceManagementConfigurationApplicationSettingApplicabilityPlatform_Windows10X),
	}
}

func (s *DeviceManagementConfigurationApplicationSettingApplicabilityPlatform) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationApplicationSettingApplicabilityPlatform(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationApplicationSettingApplicabilityPlatform(input string) (*DeviceManagementConfigurationApplicationSettingApplicabilityPlatform, error) {
	vals := map[string]DeviceManagementConfigurationApplicationSettingApplicabilityPlatform{
		"android":    DeviceManagementConfigurationApplicationSettingApplicabilityPlatform_Android,
		"ios":        DeviceManagementConfigurationApplicationSettingApplicabilityPlatform_IOS,
		"linux":      DeviceManagementConfigurationApplicationSettingApplicabilityPlatform_Linux,
		"macos":      DeviceManagementConfigurationApplicationSettingApplicabilityPlatform_MacOS,
		"none":       DeviceManagementConfigurationApplicationSettingApplicabilityPlatform_None,
		"windows10":  DeviceManagementConfigurationApplicationSettingApplicabilityPlatform_Windows10,
		"windows10x": DeviceManagementConfigurationApplicationSettingApplicabilityPlatform_Windows10X,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationApplicationSettingApplicabilityPlatform(input)
	return &out, nil
}
