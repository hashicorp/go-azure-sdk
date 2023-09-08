package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultManagedAppProtectionAppActionIfMaximumPinRetriesExceeded string

const (
	DefaultManagedAppProtectionAppActionIfMaximumPinRetriesExceededblock DefaultManagedAppProtectionAppActionIfMaximumPinRetriesExceeded = "Block"
	DefaultManagedAppProtectionAppActionIfMaximumPinRetriesExceededwarn  DefaultManagedAppProtectionAppActionIfMaximumPinRetriesExceeded = "Warn"
	DefaultManagedAppProtectionAppActionIfMaximumPinRetriesExceededwipe  DefaultManagedAppProtectionAppActionIfMaximumPinRetriesExceeded = "Wipe"
)

func PossibleValuesForDefaultManagedAppProtectionAppActionIfMaximumPinRetriesExceeded() []string {
	return []string{
		string(DefaultManagedAppProtectionAppActionIfMaximumPinRetriesExceededblock),
		string(DefaultManagedAppProtectionAppActionIfMaximumPinRetriesExceededwarn),
		string(DefaultManagedAppProtectionAppActionIfMaximumPinRetriesExceededwipe),
	}
}

func parseDefaultManagedAppProtectionAppActionIfMaximumPinRetriesExceeded(input string) (*DefaultManagedAppProtectionAppActionIfMaximumPinRetriesExceeded, error) {
	vals := map[string]DefaultManagedAppProtectionAppActionIfMaximumPinRetriesExceeded{
		"block": DefaultManagedAppProtectionAppActionIfMaximumPinRetriesExceededblock,
		"warn":  DefaultManagedAppProtectionAppActionIfMaximumPinRetriesExceededwarn,
		"wipe":  DefaultManagedAppProtectionAppActionIfMaximumPinRetriesExceededwipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultManagedAppProtectionAppActionIfMaximumPinRetriesExceeded(input)
	return &out, nil
}
