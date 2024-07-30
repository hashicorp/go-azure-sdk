package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationPolicyTemplatePlatforms string

const (
	DeviceManagementConfigurationPolicyTemplatePlatforms_Android    DeviceManagementConfigurationPolicyTemplatePlatforms = "android"
	DeviceManagementConfigurationPolicyTemplatePlatforms_IOS        DeviceManagementConfigurationPolicyTemplatePlatforms = "iOS"
	DeviceManagementConfigurationPolicyTemplatePlatforms_Linux      DeviceManagementConfigurationPolicyTemplatePlatforms = "linux"
	DeviceManagementConfigurationPolicyTemplatePlatforms_MacOS      DeviceManagementConfigurationPolicyTemplatePlatforms = "macOS"
	DeviceManagementConfigurationPolicyTemplatePlatforms_None       DeviceManagementConfigurationPolicyTemplatePlatforms = "none"
	DeviceManagementConfigurationPolicyTemplatePlatforms_Windows10  DeviceManagementConfigurationPolicyTemplatePlatforms = "windows10"
	DeviceManagementConfigurationPolicyTemplatePlatforms_Windows10X DeviceManagementConfigurationPolicyTemplatePlatforms = "windows10X"
)

func PossibleValuesForDeviceManagementConfigurationPolicyTemplatePlatforms() []string {
	return []string{
		string(DeviceManagementConfigurationPolicyTemplatePlatforms_Android),
		string(DeviceManagementConfigurationPolicyTemplatePlatforms_IOS),
		string(DeviceManagementConfigurationPolicyTemplatePlatforms_Linux),
		string(DeviceManagementConfigurationPolicyTemplatePlatforms_MacOS),
		string(DeviceManagementConfigurationPolicyTemplatePlatforms_None),
		string(DeviceManagementConfigurationPolicyTemplatePlatforms_Windows10),
		string(DeviceManagementConfigurationPolicyTemplatePlatforms_Windows10X),
	}
}

func (s *DeviceManagementConfigurationPolicyTemplatePlatforms) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationPolicyTemplatePlatforms(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationPolicyTemplatePlatforms(input string) (*DeviceManagementConfigurationPolicyTemplatePlatforms, error) {
	vals := map[string]DeviceManagementConfigurationPolicyTemplatePlatforms{
		"android":    DeviceManagementConfigurationPolicyTemplatePlatforms_Android,
		"ios":        DeviceManagementConfigurationPolicyTemplatePlatforms_IOS,
		"linux":      DeviceManagementConfigurationPolicyTemplatePlatforms_Linux,
		"macos":      DeviceManagementConfigurationPolicyTemplatePlatforms_MacOS,
		"none":       DeviceManagementConfigurationPolicyTemplatePlatforms_None,
		"windows10":  DeviceManagementConfigurationPolicyTemplatePlatforms_Windows10,
		"windows10x": DeviceManagementConfigurationPolicyTemplatePlatforms_Windows10X,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationPolicyTemplatePlatforms(input)
	return &out, nil
}
