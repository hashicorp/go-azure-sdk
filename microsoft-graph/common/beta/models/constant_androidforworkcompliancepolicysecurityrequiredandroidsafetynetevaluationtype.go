package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkCompliancePolicySecurityRequiredAndroidSafetyNetEvaluationType string

const (
	AndroidForWorkCompliancePolicySecurityRequiredAndroidSafetyNetEvaluationTypebasic          AndroidForWorkCompliancePolicySecurityRequiredAndroidSafetyNetEvaluationType = "Basic"
	AndroidForWorkCompliancePolicySecurityRequiredAndroidSafetyNetEvaluationTypehardwareBacked AndroidForWorkCompliancePolicySecurityRequiredAndroidSafetyNetEvaluationType = "HardwareBacked"
)

func PossibleValuesForAndroidForWorkCompliancePolicySecurityRequiredAndroidSafetyNetEvaluationType() []string {
	return []string{
		string(AndroidForWorkCompliancePolicySecurityRequiredAndroidSafetyNetEvaluationTypebasic),
		string(AndroidForWorkCompliancePolicySecurityRequiredAndroidSafetyNetEvaluationTypehardwareBacked),
	}
}

func parseAndroidForWorkCompliancePolicySecurityRequiredAndroidSafetyNetEvaluationType(input string) (*AndroidForWorkCompliancePolicySecurityRequiredAndroidSafetyNetEvaluationType, error) {
	vals := map[string]AndroidForWorkCompliancePolicySecurityRequiredAndroidSafetyNetEvaluationType{
		"basic":          AndroidForWorkCompliancePolicySecurityRequiredAndroidSafetyNetEvaluationTypebasic,
		"hardwarebacked": AndroidForWorkCompliancePolicySecurityRequiredAndroidSafetyNetEvaluationTypehardwareBacked,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkCompliancePolicySecurityRequiredAndroidSafetyNetEvaluationType(input)
	return &out, nil
}
