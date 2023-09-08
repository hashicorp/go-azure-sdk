package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityDeviceEvidenceRiskScore string

const (
	SecurityDeviceEvidenceRiskScorehigh          SecurityDeviceEvidenceRiskScore = "High"
	SecurityDeviceEvidenceRiskScoreinformational SecurityDeviceEvidenceRiskScore = "Informational"
	SecurityDeviceEvidenceRiskScorelow           SecurityDeviceEvidenceRiskScore = "Low"
	SecurityDeviceEvidenceRiskScoremedium        SecurityDeviceEvidenceRiskScore = "Medium"
	SecurityDeviceEvidenceRiskScorenone          SecurityDeviceEvidenceRiskScore = "None"
)

func PossibleValuesForSecurityDeviceEvidenceRiskScore() []string {
	return []string{
		string(SecurityDeviceEvidenceRiskScorehigh),
		string(SecurityDeviceEvidenceRiskScoreinformational),
		string(SecurityDeviceEvidenceRiskScorelow),
		string(SecurityDeviceEvidenceRiskScoremedium),
		string(SecurityDeviceEvidenceRiskScorenone),
	}
}

func parseSecurityDeviceEvidenceRiskScore(input string) (*SecurityDeviceEvidenceRiskScore, error) {
	vals := map[string]SecurityDeviceEvidenceRiskScore{
		"high":          SecurityDeviceEvidenceRiskScorehigh,
		"informational": SecurityDeviceEvidenceRiskScoreinformational,
		"low":           SecurityDeviceEvidenceRiskScorelow,
		"medium":        SecurityDeviceEvidenceRiskScoremedium,
		"none":          SecurityDeviceEvidenceRiskScorenone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityDeviceEvidenceRiskScore(input)
	return &out, nil
}
