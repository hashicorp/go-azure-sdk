package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppProtectionRequiredAndroidSafetyNetEvaluationType string

const (
	AndroidManagedAppProtectionRequiredAndroidSafetyNetEvaluationTypebasic          AndroidManagedAppProtectionRequiredAndroidSafetyNetEvaluationType = "Basic"
	AndroidManagedAppProtectionRequiredAndroidSafetyNetEvaluationTypehardwareBacked AndroidManagedAppProtectionRequiredAndroidSafetyNetEvaluationType = "HardwareBacked"
)

func PossibleValuesForAndroidManagedAppProtectionRequiredAndroidSafetyNetEvaluationType() []string {
	return []string{
		string(AndroidManagedAppProtectionRequiredAndroidSafetyNetEvaluationTypebasic),
		string(AndroidManagedAppProtectionRequiredAndroidSafetyNetEvaluationTypehardwareBacked),
	}
}

func parseAndroidManagedAppProtectionRequiredAndroidSafetyNetEvaluationType(input string) (*AndroidManagedAppProtectionRequiredAndroidSafetyNetEvaluationType, error) {
	vals := map[string]AndroidManagedAppProtectionRequiredAndroidSafetyNetEvaluationType{
		"basic":          AndroidManagedAppProtectionRequiredAndroidSafetyNetEvaluationTypebasic,
		"hardwarebacked": AndroidManagedAppProtectionRequiredAndroidSafetyNetEvaluationTypehardwareBacked,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppProtectionRequiredAndroidSafetyNetEvaluationType(input)
	return &out, nil
}
