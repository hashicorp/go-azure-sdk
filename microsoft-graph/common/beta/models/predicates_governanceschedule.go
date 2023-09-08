package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GovernanceScheduleOperationPredicate struct {
	Duration      *string
	EndDateTime   *string
	ODataType     *string
	StartDateTime *string
	Type          *string
}

func (p GovernanceScheduleOperationPredicate) Matches(input GovernanceSchedule) bool {

	if p.Duration != nil && (input.Duration == nil || *p.Duration != *input.Duration) {
		return false
	}

	if p.EndDateTime != nil && (input.EndDateTime == nil || *p.EndDateTime != *input.EndDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.StartDateTime != nil && (input.StartDateTime == nil || *p.StartDateTime != *input.StartDateTime) {
		return false
	}

	if p.Type != nil && (input.Type == nil || *p.Type != *input.Type) {
		return false
	}

	return true
}
