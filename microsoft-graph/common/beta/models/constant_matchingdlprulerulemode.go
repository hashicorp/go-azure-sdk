package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MatchingDlpRuleRuleMode string

const (
	MatchingDlpRuleRuleModeaudit           MatchingDlpRuleRuleMode = "Audit"
	MatchingDlpRuleRuleModeauditAndNotify  MatchingDlpRuleRuleMode = "AuditAndNotify"
	MatchingDlpRuleRuleModeenforce         MatchingDlpRuleRuleMode = "Enforce"
	MatchingDlpRuleRuleModependingDeletion MatchingDlpRuleRuleMode = "PendingDeletion"
	MatchingDlpRuleRuleModetest            MatchingDlpRuleRuleMode = "Test"
)

func PossibleValuesForMatchingDlpRuleRuleMode() []string {
	return []string{
		string(MatchingDlpRuleRuleModeaudit),
		string(MatchingDlpRuleRuleModeauditAndNotify),
		string(MatchingDlpRuleRuleModeenforce),
		string(MatchingDlpRuleRuleModependingDeletion),
		string(MatchingDlpRuleRuleModetest),
	}
}

func parseMatchingDlpRuleRuleMode(input string) (*MatchingDlpRuleRuleMode, error) {
	vals := map[string]MatchingDlpRuleRuleMode{
		"audit":           MatchingDlpRuleRuleModeaudit,
		"auditandnotify":  MatchingDlpRuleRuleModeauditAndNotify,
		"enforce":         MatchingDlpRuleRuleModeenforce,
		"pendingdeletion": MatchingDlpRuleRuleModependingDeletion,
		"test":            MatchingDlpRuleRuleModetest,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MatchingDlpRuleRuleMode(input)
	return &out, nil
}
