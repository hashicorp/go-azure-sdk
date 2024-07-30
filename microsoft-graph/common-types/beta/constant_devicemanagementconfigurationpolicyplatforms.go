package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationPolicyPlatforms string

const (
	DeviceManagementConfigurationPolicyPlatforms_Android    DeviceManagementConfigurationPolicyPlatforms = "android"
	DeviceManagementConfigurationPolicyPlatforms_IOS        DeviceManagementConfigurationPolicyPlatforms = "iOS"
	DeviceManagementConfigurationPolicyPlatforms_Linux      DeviceManagementConfigurationPolicyPlatforms = "linux"
	DeviceManagementConfigurationPolicyPlatforms_MacOS      DeviceManagementConfigurationPolicyPlatforms = "macOS"
	DeviceManagementConfigurationPolicyPlatforms_None       DeviceManagementConfigurationPolicyPlatforms = "none"
	DeviceManagementConfigurationPolicyPlatforms_Windows10  DeviceManagementConfigurationPolicyPlatforms = "windows10"
	DeviceManagementConfigurationPolicyPlatforms_Windows10X DeviceManagementConfigurationPolicyPlatforms = "windows10X"
)

func PossibleValuesForDeviceManagementConfigurationPolicyPlatforms() []string {
	return []string{
		string(DeviceManagementConfigurationPolicyPlatforms_Android),
		string(DeviceManagementConfigurationPolicyPlatforms_IOS),
		string(DeviceManagementConfigurationPolicyPlatforms_Linux),
		string(DeviceManagementConfigurationPolicyPlatforms_MacOS),
		string(DeviceManagementConfigurationPolicyPlatforms_None),
		string(DeviceManagementConfigurationPolicyPlatforms_Windows10),
		string(DeviceManagementConfigurationPolicyPlatforms_Windows10X),
	}
}

func (s *DeviceManagementConfigurationPolicyPlatforms) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationPolicyPlatforms(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationPolicyPlatforms(input string) (*DeviceManagementConfigurationPolicyPlatforms, error) {
	vals := map[string]DeviceManagementConfigurationPolicyPlatforms{
		"android":    DeviceManagementConfigurationPolicyPlatforms_Android,
		"ios":        DeviceManagementConfigurationPolicyPlatforms_IOS,
		"linux":      DeviceManagementConfigurationPolicyPlatforms_Linux,
		"macos":      DeviceManagementConfigurationPolicyPlatforms_MacOS,
		"none":       DeviceManagementConfigurationPolicyPlatforms_None,
		"windows10":  DeviceManagementConfigurationPolicyPlatforms_Windows10,
		"windows10x": DeviceManagementConfigurationPolicyPlatforms_Windows10X,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationPolicyPlatforms(input)
	return &out, nil
}
