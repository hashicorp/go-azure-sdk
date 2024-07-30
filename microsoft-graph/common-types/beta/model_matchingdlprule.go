package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MatchingDlpRule struct {
	Actions           *[]DlpActionInfo         `json:"actions,omitempty"`
	IsMostRestrictive *bool                    `json:"isMostRestrictive,omitempty"`
	ODataType         *string                  `json:"@odata.type,omitempty"`
	PolicyId          *string                  `json:"policyId,omitempty"`
	PolicyName        *string                  `json:"policyName,omitempty"`
	Priority          *int64                   `json:"priority,omitempty"`
	RuleId            *string                  `json:"ruleId,omitempty"`
	RuleMode          *MatchingDlpRuleRuleMode `json:"ruleMode,omitempty"`
	RuleName          *string                  `json:"ruleName,omitempty"`
}
