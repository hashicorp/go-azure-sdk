package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMedium string

const (
	AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMediumblock AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMedium = "Block"
	AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMediumwarn  AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMedium = "Warn"
	AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMediumwipe  AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMedium = "Wipe"
)

func PossibleValuesForAndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMedium() []string {
	return []string{
		string(AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMediumblock),
		string(AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMediumwarn),
		string(AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMediumwipe),
	}
}

func parseAndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMedium(input string) (*AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMedium, error) {
	vals := map[string]AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMedium{
		"block": AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMediumblock,
		"warn":  AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMediumwarn,
		"wipe":  AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMediumwipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMedium(input)
	return &out, nil
}
