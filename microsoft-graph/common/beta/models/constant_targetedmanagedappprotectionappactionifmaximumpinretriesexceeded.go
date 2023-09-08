package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TargetedManagedAppProtectionAppActionIfMaximumPinRetriesExceeded string

const (
	TargetedManagedAppProtectionAppActionIfMaximumPinRetriesExceededblock TargetedManagedAppProtectionAppActionIfMaximumPinRetriesExceeded = "Block"
	TargetedManagedAppProtectionAppActionIfMaximumPinRetriesExceededwarn  TargetedManagedAppProtectionAppActionIfMaximumPinRetriesExceeded = "Warn"
	TargetedManagedAppProtectionAppActionIfMaximumPinRetriesExceededwipe  TargetedManagedAppProtectionAppActionIfMaximumPinRetriesExceeded = "Wipe"
)

func PossibleValuesForTargetedManagedAppProtectionAppActionIfMaximumPinRetriesExceeded() []string {
	return []string{
		string(TargetedManagedAppProtectionAppActionIfMaximumPinRetriesExceededblock),
		string(TargetedManagedAppProtectionAppActionIfMaximumPinRetriesExceededwarn),
		string(TargetedManagedAppProtectionAppActionIfMaximumPinRetriesExceededwipe),
	}
}

func parseTargetedManagedAppProtectionAppActionIfMaximumPinRetriesExceeded(input string) (*TargetedManagedAppProtectionAppActionIfMaximumPinRetriesExceeded, error) {
	vals := map[string]TargetedManagedAppProtectionAppActionIfMaximumPinRetriesExceeded{
		"block": TargetedManagedAppProtectionAppActionIfMaximumPinRetriesExceededblock,
		"warn":  TargetedManagedAppProtectionAppActionIfMaximumPinRetriesExceededwarn,
		"wipe":  TargetedManagedAppProtectionAppActionIfMaximumPinRetriesExceededwipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TargetedManagedAppProtectionAppActionIfMaximumPinRetriesExceeded(input)
	return &out, nil
}
