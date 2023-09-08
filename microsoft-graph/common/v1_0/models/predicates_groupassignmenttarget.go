package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupAssignmentTargetOperationPredicate struct {
	GroupId   *string
	ODataType *string
}

func (p GroupAssignmentTargetOperationPredicate) Matches(input GroupAssignmentTarget) bool {

	if p.GroupId != nil && (input.GroupId == nil || *p.GroupId != *input.GroupId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
