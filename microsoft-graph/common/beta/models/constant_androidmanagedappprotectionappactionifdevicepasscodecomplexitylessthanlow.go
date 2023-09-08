package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanLow string

const (
	AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanLowblock AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanLow = "Block"
	AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanLowwarn  AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanLow = "Warn"
	AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanLowwipe  AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanLow = "Wipe"
)

func PossibleValuesForAndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanLow() []string {
	return []string{
		string(AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanLowblock),
		string(AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanLowwarn),
		string(AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanLowwipe),
	}
}

func parseAndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanLow(input string) (*AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanLow, error) {
	vals := map[string]AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanLow{
		"block": AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanLowblock,
		"warn":  AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanLowwarn,
		"wipe":  AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanLowwipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanLow(input)
	return &out, nil
}
