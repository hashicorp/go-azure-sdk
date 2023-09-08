package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailed string

const (
	AndroidManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailedblock AndroidManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailed = "Block"
	AndroidManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailedwarn  AndroidManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailed = "Warn"
	AndroidManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailedwipe  AndroidManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailed = "Wipe"
)

func PossibleValuesForAndroidManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailed() []string {
	return []string{
		string(AndroidManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailedblock),
		string(AndroidManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailedwarn),
		string(AndroidManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailedwipe),
	}
}

func parseAndroidManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailed(input string) (*AndroidManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailed, error) {
	vals := map[string]AndroidManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailed{
		"block": AndroidManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailedblock,
		"warn":  AndroidManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailedwarn,
		"wipe":  AndroidManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailedwipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailed(input)
	return &out, nil
}
