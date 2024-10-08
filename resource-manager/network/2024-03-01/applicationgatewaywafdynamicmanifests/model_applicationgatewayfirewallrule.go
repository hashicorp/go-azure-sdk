package applicationgatewaywafdynamicmanifests

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationGatewayFirewallRule struct {
	Action       *ApplicationGatewayWafRuleActionTypes      `json:"action,omitempty"`
	Description  *string                                    `json:"description,omitempty"`
	RuleId       int64                                      `json:"ruleId"`
	RuleIdString *string                                    `json:"ruleIdString,omitempty"`
	Sensitivity  *ApplicationGatewayWafRuleSensitivityTypes `json:"sensitivity,omitempty"`
	State        *ApplicationGatewayWafRuleStateTypes       `json:"state,omitempty"`
}
