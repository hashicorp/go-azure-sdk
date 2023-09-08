package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerAssignedToTaskBoardTaskFormatOperationPredicate struct {
	Id                  *string
	ODataType           *string
	UnassignedOrderHint *string
}

func (p PlannerAssignedToTaskBoardTaskFormatOperationPredicate) Matches(input PlannerAssignedToTaskBoardTaskFormat) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.UnassignedOrderHint != nil && (input.UnassignedOrderHint == nil || *p.UnassignedOrderHint != *input.UnassignedOrderHint) {
		return false
	}

	return true
}
