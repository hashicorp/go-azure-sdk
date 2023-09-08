package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeVirtualHomeButtonType string

const (
	AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeVirtualHomeButtonTypefloating      AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeVirtualHomeButtonType = "Floating"
	AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeVirtualHomeButtonTypenotConfigured AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeVirtualHomeButtonType = "NotConfigured"
	AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeVirtualHomeButtonTypeswipeUp       AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeVirtualHomeButtonType = "SwipeUp"
)

func PossibleValuesForAndroidDeviceOwnerGeneralDeviceConfigurationKioskModeVirtualHomeButtonType() []string {
	return []string{
		string(AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeVirtualHomeButtonTypefloating),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeVirtualHomeButtonTypenotConfigured),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeVirtualHomeButtonTypeswipeUp),
	}
}

func parseAndroidDeviceOwnerGeneralDeviceConfigurationKioskModeVirtualHomeButtonType(input string) (*AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeVirtualHomeButtonType, error) {
	vals := map[string]AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeVirtualHomeButtonType{
		"floating":      AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeVirtualHomeButtonTypefloating,
		"notconfigured": AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeVirtualHomeButtonTypenotConfigured,
		"swipeup":       AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeVirtualHomeButtonTypeswipeUp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeVirtualHomeButtonType(input)
	return &out, nil
}
