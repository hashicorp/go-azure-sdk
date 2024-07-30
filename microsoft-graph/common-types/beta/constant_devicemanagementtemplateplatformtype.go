package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementTemplatePlatformType string

const (
	DeviceManagementTemplatePlatformType_All                DeviceManagementTemplatePlatformType = "all"
	DeviceManagementTemplatePlatformType_Android            DeviceManagementTemplatePlatformType = "android"
	DeviceManagementTemplatePlatformType_AndroidAOSP        DeviceManagementTemplatePlatformType = "androidAOSP"
	DeviceManagementTemplatePlatformType_AndroidForWork     DeviceManagementTemplatePlatformType = "androidForWork"
	DeviceManagementTemplatePlatformType_AndroidWorkProfile DeviceManagementTemplatePlatformType = "androidWorkProfile"
	DeviceManagementTemplatePlatformType_IOS                DeviceManagementTemplatePlatformType = "iOS"
	DeviceManagementTemplatePlatformType_MacOS              DeviceManagementTemplatePlatformType = "macOS"
	DeviceManagementTemplatePlatformType_Windows10AndLater  DeviceManagementTemplatePlatformType = "windows10AndLater"
	DeviceManagementTemplatePlatformType_Windows10XProfile  DeviceManagementTemplatePlatformType = "windows10XProfile"
	DeviceManagementTemplatePlatformType_Windows81AndLater  DeviceManagementTemplatePlatformType = "windows81AndLater"
	DeviceManagementTemplatePlatformType_WindowsPhone81     DeviceManagementTemplatePlatformType = "windowsPhone81"
)

func PossibleValuesForDeviceManagementTemplatePlatformType() []string {
	return []string{
		string(DeviceManagementTemplatePlatformType_All),
		string(DeviceManagementTemplatePlatformType_Android),
		string(DeviceManagementTemplatePlatformType_AndroidAOSP),
		string(DeviceManagementTemplatePlatformType_AndroidForWork),
		string(DeviceManagementTemplatePlatformType_AndroidWorkProfile),
		string(DeviceManagementTemplatePlatformType_IOS),
		string(DeviceManagementTemplatePlatformType_MacOS),
		string(DeviceManagementTemplatePlatformType_Windows10AndLater),
		string(DeviceManagementTemplatePlatformType_Windows10XProfile),
		string(DeviceManagementTemplatePlatformType_Windows81AndLater),
		string(DeviceManagementTemplatePlatformType_WindowsPhone81),
	}
}

func (s *DeviceManagementTemplatePlatformType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementTemplatePlatformType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementTemplatePlatformType(input string) (*DeviceManagementTemplatePlatformType, error) {
	vals := map[string]DeviceManagementTemplatePlatformType{
		"all":                DeviceManagementTemplatePlatformType_All,
		"android":            DeviceManagementTemplatePlatformType_Android,
		"androidaosp":        DeviceManagementTemplatePlatformType_AndroidAOSP,
		"androidforwork":     DeviceManagementTemplatePlatformType_AndroidForWork,
		"androidworkprofile": DeviceManagementTemplatePlatformType_AndroidWorkProfile,
		"ios":                DeviceManagementTemplatePlatformType_IOS,
		"macos":              DeviceManagementTemplatePlatformType_MacOS,
		"windows10andlater":  DeviceManagementTemplatePlatformType_Windows10AndLater,
		"windows10xprofile":  DeviceManagementTemplatePlatformType_Windows10XProfile,
		"windows81andlater":  DeviceManagementTemplatePlatformType_Windows81AndLater,
		"windowsphone81":     DeviceManagementTemplatePlatformType_WindowsPhone81,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementTemplatePlatformType(input)
	return &out, nil
}
