package webapplicationfirewallmanagedrulesets

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedRuleDefinition struct {
	DefaultAction *ActionType              `json:"defaultAction,omitempty"`
	DefaultState  *ManagedRuleEnabledState `json:"defaultState,omitempty"`
	Description   *string                  `json:"description,omitempty"`
	RuleId        *string                  `json:"ruleId,omitempty"`
}
