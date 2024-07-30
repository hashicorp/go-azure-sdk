package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DetectedAppPlatform string

const (
	DetectedAppPlatform_AndroidDedicatedAndFullyManaged DetectedAppPlatform = "androidDedicatedAndFullyManaged"
	DetectedAppPlatform_AndroidDeviceAdministrator      DetectedAppPlatform = "androidDeviceAdministrator"
	DetectedAppPlatform_AndroidOSP                      DetectedAppPlatform = "androidOSP"
	DetectedAppPlatform_AndroidWorkProfile              DetectedAppPlatform = "androidWorkProfile"
	DetectedAppPlatform_ChromeOS                        DetectedAppPlatform = "chromeOS"
	DetectedAppPlatform_Ios                             DetectedAppPlatform = "ios"
	DetectedAppPlatform_MacOS                           DetectedAppPlatform = "macOS"
	DetectedAppPlatform_Unknown                         DetectedAppPlatform = "unknown"
	DetectedAppPlatform_Windows                         DetectedAppPlatform = "windows"
	DetectedAppPlatform_WindowsHolographic              DetectedAppPlatform = "windowsHolographic"
	DetectedAppPlatform_WindowsMobile                   DetectedAppPlatform = "windowsMobile"
)

func PossibleValuesForDetectedAppPlatform() []string {
	return []string{
		string(DetectedAppPlatform_AndroidDedicatedAndFullyManaged),
		string(DetectedAppPlatform_AndroidDeviceAdministrator),
		string(DetectedAppPlatform_AndroidOSP),
		string(DetectedAppPlatform_AndroidWorkProfile),
		string(DetectedAppPlatform_ChromeOS),
		string(DetectedAppPlatform_Ios),
		string(DetectedAppPlatform_MacOS),
		string(DetectedAppPlatform_Unknown),
		string(DetectedAppPlatform_Windows),
		string(DetectedAppPlatform_WindowsHolographic),
		string(DetectedAppPlatform_WindowsMobile),
	}
}

func (s *DetectedAppPlatform) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDetectedAppPlatform(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDetectedAppPlatform(input string) (*DetectedAppPlatform, error) {
	vals := map[string]DetectedAppPlatform{
		"androiddedicatedandfullymanaged": DetectedAppPlatform_AndroidDedicatedAndFullyManaged,
		"androiddeviceadministrator":      DetectedAppPlatform_AndroidDeviceAdministrator,
		"androidosp":                      DetectedAppPlatform_AndroidOSP,
		"androidworkprofile":              DetectedAppPlatform_AndroidWorkProfile,
		"chromeos":                        DetectedAppPlatform_ChromeOS,
		"ios":                             DetectedAppPlatform_Ios,
		"macos":                           DetectedAppPlatform_MacOS,
		"unknown":                         DetectedAppPlatform_Unknown,
		"windows":                         DetectedAppPlatform_Windows,
		"windowsholographic":              DetectedAppPlatform_WindowsHolographic,
		"windowsmobile":                   DetectedAppPlatform_WindowsMobile,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DetectedAppPlatform(input)
	return &out, nil
}
