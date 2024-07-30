package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceAndAppManagementAssignmentFilterPlatform string

const (
	DeviceAndAppManagementAssignmentFilterPlatform_Android                            DeviceAndAppManagementAssignmentFilterPlatform = "android"
	DeviceAndAppManagementAssignmentFilterPlatform_AndroidAOSP                        DeviceAndAppManagementAssignmentFilterPlatform = "androidAOSP"
	DeviceAndAppManagementAssignmentFilterPlatform_AndroidForWork                     DeviceAndAppManagementAssignmentFilterPlatform = "androidForWork"
	DeviceAndAppManagementAssignmentFilterPlatform_AndroidMobileApplicationManagement DeviceAndAppManagementAssignmentFilterPlatform = "androidMobileApplicationManagement"
	DeviceAndAppManagementAssignmentFilterPlatform_AndroidWorkProfile                 DeviceAndAppManagementAssignmentFilterPlatform = "androidWorkProfile"
	DeviceAndAppManagementAssignmentFilterPlatform_IOS                                DeviceAndAppManagementAssignmentFilterPlatform = "iOS"
	DeviceAndAppManagementAssignmentFilterPlatform_IOSMobileApplicationManagement     DeviceAndAppManagementAssignmentFilterPlatform = "iOSMobileApplicationManagement"
	DeviceAndAppManagementAssignmentFilterPlatform_MacOS                              DeviceAndAppManagementAssignmentFilterPlatform = "macOS"
	DeviceAndAppManagementAssignmentFilterPlatform_Unknown                            DeviceAndAppManagementAssignmentFilterPlatform = "unknown"
	DeviceAndAppManagementAssignmentFilterPlatform_Windows10AndLater                  DeviceAndAppManagementAssignmentFilterPlatform = "windows10AndLater"
	DeviceAndAppManagementAssignmentFilterPlatform_Windows81AndLater                  DeviceAndAppManagementAssignmentFilterPlatform = "windows81AndLater"
	DeviceAndAppManagementAssignmentFilterPlatform_WindowsPhone81                     DeviceAndAppManagementAssignmentFilterPlatform = "windowsPhone81"
)

func PossibleValuesForDeviceAndAppManagementAssignmentFilterPlatform() []string {
	return []string{
		string(DeviceAndAppManagementAssignmentFilterPlatform_Android),
		string(DeviceAndAppManagementAssignmentFilterPlatform_AndroidAOSP),
		string(DeviceAndAppManagementAssignmentFilterPlatform_AndroidForWork),
		string(DeviceAndAppManagementAssignmentFilterPlatform_AndroidMobileApplicationManagement),
		string(DeviceAndAppManagementAssignmentFilterPlatform_AndroidWorkProfile),
		string(DeviceAndAppManagementAssignmentFilterPlatform_IOS),
		string(DeviceAndAppManagementAssignmentFilterPlatform_IOSMobileApplicationManagement),
		string(DeviceAndAppManagementAssignmentFilterPlatform_MacOS),
		string(DeviceAndAppManagementAssignmentFilterPlatform_Unknown),
		string(DeviceAndAppManagementAssignmentFilterPlatform_Windows10AndLater),
		string(DeviceAndAppManagementAssignmentFilterPlatform_Windows81AndLater),
		string(DeviceAndAppManagementAssignmentFilterPlatform_WindowsPhone81),
	}
}

func (s *DeviceAndAppManagementAssignmentFilterPlatform) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceAndAppManagementAssignmentFilterPlatform(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceAndAppManagementAssignmentFilterPlatform(input string) (*DeviceAndAppManagementAssignmentFilterPlatform, error) {
	vals := map[string]DeviceAndAppManagementAssignmentFilterPlatform{
		"android":                            DeviceAndAppManagementAssignmentFilterPlatform_Android,
		"androidaosp":                        DeviceAndAppManagementAssignmentFilterPlatform_AndroidAOSP,
		"androidforwork":                     DeviceAndAppManagementAssignmentFilterPlatform_AndroidForWork,
		"androidmobileapplicationmanagement": DeviceAndAppManagementAssignmentFilterPlatform_AndroidMobileApplicationManagement,
		"androidworkprofile":                 DeviceAndAppManagementAssignmentFilterPlatform_AndroidWorkProfile,
		"ios":                                DeviceAndAppManagementAssignmentFilterPlatform_IOS,
		"iosmobileapplicationmanagement":     DeviceAndAppManagementAssignmentFilterPlatform_IOSMobileApplicationManagement,
		"macos":                              DeviceAndAppManagementAssignmentFilterPlatform_MacOS,
		"unknown":                            DeviceAndAppManagementAssignmentFilterPlatform_Unknown,
		"windows10andlater":                  DeviceAndAppManagementAssignmentFilterPlatform_Windows10AndLater,
		"windows81andlater":                  DeviceAndAppManagementAssignmentFilterPlatform_Windows81AndLater,
		"windowsphone81":                     DeviceAndAppManagementAssignmentFilterPlatform_WindowsPhone81,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceAndAppManagementAssignmentFilterPlatform(input)
	return &out, nil
}
