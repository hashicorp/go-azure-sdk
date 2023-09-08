package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignedTrainingInfoOperationPredicate struct {
	AssignedUserCount  *int64
	CompletedUserCount *int64
	DisplayName        *string
	ODataType          *string
}

func (p AssignedTrainingInfoOperationPredicate) Matches(input AssignedTrainingInfo) bool {

	if p.AssignedUserCount != nil && (input.AssignedUserCount == nil || *p.AssignedUserCount != *input.AssignedUserCount) {
		return false
	}

	if p.CompletedUserCount != nil && (input.CompletedUserCount == nil || *p.CompletedUserCount != *input.CompletedUserCount) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
