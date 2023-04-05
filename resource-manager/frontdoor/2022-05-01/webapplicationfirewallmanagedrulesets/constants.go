package webapplicationfirewallmanagedrulesets

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
