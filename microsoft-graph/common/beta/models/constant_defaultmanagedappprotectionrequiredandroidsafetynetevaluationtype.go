package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultManagedAppProtectionRequiredAndroidSafetyNetEvaluationType string

const (
	DefaultManagedAppProtectionRequiredAndroidSafetyNetEvaluationTypebasic          DefaultManagedAppProtectionRequiredAndroidSafetyNetEvaluationType = "Basic"
	DefaultManagedAppProtectionRequiredAndroidSafetyNetEvaluationTypehardwareBacked DefaultManagedAppProtectionRequiredAndroidSafetyNetEvaluationType = "HardwareBacked"
)

func PossibleValuesForDefaultManagedAppProtectionRequiredAndroidSafetyNetEvaluationType() []string {
	return []string{
		string(DefaultManagedAppProtectionRequiredAndroidSafetyNetEvaluationTypebasic),
		string(DefaultManagedAppProtectionRequiredAndroidSafetyNetEvaluationTypehardwareBacked),
	}
}

func parseDefaultManagedAppProtectionRequiredAndroidSafetyNetEvaluationType(input string) (*DefaultManagedAppProtectionRequiredAndroidSafetyNetEvaluationType, error) {
	vals := map[string]DefaultManagedAppProtectionRequiredAndroidSafetyNetEvaluationType{
		"basic":          DefaultManagedAppProtectionRequiredAndroidSafetyNetEvaluationTypebasic,
		"hardwarebacked": DefaultManagedAppProtectionRequiredAndroidSafetyNetEvaluationTypehardwareBacked,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultManagedAppProtectionRequiredAndroidSafetyNetEvaluationType(input)
	return &out, nil
}
