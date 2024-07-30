package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceCompliancePolicyStatePlatformType string

const (
	DeviceCompliancePolicyStatePlatformType_All               DeviceCompliancePolicyStatePlatformType = "all"
	DeviceCompliancePolicyStatePlatformType_Android           DeviceCompliancePolicyStatePlatformType = "android"
	DeviceCompliancePolicyStatePlatformType_AndroidForWork    DeviceCompliancePolicyStatePlatformType = "androidForWork"
	DeviceCompliancePolicyStatePlatformType_IOS               DeviceCompliancePolicyStatePlatformType = "iOS"
	DeviceCompliancePolicyStatePlatformType_MacOS             DeviceCompliancePolicyStatePlatformType = "macOS"
	DeviceCompliancePolicyStatePlatformType_Windows10AndLater DeviceCompliancePolicyStatePlatformType = "windows10AndLater"
	DeviceCompliancePolicyStatePlatformType_Windows81AndLater DeviceCompliancePolicyStatePlatformType = "windows81AndLater"
	DeviceCompliancePolicyStatePlatformType_WindowsPhone81    DeviceCompliancePolicyStatePlatformType = "windowsPhone81"
)

func PossibleValuesForDeviceCompliancePolicyStatePlatformType() []string {
	return []string{
		string(DeviceCompliancePolicyStatePlatformType_All),
		string(DeviceCompliancePolicyStatePlatformType_Android),
		string(DeviceCompliancePolicyStatePlatformType_AndroidForWork),
		string(DeviceCompliancePolicyStatePlatformType_IOS),
		string(DeviceCompliancePolicyStatePlatformType_MacOS),
		string(DeviceCompliancePolicyStatePlatformType_Windows10AndLater),
		string(DeviceCompliancePolicyStatePlatformType_Windows81AndLater),
		string(DeviceCompliancePolicyStatePlatformType_WindowsPhone81),
	}
}

func (s *DeviceCompliancePolicyStatePlatformType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceCompliancePolicyStatePlatformType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceCompliancePolicyStatePlatformType(input string) (*DeviceCompliancePolicyStatePlatformType, error) {
	vals := map[string]DeviceCompliancePolicyStatePlatformType{
		"all":               DeviceCompliancePolicyStatePlatformType_All,
		"android":           DeviceCompliancePolicyStatePlatformType_Android,
		"androidforwork":    DeviceCompliancePolicyStatePlatformType_AndroidForWork,
		"ios":               DeviceCompliancePolicyStatePlatformType_IOS,
		"macos":             DeviceCompliancePolicyStatePlatformType_MacOS,
		"windows10andlater": DeviceCompliancePolicyStatePlatformType_Windows10AndLater,
		"windows81andlater": DeviceCompliancePolicyStatePlatformType_Windows81AndLater,
		"windowsphone81":    DeviceCompliancePolicyStatePlatformType_WindowsPhone81,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceCompliancePolicyStatePlatformType(input)
	return &out, nil
}
