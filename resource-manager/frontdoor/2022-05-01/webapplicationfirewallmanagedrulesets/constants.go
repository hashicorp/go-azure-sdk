package webapplicationfirewallmanagedrulesets

import (
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActionType string

const (
	ActionTypeAllow          ActionType = "Allow"
	ActionTypeAnomalyScoring ActionType = "AnomalyScoring"
	ActionTypeBlock          ActionType = "Block"
	ActionTypeLog            ActionType = "Log"
	ActionTypeRedirect       ActionType = "Redirect"
)

func PossibleValuesForActionType() []string {
	return []string{
		string(ActionTypeAllow),
		string(ActionTypeAnomalyScoring),
		string(ActionTypeBlock),
		string(ActionTypeLog),
		string(ActionTypeRedirect),
	}
}

func parseActionType(input string) (*ActionType, error) {
	vals := map[string]ActionType{
		"allow":          ActionTypeAllow,
		"anomalyscoring": ActionTypeAnomalyScoring,
		"block":          ActionTypeBlock,
		"log":            ActionTypeLog,
		"redirect":       ActionTypeRedirect,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ActionType(input)
	return &out, nil
}

type ManagedRuleEnabledState string

const (
	ManagedRuleEnabledStateDisabled ManagedRuleEnabledState = "Disabled"
	ManagedRuleEnabledStateEnabled  ManagedRuleEnabledState = "Enabled"
)

func PossibleValuesForManagedRuleEnabledState() []string {
	return []string{
		string(ManagedRuleEnabledStateDisabled),
		string(ManagedRuleEnabledStateEnabled),
	}
}

func parseManagedRuleEnabledState(input string) (*ManagedRuleEnabledState, error) {
	vals := map[string]ManagedRuleEnabledState{
		"disabled": ManagedRuleEnabledStateDisabled,
		"enabled":  ManagedRuleEnabledStateEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedRuleEnabledState(input)
	return &out, nil
}
