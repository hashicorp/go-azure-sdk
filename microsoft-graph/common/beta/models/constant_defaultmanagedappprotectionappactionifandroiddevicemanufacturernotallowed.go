package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowed string

const (
	DefaultManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowedblock DefaultManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowed = "Block"
	DefaultManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowedwarn  DefaultManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowed = "Warn"
	DefaultManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowedwipe  DefaultManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowed = "Wipe"
)

func PossibleValuesForDefaultManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowed() []string {
	return []string{
		string(DefaultManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowedblock),
		string(DefaultManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowedwarn),
		string(DefaultManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowedwipe),
	}
}

func parseDefaultManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowed(input string) (*DefaultManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowed, error) {
	vals := map[string]DefaultManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowed{
		"block": DefaultManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowedblock,
		"warn":  DefaultManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowedwarn,
		"wipe":  DefaultManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowedwipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowed(input)
	return &out, nil
}
