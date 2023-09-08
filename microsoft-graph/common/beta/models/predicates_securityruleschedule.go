package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityRuleScheduleOperationPredicate struct {
	NextRunDateTime *string
	ODataType       *string
	Period          *string
}

func (p SecurityRuleScheduleOperationPredicate) Matches(input SecurityRuleSchedule) bool {

	if p.NextRunDateTime != nil && (input.NextRunDateTime == nil || *p.NextRunDateTime != *input.NextRunDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Period != nil && (input.Period == nil || *p.Period != *input.Period) {
		return false
	}

	return true
}
