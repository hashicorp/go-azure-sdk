package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMedium string

const (
	DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMediumblock DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMedium = "Block"
	DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMediumwarn  DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMedium = "Warn"
	DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMediumwipe  DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMedium = "Wipe"
)

func PossibleValuesForDefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMedium() []string {
	return []string{
		string(DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMediumblock),
		string(DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMediumwarn),
		string(DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMediumwipe),
	}
}

func parseDefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMedium(input string) (*DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMedium, error) {
	vals := map[string]DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMedium{
		"block": DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMediumblock,
		"warn":  DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMediumwarn,
		"wipe":  DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMediumwipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMedium(input)
	return &out, nil
}
