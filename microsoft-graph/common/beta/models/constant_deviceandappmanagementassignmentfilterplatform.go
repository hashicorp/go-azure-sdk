package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceAndAppManagementAssignmentFilterPlatform string

const (
	DeviceAndAppManagementAssignmentFilterPlatformandroid                            DeviceAndAppManagementAssignmentFilterPlatform = "Android"
	DeviceAndAppManagementAssignmentFilterPlatformandroidAOSP                        DeviceAndAppManagementAssignmentFilterPlatform = "AndroidAOSP"
	DeviceAndAppManagementAssignmentFilterPlatformandroidForWork                     DeviceAndAppManagementAssignmentFilterPlatform = "AndroidForWork"
	DeviceAndAppManagementAssignmentFilterPlatformandroidMobileApplicationManagement DeviceAndAppManagementAssignmentFilterPlatform = "AndroidMobileApplicationManagement"
	DeviceAndAppManagementAssignmentFilterPlatformandroidWorkProfile                 DeviceAndAppManagementAssignmentFilterPlatform = "AndroidWorkProfile"
	DeviceAndAppManagementAssignmentFilterPlatformiOS                                DeviceAndAppManagementAssignmentFilterPlatform = "IOS"
	DeviceAndAppManagementAssignmentFilterPlatformiOSMobileApplicationManagement     DeviceAndAppManagementAssignmentFilterPlatform = "IOSMobileApplicationManagement"
	DeviceAndAppManagementAssignmentFilterPlatformmacOS                              DeviceAndAppManagementAssignmentFilterPlatform = "MacOS"
	DeviceAndAppManagementAssignmentFilterPlatformunknown                            DeviceAndAppManagementAssignmentFilterPlatform = "Unknown"
	DeviceAndAppManagementAssignmentFilterPlatformwindows10AndLater                  DeviceAndAppManagementAssignmentFilterPlatform = "Windows10AndLater"
	DeviceAndAppManagementAssignmentFilterPlatformwindows81AndLater                  DeviceAndAppManagementAssignmentFilterPlatform = "Windows81AndLater"
	DeviceAndAppManagementAssignmentFilterPlatformwindowsPhone81                     DeviceAndAppManagementAssignmentFilterPlatform = "WindowsPhone81"
)

func PossibleValuesForDeviceAndAppManagementAssignmentFilterPlatform() []string {
	return []string{
		string(DeviceAndAppManagementAssignmentFilterPlatformandroid),
		string(DeviceAndAppManagementAssignmentFilterPlatformandroidAOSP),
		string(DeviceAndAppManagementAssignmentFilterPlatformandroidForWork),
		string(DeviceAndAppManagementAssignmentFilterPlatformandroidMobileApplicationManagement),
		string(DeviceAndAppManagementAssignmentFilterPlatformandroidWorkProfile),
		string(DeviceAndAppManagementAssignmentFilterPlatformiOS),
		string(DeviceAndAppManagementAssignmentFilterPlatformiOSMobileApplicationManagement),
		string(DeviceAndAppManagementAssignmentFilterPlatformmacOS),
		string(DeviceAndAppManagementAssignmentFilterPlatformunknown),
		string(DeviceAndAppManagementAssignmentFilterPlatformwindows10AndLater),
		string(DeviceAndAppManagementAssignmentFilterPlatformwindows81AndLater),
		string(DeviceAndAppManagementAssignmentFilterPlatformwindowsPhone81),
	}
}

func parseDeviceAndAppManagementAssignmentFilterPlatform(input string) (*DeviceAndAppManagementAssignmentFilterPlatform, error) {
	vals := map[string]DeviceAndAppManagementAssignmentFilterPlatform{
		"android":                            DeviceAndAppManagementAssignmentFilterPlatformandroid,
		"androidaosp":                        DeviceAndAppManagementAssignmentFilterPlatformandroidAOSP,
		"androidforwork":                     DeviceAndAppManagementAssignmentFilterPlatformandroidForWork,
		"androidmobileapplicationmanagement": DeviceAndAppManagementAssignmentFilterPlatformandroidMobileApplicationManagement,
		"androidworkprofile":                 DeviceAndAppManagementAssignmentFilterPlatformandroidWorkProfile,
		"ios":                                DeviceAndAppManagementAssignmentFilterPlatformiOS,
		"iosmobileapplicationmanagement":     DeviceAndAppManagementAssignmentFilterPlatformiOSMobileApplicationManagement,
		"macos":                              DeviceAndAppManagementAssignmentFilterPlatformmacOS,
		"unknown":                            DeviceAndAppManagementAssignmentFilterPlatformunknown,
		"windows10andlater":                  DeviceAndAppManagementAssignmentFilterPlatformwindows10AndLater,
		"windows81andlater":                  DeviceAndAppManagementAssignmentFilterPlatformwindows81AndLater,
		"windowsphone81":                     DeviceAndAppManagementAssignmentFilterPlatformwindowsPhone81,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceAndAppManagementAssignmentFilterPlatform(input)
	return &out, nil
}
