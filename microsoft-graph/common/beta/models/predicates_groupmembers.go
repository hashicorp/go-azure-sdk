package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupMembersOperationPredicate struct {
	Description *string
	Id          *string
	IsBackup    *bool
	ODataType   *string
}

func (p GroupMembersOperationPredicate) Matches(input GroupMembers) bool {

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsBackup != nil && (input.IsBackup == nil || *p.IsBackup != *input.IsBackup) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
