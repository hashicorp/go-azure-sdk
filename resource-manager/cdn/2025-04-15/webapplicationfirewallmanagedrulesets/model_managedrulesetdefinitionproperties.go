package webapplicationfirewallmanagedrulesets

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedRuleSetDefinitionProperties struct {
	ProvisioningState *string                       `json:"provisioningState,omitempty"`
	RuleGroups        *[]ManagedRuleGroupDefinition `json:"ruleGroups,omitempty"`
	RuleSetType       *string                       `json:"ruleSetType,omitempty"`
	RuleSetVersion    *string                       `json:"ruleSetVersion,omitempty"`
}
