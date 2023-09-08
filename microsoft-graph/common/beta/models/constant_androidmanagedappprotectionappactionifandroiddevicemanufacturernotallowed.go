package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowed string

const (
	AndroidManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowedblock AndroidManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowed = "Block"
	AndroidManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowedwarn  AndroidManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowed = "Warn"
	AndroidManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowedwipe  AndroidManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowed = "Wipe"
)

func PossibleValuesForAndroidManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowed() []string {
	return []string{
		string(AndroidManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowedblock),
		string(AndroidManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowedwarn),
		string(AndroidManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowedwipe),
	}
}

func parseAndroidManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowed(input string) (*AndroidManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowed, error) {
	vals := map[string]AndroidManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowed{
		"block": AndroidManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowedblock,
		"warn":  AndroidManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowedwarn,
		"wipe":  AndroidManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowedwipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowed(input)
	return &out, nil
}
