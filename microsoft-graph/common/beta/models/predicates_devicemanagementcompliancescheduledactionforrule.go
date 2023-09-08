package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementComplianceScheduledActionForRuleOperationPredicate struct {
	Id        *string
	ODataType *string
	RuleName  *string
}

func (p DeviceManagementComplianceScheduledActionForRuleOperationPredicate) Matches(input DeviceManagementComplianceScheduledActionForRule) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RuleName != nil && (input.RuleName == nil || *p.RuleName != *input.RuleName) {
		return false
	}

	return true
}
