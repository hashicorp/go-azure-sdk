package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHigh string

const (
	DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHighblock DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHigh = "Block"
	DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHighwarn  DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHigh = "Warn"
	DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHighwipe  DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHigh = "Wipe"
)

func PossibleValuesForDefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHigh() []string {
	return []string{
		string(DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHighblock),
		string(DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHighwarn),
		string(DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHighwipe),
	}
}

func parseDefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHigh(input string) (*DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHigh, error) {
	vals := map[string]DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHigh{
		"block": DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHighblock,
		"warn":  DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHighwarn,
		"wipe":  DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHighwipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHigh(input)
	return &out, nil
}
