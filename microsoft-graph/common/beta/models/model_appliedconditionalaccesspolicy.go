package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppliedConditionalAccessPolicy struct {
	AuthenticationStrength      *AuthenticationStrength                                 `json:"authenticationStrength,omitempty"`
	ConditionsNotSatisfied      *[]AppliedConditionalAccessPolicyConditionsNotSatisfied `json:"conditionsNotSatisfied,omitempty"`
	ConditionsSatisfied         *[]AppliedConditionalAccessPolicyConditionsSatisfied    `json:"conditionsSatisfied,omitempty"`
	DisplayName                 *string                                                 `json:"displayName,omitempty"`
	EnforcedGrantControls       *[]string                                               `json:"enforcedGrantControls,omitempty"`
	EnforcedSessionControls     *[]string                                               `json:"enforcedSessionControls,omitempty"`
	ExcludeRulesSatisfied       *[]ConditionalAccessRuleSatisfied                       `json:"excludeRulesSatisfied,omitempty"`
	Id                          *string                                                 `json:"id,omitempty"`
	IncludeRulesSatisfied       *[]ConditionalAccessRuleSatisfied                       `json:"includeRulesSatisfied,omitempty"`
	ODataType                   *string                                                 `json:"@odata.type,omitempty"`
	Result                      *AppliedConditionalAccessPolicyResult                   `json:"result,omitempty"`
	SessionControlsNotSatisfied *[]string                                               `json:"sessionControlsNotSatisfied,omitempty"`
}
