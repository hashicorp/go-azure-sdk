package webapplicationfirewallpolicies

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedRulesDefinition struct {
	Exceptions      *[]ExceptionEntry         `json:"exceptions,omitempty"`
	Exclusions      *[]OwaspCrsExclusionEntry `json:"exclusions,omitempty"`
	ManagedRuleSets []ManagedRuleSet          `json:"managedRuleSets"`
}
