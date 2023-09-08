package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerBucketTaskBoardTaskFormatOperationPredicate struct {
	Id        *string
	ODataType *string
	OrderHint *string
}

func (p PlannerBucketTaskBoardTaskFormatOperationPredicate) Matches(input PlannerBucketTaskBoardTaskFormat) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
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
