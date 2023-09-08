package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceEnrollmentPlatformRestrictionConfigurationPlatformType string

const (
	DeviceEnrollmentPlatformRestrictionConfigurationPlatformTypeallPlatforms   DeviceEnrollmentPlatformRestrictionConfigurationPlatformType = "AllPlatforms"
	DeviceEnrollmentPlatformRestrictionConfigurationPlatformTypeandroid        DeviceEnrollmentPlatformRestrictionConfigurationPlatformType = "Android"
	DeviceEnrollmentPlatformRestrictionConfigurationPlatformTypeandroidForWork DeviceEnrollmentPlatformRestrictionConfigurationPlatformType = "AndroidForWork"
	DeviceEnrollmentPlatformRestrictionConfigurationPlatformTypeios            DeviceEnrollmentPlatformRestrictionConfigurationPlatformType = "Ios"
	DeviceEnrollmentPlatformRestrictionConfigurationPlatformTypelinux          DeviceEnrollmentPlatformRestrictionConfigurationPlatformType = "Linux"
	DeviceEnrollmentPlatformRestrictionConfigurationPlatformTypemac            DeviceEnrollmentPlatformRestrictionConfigurationPlatformType = "Mac"
	DeviceEnrollmentPlatformRestrictionConfigurationPlatformTypewindows        DeviceEnrollmentPlatformRestrictionConfigurationPlatformType = "Windows"
	DeviceEnrollmentPlatformRestrictionConfigurationPlatformTypewindowsPhone   DeviceEnrollmentPlatformRestrictionConfigurationPlatformType = "WindowsPhone"
)

func PossibleValuesForDeviceEnrollmentPlatformRestrictionConfigurationPlatformType() []string {
	return []string{
		string(DeviceEnrollmentPlatformRestrictionConfigurationPlatformTypeallPlatforms),
		string(DeviceEnrollmentPlatformRestrictionConfigurationPlatformTypeandroid),
		string(DeviceEnrollmentPlatformRestrictionConfigurationPlatformTypeandroidForWork),
		string(DeviceEnrollmentPlatformRestrictionConfigurationPlatformTypeios),
		string(DeviceEnrollmentPlatformRestrictionConfigurationPlatformTypelinux),
		string(DeviceEnrollmentPlatformRestrictionConfigurationPlatformTypemac),
		string(DeviceEnrollmentPlatformRestrictionConfigurationPlatformTypewindows),
		string(DeviceEnrollmentPlatformRestrictionConfigurationPlatformTypewindowsPhone),
	}
}

func parseDeviceEnrollmentPlatformRestrictionConfigurationPlatformType(input string) (*DeviceEnrollmentPlatformRestrictionConfigurationPlatformType, error) {
	vals := map[string]DeviceEnrollmentPlatformRestrictionConfigurationPlatformType{
		"allplatforms":   DeviceEnrollmentPlatformRestrictionConfigurationPlatformTypeallPlatforms,
		"android":        DeviceEnrollmentPlatformRestrictionConfigurationPlatformTypeandroid,
		"androidforwork": DeviceEnrollmentPlatformRestrictionConfigurationPlatformTypeandroidForWork,
		"ios":            DeviceEnrollmentPlatformRestrictionConfigurationPlatformTypeios,
		"linux":          DeviceEnrollmentPlatformRestrictionConfigurationPlatformTypelinux,
		"mac":            DeviceEnrollmentPlatformRestrictionConfigurationPlatformTypemac,
		"windows":        DeviceEnrollmentPlatformRestrictionConfigurationPlatformTypewindows,
		"windowsphone":   DeviceEnrollmentPlatformRestrictionConfigurationPlatformTypewindowsPhone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceEnrollmentPlatformRestrictionConfigurationPlatformType(input)
	return &out, nil
}
