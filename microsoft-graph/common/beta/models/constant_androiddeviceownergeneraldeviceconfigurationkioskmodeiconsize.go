package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeIconSize string

const (
	AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeIconSizelarge         AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeIconSize = "Large"
	AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeIconSizelargest       AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeIconSize = "Largest"
	AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeIconSizenotConfigured AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeIconSize = "NotConfigured"
	AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeIconSizeregular       AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeIconSize = "Regular"
	AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeIconSizesmall         AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeIconSize = "Small"
	AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeIconSizesmallest      AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeIconSize = "Smallest"
)

func PossibleValuesForAndroidDeviceOwnerGeneralDeviceConfigurationKioskModeIconSize() []string {
	return []string{
		string(AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeIconSizelarge),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeIconSizelargest),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeIconSizenotConfigured),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeIconSizeregular),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeIconSizesmall),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeIconSizesmallest),
	}
}

func parseAndroidDeviceOwnerGeneralDeviceConfigurationKioskModeIconSize(input string) (*AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeIconSize, error) {
	vals := map[string]AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeIconSize{
		"large":         AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeIconSizelarge,
		"largest":       AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeIconSizelargest,
		"notconfigured": AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeIconSizenotConfigured,
		"regular":       AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeIconSizeregular,
		"small":         AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeIconSizesmall,
		"smallest":      AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeIconSizesmallest,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeIconSize(input)
	return &out, nil
}
