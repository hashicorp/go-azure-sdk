package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppProtectionRequiredAndroidSafetyNetAppsVerificationType string

const (
	AndroidManagedAppProtectionRequiredAndroidSafetyNetAppsVerificationTypeenabled AndroidManagedAppProtectionRequiredAndroidSafetyNetAppsVerificationType = "Enabled"
	AndroidManagedAppProtectionRequiredAndroidSafetyNetAppsVerificationTypenone    AndroidManagedAppProtectionRequiredAndroidSafetyNetAppsVerificationType = "None"
)

func PossibleValuesForAndroidManagedAppProtectionRequiredAndroidSafetyNetAppsVerificationType() []string {
	return []string{
		string(AndroidManagedAppProtectionRequiredAndroidSafetyNetAppsVerificationTypeenabled),
		string(AndroidManagedAppProtectionRequiredAndroidSafetyNetAppsVerificationTypenone),
	}
}

func parseAndroidManagedAppProtectionRequiredAndroidSafetyNetAppsVerificationType(input string) (*AndroidManagedAppProtectionRequiredAndroidSafetyNetAppsVerificationType, error) {
	vals := map[string]AndroidManagedAppProtectionRequiredAndroidSafetyNetAppsVerificationType{
		"enabled": AndroidManagedAppProtectionRequiredAndroidSafetyNetAppsVerificationTypeenabled,
		"none":    AndroidManagedAppProtectionRequiredAndroidSafetyNetAppsVerificationTypenone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppProtectionRequiredAndroidSafetyNetAppsVerificationType(input)
	return &out, nil
}
