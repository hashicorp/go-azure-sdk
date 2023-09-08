package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppProtectionAppActionIfDeviceLockNotSet string

const (
	AndroidManagedAppProtectionAppActionIfDeviceLockNotSetblock AndroidManagedAppProtectionAppActionIfDeviceLockNotSet = "Block"
	AndroidManagedAppProtectionAppActionIfDeviceLockNotSetwarn  AndroidManagedAppProtectionAppActionIfDeviceLockNotSet = "Warn"
	AndroidManagedAppProtectionAppActionIfDeviceLockNotSetwipe  AndroidManagedAppProtectionAppActionIfDeviceLockNotSet = "Wipe"
)

func PossibleValuesForAndroidManagedAppProtectionAppActionIfDeviceLockNotSet() []string {
	return []string{
		string(AndroidManagedAppProtectionAppActionIfDeviceLockNotSetblock),
		string(AndroidManagedAppProtectionAppActionIfDeviceLockNotSetwarn),
		string(AndroidManagedAppProtectionAppActionIfDeviceLockNotSetwipe),
	}
}

func parseAndroidManagedAppProtectionAppActionIfDeviceLockNotSet(input string) (*AndroidManagedAppProtectionAppActionIfDeviceLockNotSet, error) {
	vals := map[string]AndroidManagedAppProtectionAppActionIfDeviceLockNotSet{
		"block": AndroidManagedAppProtectionAppActionIfDeviceLockNotSetblock,
		"warn":  AndroidManagedAppProtectionAppActionIfDeviceLockNotSetwarn,
		"wipe":  AndroidManagedAppProtectionAppActionIfDeviceLockNotSetwipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppProtectionAppActionIfDeviceLockNotSet(input)
	return &out, nil
}
