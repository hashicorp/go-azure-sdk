package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultManagedAppProtectionAppActionIfAccountIsClockedOut string

const (
	DefaultManagedAppProtectionAppActionIfAccountIsClockedOutblock DefaultManagedAppProtectionAppActionIfAccountIsClockedOut = "Block"
	DefaultManagedAppProtectionAppActionIfAccountIsClockedOutwarn  DefaultManagedAppProtectionAppActionIfAccountIsClockedOut = "Warn"
	DefaultManagedAppProtectionAppActionIfAccountIsClockedOutwipe  DefaultManagedAppProtectionAppActionIfAccountIsClockedOut = "Wipe"
)

func PossibleValuesForDefaultManagedAppProtectionAppActionIfAccountIsClockedOut() []string {
	return []string{
		string(DefaultManagedAppProtectionAppActionIfAccountIsClockedOutblock),
		string(DefaultManagedAppProtectionAppActionIfAccountIsClockedOutwarn),
		string(DefaultManagedAppProtectionAppActionIfAccountIsClockedOutwipe),
	}
}

func parseDefaultManagedAppProtectionAppActionIfAccountIsClockedOut(input string) (*DefaultManagedAppProtectionAppActionIfAccountIsClockedOut, error) {
	vals := map[string]DefaultManagedAppProtectionAppActionIfAccountIsClockedOut{
		"block": DefaultManagedAppProtectionAppActionIfAccountIsClockedOutblock,
		"warn":  DefaultManagedAppProtectionAppActionIfAccountIsClockedOutwarn,
		"wipe":  DefaultManagedAppProtectionAppActionIfAccountIsClockedOutwipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultManagedAppProtectionAppActionIfAccountIsClockedOut(input)
	return &out, nil
}
