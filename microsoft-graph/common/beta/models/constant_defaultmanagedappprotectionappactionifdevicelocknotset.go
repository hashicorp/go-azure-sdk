package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultManagedAppProtectionAppActionIfDeviceLockNotSet string

const (
	DefaultManagedAppProtectionAppActionIfDeviceLockNotSetblock DefaultManagedAppProtectionAppActionIfDeviceLockNotSet = "Block"
	DefaultManagedAppProtectionAppActionIfDeviceLockNotSetwarn  DefaultManagedAppProtectionAppActionIfDeviceLockNotSet = "Warn"
	DefaultManagedAppProtectionAppActionIfDeviceLockNotSetwipe  DefaultManagedAppProtectionAppActionIfDeviceLockNotSet = "Wipe"
)

func PossibleValuesForDefaultManagedAppProtectionAppActionIfDeviceLockNotSet() []string {
	return []string{
		string(DefaultManagedAppProtectionAppActionIfDeviceLockNotSetblock),
		string(DefaultManagedAppProtectionAppActionIfDeviceLockNotSetwarn),
		string(DefaultManagedAppProtectionAppActionIfDeviceLockNotSetwipe),
	}
}

func parseDefaultManagedAppProtectionAppActionIfDeviceLockNotSet(input string) (*DefaultManagedAppProtectionAppActionIfDeviceLockNotSet, error) {
	vals := map[string]DefaultManagedAppProtectionAppActionIfDeviceLockNotSet{
		"block": DefaultManagedAppProtectionAppActionIfDeviceLockNotSetblock,
		"warn":  DefaultManagedAppProtectionAppActionIfDeviceLockNotSetwarn,
		"wipe":  DefaultManagedAppProtectionAppActionIfDeviceLockNotSetwipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultManagedAppProtectionAppActionIfDeviceLockNotSet(input)
	return &out, nil
}
