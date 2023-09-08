package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeManagedHomeScreenPinComplexity string

const (
	AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeManagedHomeScreenPinComplexitycomplex       AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeManagedHomeScreenPinComplexity = "Complex"
	AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeManagedHomeScreenPinComplexitynotConfigured AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeManagedHomeScreenPinComplexity = "NotConfigured"
	AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeManagedHomeScreenPinComplexitysimple        AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeManagedHomeScreenPinComplexity = "Simple"
)

func PossibleValuesForAndroidDeviceOwnerGeneralDeviceConfigurationKioskModeManagedHomeScreenPinComplexity() []string {
	return []string{
		string(AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeManagedHomeScreenPinComplexitycomplex),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeManagedHomeScreenPinComplexitynotConfigured),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeManagedHomeScreenPinComplexitysimple),
	}
}

func parseAndroidDeviceOwnerGeneralDeviceConfigurationKioskModeManagedHomeScreenPinComplexity(input string) (*AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeManagedHomeScreenPinComplexity, error) {
	vals := map[string]AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeManagedHomeScreenPinComplexity{
		"complex":       AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeManagedHomeScreenPinComplexitycomplex,
		"notconfigured": AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeManagedHomeScreenPinComplexitynotConfigured,
		"simple":        AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeManagedHomeScreenPinComplexitysimple,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeManagedHomeScreenPinComplexity(input)
	return &out, nil
}
