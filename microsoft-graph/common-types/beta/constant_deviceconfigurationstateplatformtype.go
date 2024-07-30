package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceConfigurationStatePlatformType string

const (
	DeviceConfigurationStatePlatformType_All                DeviceConfigurationStatePlatformType = "all"
	DeviceConfigurationStatePlatformType_Android            DeviceConfigurationStatePlatformType = "android"
	DeviceConfigurationStatePlatformType_AndroidAOSP        DeviceConfigurationStatePlatformType = "androidAOSP"
	DeviceConfigurationStatePlatformType_AndroidForWork     DeviceConfigurationStatePlatformType = "androidForWork"
	DeviceConfigurationStatePlatformType_AndroidWorkProfile DeviceConfigurationStatePlatformType = "androidWorkProfile"
	DeviceConfigurationStatePlatformType_IOS                DeviceConfigurationStatePlatformType = "iOS"
	DeviceConfigurationStatePlatformType_MacOS              DeviceConfigurationStatePlatformType = "macOS"
	DeviceConfigurationStatePlatformType_Windows10AndLater  DeviceConfigurationStatePlatformType = "windows10AndLater"
	DeviceConfigurationStatePlatformType_Windows10XProfile  DeviceConfigurationStatePlatformType = "windows10XProfile"
	DeviceConfigurationStatePlatformType_Windows81AndLater  DeviceConfigurationStatePlatformType = "windows81AndLater"
	DeviceConfigurationStatePlatformType_WindowsPhone81     DeviceConfigurationStatePlatformType = "windowsPhone81"
)

func PossibleValuesForDeviceConfigurationStatePlatformType() []string {
	return []string{
		string(DeviceConfigurationStatePlatformType_All),
		string(DeviceConfigurationStatePlatformType_Android),
		string(DeviceConfigurationStatePlatformType_AndroidAOSP),
		string(DeviceConfigurationStatePlatformType_AndroidForWork),
		string(DeviceConfigurationStatePlatformType_AndroidWorkProfile),
		string(DeviceConfigurationStatePlatformType_IOS),
		string(DeviceConfigurationStatePlatformType_MacOS),
		string(DeviceConfigurationStatePlatformType_Windows10AndLater),
		string(DeviceConfigurationStatePlatformType_Windows10XProfile),
		string(DeviceConfigurationStatePlatformType_Windows81AndLater),
		string(DeviceConfigurationStatePlatformType_WindowsPhone81),
	}
}

func (s *DeviceConfigurationStatePlatformType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceConfigurationStatePlatformType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceConfigurationStatePlatformType(input string) (*DeviceConfigurationStatePlatformType, error) {
	vals := map[string]DeviceConfigurationStatePlatformType{
		"all":                DeviceConfigurationStatePlatformType_All,
		"android":            DeviceConfigurationStatePlatformType_Android,
		"androidaosp":        DeviceConfigurationStatePlatformType_AndroidAOSP,
		"androidforwork":     DeviceConfigurationStatePlatformType_AndroidForWork,
		"androidworkprofile": DeviceConfigurationStatePlatformType_AndroidWorkProfile,
		"ios":                DeviceConfigurationStatePlatformType_IOS,
		"macos":              DeviceConfigurationStatePlatformType_MacOS,
		"windows10andlater":  DeviceConfigurationStatePlatformType_Windows10AndLater,
		"windows10xprofile":  DeviceConfigurationStatePlatformType_Windows10XProfile,
		"windows81andlater":  DeviceConfigurationStatePlatformType_Windows81AndLater,
		"windowsphone81":     DeviceConfigurationStatePlatformType_WindowsPhone81,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceConfigurationStatePlatformType(input)
	return &out, nil
}
