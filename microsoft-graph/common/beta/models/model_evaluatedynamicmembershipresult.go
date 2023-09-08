package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EvaluateDynamicMembershipResult struct {
	MembershipRule                  *string                      `json:"membershipRule,omitempty"`
	MembershipRuleEvaluationDetails *ExpressionEvaluationDetails `json:"membershipRuleEvaluationDetails,omitempty"`
	MembershipRuleEvaluationResult  *bool                        `json:"membershipRuleEvaluationResult,omitempty"`
	ODataType                       *string                      `json:"@odata.type,omitempty"`
}
