package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHigh string

const (
	AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHighblock AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHigh = "Block"
	AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHighwarn  AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHigh = "Warn"
	AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHighwipe  AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHigh = "Wipe"
)

func PossibleValuesForAndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHigh() []string {
	return []string{
		string(AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHighblock),
		string(AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHighwarn),
		string(AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHighwipe),
	}
}

func parseAndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHigh(input string) (*AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHigh, error) {
	vals := map[string]AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHigh{
		"block": AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHighblock,
		"warn":  AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHighwarn,
		"wipe":  AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHighwipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHigh(input)
	return &out, nil
}
