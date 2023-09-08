package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailed string

const (
	AndroidManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailedblock AndroidManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailed = "Block"
	AndroidManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailedwarn  AndroidManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailed = "Warn"
	AndroidManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailedwipe  AndroidManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailed = "Wipe"
)

func PossibleValuesForAndroidManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailed() []string {
	return []string{
		string(AndroidManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailedblock),
		string(AndroidManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailedwarn),
		string(AndroidManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailedwipe),
	}
}

func parseAndroidManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailed(input string) (*AndroidManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailed, error) {
	vals := map[string]AndroidManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailed{
		"block": AndroidManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailedblock,
		"warn":  AndroidManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailedwarn,
		"wipe":  AndroidManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailedwipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailed(input)
	return &out, nil
}
