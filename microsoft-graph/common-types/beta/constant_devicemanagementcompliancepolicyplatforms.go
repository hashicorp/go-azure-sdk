package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementCompliancePolicyPlatforms string

const (
	DeviceManagementCompliancePolicyPlatforms_Android    DeviceManagementCompliancePolicyPlatforms = "android"
	DeviceManagementCompliancePolicyPlatforms_IOS        DeviceManagementCompliancePolicyPlatforms = "iOS"
	DeviceManagementCompliancePolicyPlatforms_Linux      DeviceManagementCompliancePolicyPlatforms = "linux"
	DeviceManagementCompliancePolicyPlatforms_MacOS      DeviceManagementCompliancePolicyPlatforms = "macOS"
	DeviceManagementCompliancePolicyPlatforms_None       DeviceManagementCompliancePolicyPlatforms = "none"
	DeviceManagementCompliancePolicyPlatforms_Windows10  DeviceManagementCompliancePolicyPlatforms = "windows10"
	DeviceManagementCompliancePolicyPlatforms_Windows10X DeviceManagementCompliancePolicyPlatforms = "windows10X"
)

func PossibleValuesForDeviceManagementCompliancePolicyPlatforms() []string {
	return []string{
		string(DeviceManagementCompliancePolicyPlatforms_Android),
		string(DeviceManagementCompliancePolicyPlatforms_IOS),
		string(DeviceManagementCompliancePolicyPlatforms_Linux),
		string(DeviceManagementCompliancePolicyPlatforms_MacOS),
		string(DeviceManagementCompliancePolicyPlatforms_None),
		string(DeviceManagementCompliancePolicyPlatforms_Windows10),
		string(DeviceManagementCompliancePolicyPlatforms_Windows10X),
	}
}

func (s *DeviceManagementCompliancePolicyPlatforms) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementCompliancePolicyPlatforms(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementCompliancePolicyPlatforms(input string) (*DeviceManagementCompliancePolicyPlatforms, error) {
	vals := map[string]DeviceManagementCompliancePolicyPlatforms{
		"android":    DeviceManagementCompliancePolicyPlatforms_Android,
		"ios":        DeviceManagementCompliancePolicyPlatforms_IOS,
		"linux":      DeviceManagementCompliancePolicyPlatforms_Linux,
		"macos":      DeviceManagementCompliancePolicyPlatforms_MacOS,
		"none":       DeviceManagementCompliancePolicyPlatforms_None,
		"windows10":  DeviceManagementCompliancePolicyPlatforms_Windows10,
		"windows10x": DeviceManagementCompliancePolicyPlatforms_Windows10X,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementCompliancePolicyPlatforms(input)
	return &out, nil
}
