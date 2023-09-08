package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowed string

const (
	DefaultManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowedblock DefaultManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowed = "Block"
	DefaultManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowedwarn  DefaultManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowed = "Warn"
	DefaultManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowedwipe  DefaultManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowed = "Wipe"
)

func PossibleValuesForDefaultManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowed() []string {
	return []string{
		string(DefaultManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowedblock),
		string(DefaultManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowedwarn),
		string(DefaultManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowedwipe),
	}
}

func parseDefaultManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowed(input string) (*DefaultManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowed, error) {
	vals := map[string]DefaultManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowed{
		"block": DefaultManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowedblock,
		"warn":  DefaultManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowedwarn,
		"wipe":  DefaultManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowedwipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowed(input)
	return &out, nil
}
