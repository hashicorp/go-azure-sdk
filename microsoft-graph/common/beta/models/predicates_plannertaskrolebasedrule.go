package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerTaskRoleBasedRuleOperationPredicate struct {
	DefaultRule *string
	ODataType   *string
}

func (p PlannerTaskRoleBasedRuleOperationPredicate) Matches(input PlannerTaskRoleBasedRule) bool {

	if p.DefaultRule != nil && (input.DefaultRule == nil || *p.DefaultRule != *input.DefaultRule) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
