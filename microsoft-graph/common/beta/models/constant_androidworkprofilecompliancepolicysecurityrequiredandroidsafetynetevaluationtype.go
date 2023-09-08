package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileCompliancePolicySecurityRequiredAndroidSafetyNetEvaluationType string

const (
	AndroidWorkProfileCompliancePolicySecurityRequiredAndroidSafetyNetEvaluationTypebasic          AndroidWorkProfileCompliancePolicySecurityRequiredAndroidSafetyNetEvaluationType = "Basic"
	AndroidWorkProfileCompliancePolicySecurityRequiredAndroidSafetyNetEvaluationTypehardwareBacked AndroidWorkProfileCompliancePolicySecurityRequiredAndroidSafetyNetEvaluationType = "HardwareBacked"
)

func PossibleValuesForAndroidWorkProfileCompliancePolicySecurityRequiredAndroidSafetyNetEvaluationType() []string {
	return []string{
		string(AndroidWorkProfileCompliancePolicySecurityRequiredAndroidSafetyNetEvaluationTypebasic),
		string(AndroidWorkProfileCompliancePolicySecurityRequiredAndroidSafetyNetEvaluationTypehardwareBacked),
	}
}

func parseAndroidWorkProfileCompliancePolicySecurityRequiredAndroidSafetyNetEvaluationType(input string) (*AndroidWorkProfileCompliancePolicySecurityRequiredAndroidSafetyNetEvaluationType, error) {
	vals := map[string]AndroidWorkProfileCompliancePolicySecurityRequiredAndroidSafetyNetEvaluationType{
		"basic":          AndroidWorkProfileCompliancePolicySecurityRequiredAndroidSafetyNetEvaluationTypebasic,
		"hardwarebacked": AndroidWorkProfileCompliancePolicySecurityRequiredAndroidSafetyNetEvaluationTypehardwareBacked,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileCompliancePolicySecurityRequiredAndroidSafetyNetEvaluationType(input)
	return &out, nil
}
