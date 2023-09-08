package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlertHistoryStateOperationPredicate struct {
	AppId           *string
	AssignedTo      *string
	ODataType       *string
	UpdatedDateTime *string
	User            *string
}

func (p AlertHistoryStateOperationPredicate) Matches(input AlertHistoryState) bool {

	if p.AppId != nil && (input.AppId == nil || *p.AppId != *input.AppId) {
		return false
	}

	if p.AssignedTo != nil && (input.AssignedTo == nil || *p.AssignedTo != *input.AssignedTo) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.UpdatedDateTime != nil && (input.UpdatedDateTime == nil || *p.UpdatedDateTime != *input.UpdatedDateTime) {
		return false
	}

	if p.User != nil && (input.User == nil || *p.User != *input.User) {
		return false
	}

	return true
}
