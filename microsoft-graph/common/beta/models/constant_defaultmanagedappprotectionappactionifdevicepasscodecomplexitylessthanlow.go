package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanLow string

const (
	DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanLowblock DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanLow = "Block"
	DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanLowwarn  DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanLow = "Warn"
	DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanLowwipe  DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanLow = "Wipe"
)

func PossibleValuesForDefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanLow() []string {
	return []string{
		string(DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanLowblock),
		string(DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanLowwarn),
		string(DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanLowwipe),
	}
}

func parseDefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanLow(input string) (*DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanLow, error) {
	vals := map[string]DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanLow{
		"block": DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanLowblock,
		"warn":  DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanLowwarn,
		"wipe":  DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanLowwipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanLow(input)
	return &out, nil
}
