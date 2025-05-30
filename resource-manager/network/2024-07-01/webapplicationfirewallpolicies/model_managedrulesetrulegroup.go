package webapplicationfirewallpolicies

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedRuleSetRuleGroup struct {
	RuleGroupName string    `json:"ruleGroupName"`
	Rules         *[]string `json:"rules,omitempty"`
}
