package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailed string

const (
	DefaultManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailedblock DefaultManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailed = "Block"
	DefaultManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailedwarn  DefaultManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailed = "Warn"
	DefaultManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailedwipe  DefaultManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailed = "Wipe"
)

func PossibleValuesForDefaultManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailed() []string {
	return []string{
		string(DefaultManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailedblock),
		string(DefaultManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailedwarn),
		string(DefaultManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailedwipe),
	}
}

func parseDefaultManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailed(input string) (*DefaultManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailed, error) {
	vals := map[string]DefaultManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailed{
		"block": DefaultManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailedblock,
		"warn":  DefaultManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailedwarn,
		"wipe":  DefaultManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailedwipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailed(input)
	return &out, nil
}
