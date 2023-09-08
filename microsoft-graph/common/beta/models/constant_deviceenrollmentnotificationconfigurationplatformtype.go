package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceEnrollmentNotificationConfigurationPlatformType string

const (
	DeviceEnrollmentNotificationConfigurationPlatformTypeallPlatforms   DeviceEnrollmentNotificationConfigurationPlatformType = "AllPlatforms"
	DeviceEnrollmentNotificationConfigurationPlatformTypeandroid        DeviceEnrollmentNotificationConfigurationPlatformType = "Android"
	DeviceEnrollmentNotificationConfigurationPlatformTypeandroidForWork DeviceEnrollmentNotificationConfigurationPlatformType = "AndroidForWork"
	DeviceEnrollmentNotificationConfigurationPlatformTypeios            DeviceEnrollmentNotificationConfigurationPlatformType = "Ios"
	DeviceEnrollmentNotificationConfigurationPlatformTypelinux          DeviceEnrollmentNotificationConfigurationPlatformType = "Linux"
	DeviceEnrollmentNotificationConfigurationPlatformTypemac            DeviceEnrollmentNotificationConfigurationPlatformType = "Mac"
	DeviceEnrollmentNotificationConfigurationPlatformTypewindows        DeviceEnrollmentNotificationConfigurationPlatformType = "Windows"
	DeviceEnrollmentNotificationConfigurationPlatformTypewindowsPhone   DeviceEnrollmentNotificationConfigurationPlatformType = "WindowsPhone"
)

func PossibleValuesForDeviceEnrollmentNotificationConfigurationPlatformType() []string {
	return []string{
		string(DeviceEnrollmentNotificationConfigurationPlatformTypeallPlatforms),
		string(DeviceEnrollmentNotificationConfigurationPlatformTypeandroid),
		string(DeviceEnrollmentNotificationConfigurationPlatformTypeandroidForWork),
		string(DeviceEnrollmentNotificationConfigurationPlatformTypeios),
		string(DeviceEnrollmentNotificationConfigurationPlatformTypelinux),
		string(DeviceEnrollmentNotificationConfigurationPlatformTypemac),
		string(DeviceEnrollmentNotificationConfigurationPlatformTypewindows),
		string(DeviceEnrollmentNotificationConfigurationPlatformTypewindowsPhone),
	}
}

func parseDeviceEnrollmentNotificationConfigurationPlatformType(input string) (*DeviceEnrollmentNotificationConfigurationPlatformType, error) {
	vals := map[string]DeviceEnrollmentNotificationConfigurationPlatformType{
		"allplatforms":   DeviceEnrollmentNotificationConfigurationPlatformTypeallPlatforms,
		"android":        DeviceEnrollmentNotificationConfigurationPlatformTypeandroid,
		"androidforwork": DeviceEnrollmentNotificationConfigurationPlatformTypeandroidForWork,
		"ios":            DeviceEnrollmentNotificationConfigurationPlatformTypeios,
		"linux":          DeviceEnrollmentNotificationConfigurationPlatformTypelinux,
		"mac":            DeviceEnrollmentNotificationConfigurationPlatformTypemac,
		"windows":        DeviceEnrollmentNotificationConfigurationPlatformTypewindows,
		"windowsphone":   DeviceEnrollmentNotificationConfigurationPlatformTypewindowsPhone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceEnrollmentNotificationConfigurationPlatformType(input)
	return &out, nil
}
