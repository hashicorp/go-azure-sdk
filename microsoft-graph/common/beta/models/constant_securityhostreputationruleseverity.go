package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityHostReputationRuleSeverity string

const (
	SecurityHostReputationRuleSeverityhigh    SecurityHostReputationRuleSeverity = "High"
	SecurityHostReputationRuleSeveritylow     SecurityHostReputationRuleSeverity = "Low"
	SecurityHostReputationRuleSeveritymedium  SecurityHostReputationRuleSeverity = "Medium"
	SecurityHostReputationRuleSeverityunknown SecurityHostReputationRuleSeverity = "Unknown"
)

func PossibleValuesForSecurityHostReputationRuleSeverity() []string {
	return []string{
		string(SecurityHostReputationRuleSeverityhigh),
		string(SecurityHostReputationRuleSeveritylow),
		string(SecurityHostReputationRuleSeveritymedium),
		string(SecurityHostReputationRuleSeverityunknown),
	}
}

func parseSecurityHostReputationRuleSeverity(input string) (*SecurityHostReputationRuleSeverity, error) {
	vals := map[string]SecurityHostReputationRuleSeverity{
		"high":    SecurityHostReputationRuleSeverityhigh,
		"low":     SecurityHostReputationRuleSeveritylow,
		"medium":  SecurityHostReputationRuleSeveritymedium,
		"unknown": SecurityHostReputationRuleSeverityunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityHostReputationRuleSeverity(input)
	return &out, nil
}
