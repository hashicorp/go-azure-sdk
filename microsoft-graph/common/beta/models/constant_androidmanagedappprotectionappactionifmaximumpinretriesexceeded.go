package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppProtectionAppActionIfMaximumPinRetriesExceeded string

const (
	AndroidManagedAppProtectionAppActionIfMaximumPinRetriesExceededblock AndroidManagedAppProtectionAppActionIfMaximumPinRetriesExceeded = "Block"
	AndroidManagedAppProtectionAppActionIfMaximumPinRetriesExceededwarn  AndroidManagedAppProtectionAppActionIfMaximumPinRetriesExceeded = "Warn"
	AndroidManagedAppProtectionAppActionIfMaximumPinRetriesExceededwipe  AndroidManagedAppProtectionAppActionIfMaximumPinRetriesExceeded = "Wipe"
)

func PossibleValuesForAndroidManagedAppProtectionAppActionIfMaximumPinRetriesExceeded() []string {
	return []string{
		string(AndroidManagedAppProtectionAppActionIfMaximumPinRetriesExceededblock),
		string(AndroidManagedAppProtectionAppActionIfMaximumPinRetriesExceededwarn),
		string(AndroidManagedAppProtectionAppActionIfMaximumPinRetriesExceededwipe),
	}
}

func parseAndroidManagedAppProtectionAppActionIfMaximumPinRetriesExceeded(input string) (*AndroidManagedAppProtectionAppActionIfMaximumPinRetriesExceeded, error) {
	vals := map[string]AndroidManagedAppProtectionAppActionIfMaximumPinRetriesExceeded{
		"block": AndroidManagedAppProtectionAppActionIfMaximumPinRetriesExceededblock,
		"warn":  AndroidManagedAppProtectionAppActionIfMaximumPinRetriesExceededwarn,
		"wipe":  AndroidManagedAppProtectionAppActionIfMaximumPinRetriesExceededwipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppProtectionAppActionIfMaximumPinRetriesExceeded(input)
	return &out, nil
}
