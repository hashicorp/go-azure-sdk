package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileCompliancePolicyRequiredPasswordComplexity string

const (
	AndroidWorkProfileCompliancePolicyRequiredPasswordComplexityhigh   AndroidWorkProfileCompliancePolicyRequiredPasswordComplexity = "High"
	AndroidWorkProfileCompliancePolicyRequiredPasswordComplexitylow    AndroidWorkProfileCompliancePolicyRequiredPasswordComplexity = "Low"
	AndroidWorkProfileCompliancePolicyRequiredPasswordComplexitymedium AndroidWorkProfileCompliancePolicyRequiredPasswordComplexity = "Medium"
	AndroidWorkProfileCompliancePolicyRequiredPasswordComplexitynone   AndroidWorkProfileCompliancePolicyRequiredPasswordComplexity = "None"
)

func PossibleValuesForAndroidWorkProfileCompliancePolicyRequiredPasswordComplexity() []string {
	return []string{
		string(AndroidWorkProfileCompliancePolicyRequiredPasswordComplexityhigh),
		string(AndroidWorkProfileCompliancePolicyRequiredPasswordComplexitylow),
		string(AndroidWorkProfileCompliancePolicyRequiredPasswordComplexitymedium),
		string(AndroidWorkProfileCompliancePolicyRequiredPasswordComplexitynone),
	}
}

func parseAndroidWorkProfileCompliancePolicyRequiredPasswordComplexity(input string) (*AndroidWorkProfileCompliancePolicyRequiredPasswordComplexity, error) {
	vals := map[string]AndroidWorkProfileCompliancePolicyRequiredPasswordComplexity{
		"high":   AndroidWorkProfileCompliancePolicyRequiredPasswordComplexityhigh,
		"low":    AndroidWorkProfileCompliancePolicyRequiredPasswordComplexitylow,
		"medium": AndroidWorkProfileCompliancePolicyRequiredPasswordComplexitymedium,
		"none":   AndroidWorkProfileCompliancePolicyRequiredPasswordComplexitynone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileCompliancePolicyRequiredPasswordComplexity(input)
	return &out, nil
}
