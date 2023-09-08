package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidCompliancePolicyRequiredPasswordComplexity string

const (
	AndroidCompliancePolicyRequiredPasswordComplexityhigh   AndroidCompliancePolicyRequiredPasswordComplexity = "High"
	AndroidCompliancePolicyRequiredPasswordComplexitylow    AndroidCompliancePolicyRequiredPasswordComplexity = "Low"
	AndroidCompliancePolicyRequiredPasswordComplexitymedium AndroidCompliancePolicyRequiredPasswordComplexity = "Medium"
	AndroidCompliancePolicyRequiredPasswordComplexitynone   AndroidCompliancePolicyRequiredPasswordComplexity = "None"
)

func PossibleValuesForAndroidCompliancePolicyRequiredPasswordComplexity() []string {
	return []string{
		string(AndroidCompliancePolicyRequiredPasswordComplexityhigh),
		string(AndroidCompliancePolicyRequiredPasswordComplexitylow),
		string(AndroidCompliancePolicyRequiredPasswordComplexitymedium),
		string(AndroidCompliancePolicyRequiredPasswordComplexitynone),
	}
}

func parseAndroidCompliancePolicyRequiredPasswordComplexity(input string) (*AndroidCompliancePolicyRequiredPasswordComplexity, error) {
	vals := map[string]AndroidCompliancePolicyRequiredPasswordComplexity{
		"high":   AndroidCompliancePolicyRequiredPasswordComplexityhigh,
		"low":    AndroidCompliancePolicyRequiredPasswordComplexitylow,
		"medium": AndroidCompliancePolicyRequiredPasswordComplexitymedium,
		"none":   AndroidCompliancePolicyRequiredPasswordComplexitynone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidCompliancePolicyRequiredPasswordComplexity(input)
	return &out, nil
}
