package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailed string

const (
	DefaultManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailedblock DefaultManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailed = "Block"
	DefaultManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailedwarn  DefaultManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailed = "Warn"
	DefaultManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailedwipe  DefaultManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailed = "Wipe"
)

func PossibleValuesForDefaultManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailed() []string {
	return []string{
		string(DefaultManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailedblock),
		string(DefaultManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailedwarn),
		string(DefaultManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailedwipe),
	}
}

func parseDefaultManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailed(input string) (*DefaultManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailed, error) {
	vals := map[string]DefaultManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailed{
		"block": DefaultManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailedblock,
		"warn":  DefaultManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailedwarn,
		"wipe":  DefaultManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailedwipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailed(input)
	return &out, nil
}
