package openapis

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedRuleGroupDefinition struct {
	Description   *string                  `json:"description,omitempty"`
	RuleGroupName *string                  `json:"ruleGroupName,omitempty"`
	Rules         *[]ManagedRuleDefinition `json:"rules,omitempty"`
}
