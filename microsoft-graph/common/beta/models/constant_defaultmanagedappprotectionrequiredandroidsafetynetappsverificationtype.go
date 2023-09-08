package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultManagedAppProtectionRequiredAndroidSafetyNetAppsVerificationType string

const (
	DefaultManagedAppProtectionRequiredAndroidSafetyNetAppsVerificationTypeenabled DefaultManagedAppProtectionRequiredAndroidSafetyNetAppsVerificationType = "Enabled"
	DefaultManagedAppProtectionRequiredAndroidSafetyNetAppsVerificationTypenone    DefaultManagedAppProtectionRequiredAndroidSafetyNetAppsVerificationType = "None"
)

func PossibleValuesForDefaultManagedAppProtectionRequiredAndroidSafetyNetAppsVerificationType() []string {
	return []string{
		string(DefaultManagedAppProtectionRequiredAndroidSafetyNetAppsVerificationTypeenabled),
		string(DefaultManagedAppProtectionRequiredAndroidSafetyNetAppsVerificationTypenone),
	}
}

func parseDefaultManagedAppProtectionRequiredAndroidSafetyNetAppsVerificationType(input string) (*DefaultManagedAppProtectionRequiredAndroidSafetyNetAppsVerificationType, error) {
	vals := map[string]DefaultManagedAppProtectionRequiredAndroidSafetyNetAppsVerificationType{
		"enabled": DefaultManagedAppProtectionRequiredAndroidSafetyNetAppsVerificationTypeenabled,
		"none":    DefaultManagedAppProtectionRequiredAndroidSafetyNetAppsVerificationTypenone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultManagedAppProtectionRequiredAndroidSafetyNetAppsVerificationType(input)
	return &out, nil
}
