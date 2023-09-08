package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeScreenOrientation string

const (
	AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeScreenOrientationautoRotate    AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeScreenOrientation = "AutoRotate"
	AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeScreenOrientationlandscape     AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeScreenOrientation = "Landscape"
	AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeScreenOrientationnotConfigured AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeScreenOrientation = "NotConfigured"
	AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeScreenOrientationportrait      AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeScreenOrientation = "Portrait"
)

func PossibleValuesForAndroidDeviceOwnerGeneralDeviceConfigurationKioskModeScreenOrientation() []string {
	return []string{
		string(AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeScreenOrientationautoRotate),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeScreenOrientationlandscape),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeScreenOrientationnotConfigured),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeScreenOrientationportrait),
	}
}

func parseAndroidDeviceOwnerGeneralDeviceConfigurationKioskModeScreenOrientation(input string) (*AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeScreenOrientation, error) {
	vals := map[string]AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeScreenOrientation{
		"autorotate":    AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeScreenOrientationautoRotate,
		"landscape":     AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeScreenOrientationlandscape,
		"notconfigured": AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeScreenOrientationnotConfigured,
		"portrait":      AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeScreenOrientationportrait,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeScreenOrientation(input)
	return &out, nil
}
