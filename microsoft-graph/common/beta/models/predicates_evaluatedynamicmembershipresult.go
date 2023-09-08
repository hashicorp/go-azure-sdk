package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EvaluateDynamicMembershipResultOperationPredicate struct {
	MembershipRule                 *string
	MembershipRuleEvaluationResult *bool
	ODataType                      *string
}

func (p EvaluateDynamicMembershipResultOperationPredicate) Matches(input EvaluateDynamicMembershipResult) bool {

	if p.MembershipRule != nil && (input.MembershipRule == nil || *p.MembershipRule != *input.MembershipRule) {
		return false
	}

	if p.MembershipRuleEvaluationResult != nil && (input.MembershipRuleEvaluationResult == nil || *p.MembershipRuleEvaluationResult != *input.MembershipRuleEvaluationResult) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
