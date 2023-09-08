package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerAssignmentOperationPredicate struct {
	AssignedDateTime *string
	ODataType        *string
	OrderHint        *string
}

func (p PlannerAssignmentOperationPredicate) Matches(input PlannerAssignment) bool {

	if p.AssignedDateTime != nil && (input.AssignedDateTime == nil || *p.AssignedDateTime != *input.AssignedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OrderHint != nil && (input.OrderHint == nil || *p.OrderHint != *input.OrderHint) {
		return false
	}

	return true
}
