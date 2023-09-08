package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkCompliancePolicyRequiredPasswordComplexity string

const (
	AndroidForWorkCompliancePolicyRequiredPasswordComplexityhigh   AndroidForWorkCompliancePolicyRequiredPasswordComplexity = "High"
	AndroidForWorkCompliancePolicyRequiredPasswordComplexitylow    AndroidForWorkCompliancePolicyRequiredPasswordComplexity = "Low"
	AndroidForWorkCompliancePolicyRequiredPasswordComplexitymedium AndroidForWorkCompliancePolicyRequiredPasswordComplexity = "Medium"
	AndroidForWorkCompliancePolicyRequiredPasswordComplexitynone   AndroidForWorkCompliancePolicyRequiredPasswordComplexity = "None"
)

func PossibleValuesForAndroidForWorkCompliancePolicyRequiredPasswordComplexity() []string {
	return []string{
		string(AndroidForWorkCompliancePolicyRequiredPasswordComplexityhigh),
		string(AndroidForWorkCompliancePolicyRequiredPasswordComplexitylow),
		string(AndroidForWorkCompliancePolicyRequiredPasswordComplexitymedium),
		string(AndroidForWorkCompliancePolicyRequiredPasswordComplexitynone),
	}
}

func parseAndroidForWorkCompliancePolicyRequiredPasswordComplexity(input string) (*AndroidForWorkCompliancePolicyRequiredPasswordComplexity, error) {
	vals := map[string]AndroidForWorkCompliancePolicyRequiredPasswordComplexity{
		"high":   AndroidForWorkCompliancePolicyRequiredPasswordComplexityhigh,
		"low":    AndroidForWorkCompliancePolicyRequiredPasswordComplexitylow,
		"medium": AndroidForWorkCompliancePolicyRequiredPasswordComplexitymedium,
		"none":   AndroidForWorkCompliancePolicyRequiredPasswordComplexitynone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkCompliancePolicyRequiredPasswordComplexity(input)
	return &out, nil
}
