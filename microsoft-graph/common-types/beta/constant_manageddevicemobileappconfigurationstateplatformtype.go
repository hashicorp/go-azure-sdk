package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceMobileAppConfigurationStatePlatformType string

const (
	ManagedDeviceMobileAppConfigurationStatePlatformType_All                ManagedDeviceMobileAppConfigurationStatePlatformType = "all"
	ManagedDeviceMobileAppConfigurationStatePlatformType_Android            ManagedDeviceMobileAppConfigurationStatePlatformType = "android"
	ManagedDeviceMobileAppConfigurationStatePlatformType_AndroidAOSP        ManagedDeviceMobileAppConfigurationStatePlatformType = "androidAOSP"
	ManagedDeviceMobileAppConfigurationStatePlatformType_AndroidForWork     ManagedDeviceMobileAppConfigurationStatePlatformType = "androidForWork"
	ManagedDeviceMobileAppConfigurationStatePlatformType_AndroidWorkProfile ManagedDeviceMobileAppConfigurationStatePlatformType = "androidWorkProfile"
	ManagedDeviceMobileAppConfigurationStatePlatformType_IOS                ManagedDeviceMobileAppConfigurationStatePlatformType = "iOS"
	ManagedDeviceMobileAppConfigurationStatePlatformType_MacOS              ManagedDeviceMobileAppConfigurationStatePlatformType = "macOS"
	ManagedDeviceMobileAppConfigurationStatePlatformType_Windows10AndLater  ManagedDeviceMobileAppConfigurationStatePlatformType = "windows10AndLater"
	ManagedDeviceMobileAppConfigurationStatePlatformType_Windows10XProfile  ManagedDeviceMobileAppConfigurationStatePlatformType = "windows10XProfile"
	ManagedDeviceMobileAppConfigurationStatePlatformType_Windows81AndLater  ManagedDeviceMobileAppConfigurationStatePlatformType = "windows81AndLater"
	ManagedDeviceMobileAppConfigurationStatePlatformType_WindowsPhone81     ManagedDeviceMobileAppConfigurationStatePlatformType = "windowsPhone81"
)

func PossibleValuesForManagedDeviceMobileAppConfigurationStatePlatformType() []string {
	return []string{
		string(ManagedDeviceMobileAppConfigurationStatePlatformType_All),
		string(ManagedDeviceMobileAppConfigurationStatePlatformType_Android),
		string(ManagedDeviceMobileAppConfigurationStatePlatformType_AndroidAOSP),
		string(ManagedDeviceMobileAppConfigurationStatePlatformType_AndroidForWork),
		string(ManagedDeviceMobileAppConfigurationStatePlatformType_AndroidWorkProfile),
		string(ManagedDeviceMobileAppConfigurationStatePlatformType_IOS),
		string(ManagedDeviceMobileAppConfigurationStatePlatformType_MacOS),
		string(ManagedDeviceMobileAppConfigurationStatePlatformType_Windows10AndLater),
		string(ManagedDeviceMobileAppConfigurationStatePlatformType_Windows10XProfile),
		string(ManagedDeviceMobileAppConfigurationStatePlatformType_Windows81AndLater),
		string(ManagedDeviceMobileAppConfigurationStatePlatformType_WindowsPhone81),
	}
}

func (s *ManagedDeviceMobileAppConfigurationStatePlatformType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedDeviceMobileAppConfigurationStatePlatformType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedDeviceMobileAppConfigurationStatePlatformType(input string) (*ManagedDeviceMobileAppConfigurationStatePlatformType, error) {
	vals := map[string]ManagedDeviceMobileAppConfigurationStatePlatformType{
		"all":                ManagedDeviceMobileAppConfigurationStatePlatformType_All,
		"android":            ManagedDeviceMobileAppConfigurationStatePlatformType_Android,
		"androidaosp":        ManagedDeviceMobileAppConfigurationStatePlatformType_AndroidAOSP,
		"androidforwork":     ManagedDeviceMobileAppConfigurationStatePlatformType_AndroidForWork,
		"androidworkprofile": ManagedDeviceMobileAppConfigurationStatePlatformType_AndroidWorkProfile,
		"ios":                ManagedDeviceMobileAppConfigurationStatePlatformType_IOS,
		"macos":              ManagedDeviceMobileAppConfigurationStatePlatformType_MacOS,
		"windows10andlater":  ManagedDeviceMobileAppConfigurationStatePlatformType_Windows10AndLater,
		"windows10xprofile":  ManagedDeviceMobileAppConfigurationStatePlatformType_Windows10XProfile,
		"windows81andlater":  ManagedDeviceMobileAppConfigurationStatePlatformType_Windows81AndLater,
		"windowsphone81":     ManagedDeviceMobileAppConfigurationStatePlatformType_WindowsPhone81,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceMobileAppConfigurationStatePlatformType(input)
	return &out, nil
}
