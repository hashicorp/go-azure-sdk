package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowed string

const (
	AndroidManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowedblock AndroidManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowed = "Block"
	AndroidManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowedwarn  AndroidManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowed = "Warn"
	AndroidManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowedwipe  AndroidManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowed = "Wipe"
)

func PossibleValuesForAndroidManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowed() []string {
	return []string{
		string(AndroidManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowedblock),
		string(AndroidManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowedwarn),
		string(AndroidManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowedwipe),
	}
}

func parseAndroidManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowed(input string) (*AndroidManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowed, error) {
	vals := map[string]AndroidManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowed{
		"block": AndroidManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowedblock,
		"warn":  AndroidManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowedwarn,
		"wipe":  AndroidManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowedwipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowed(input)
	return &out, nil
}
