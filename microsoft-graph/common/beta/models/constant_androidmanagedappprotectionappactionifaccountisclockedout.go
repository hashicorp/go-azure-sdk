package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppProtectionAppActionIfAccountIsClockedOut string

const (
	AndroidManagedAppProtectionAppActionIfAccountIsClockedOutblock AndroidManagedAppProtectionAppActionIfAccountIsClockedOut = "Block"
	AndroidManagedAppProtectionAppActionIfAccountIsClockedOutwarn  AndroidManagedAppProtectionAppActionIfAccountIsClockedOut = "Warn"
	AndroidManagedAppProtectionAppActionIfAccountIsClockedOutwipe  AndroidManagedAppProtectionAppActionIfAccountIsClockedOut = "Wipe"
)

func PossibleValuesForAndroidManagedAppProtectionAppActionIfAccountIsClockedOut() []string {
	return []string{
		string(AndroidManagedAppProtectionAppActionIfAccountIsClockedOutblock),
		string(AndroidManagedAppProtectionAppActionIfAccountIsClockedOutwarn),
		string(AndroidManagedAppProtectionAppActionIfAccountIsClockedOutwipe),
	}
}

func parseAndroidManagedAppProtectionAppActionIfAccountIsClockedOut(input string) (*AndroidManagedAppProtectionAppActionIfAccountIsClockedOut, error) {
	vals := map[string]AndroidManagedAppProtectionAppActionIfAccountIsClockedOut{
		"block": AndroidManagedAppProtectionAppActionIfAccountIsClockedOutblock,
		"warn":  AndroidManagedAppProtectionAppActionIfAccountIsClockedOutwarn,
		"wipe":  AndroidManagedAppProtectionAppActionIfAccountIsClockedOutwipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppProtectionAppActionIfAccountIsClockedOut(input)
	return &out, nil
}
