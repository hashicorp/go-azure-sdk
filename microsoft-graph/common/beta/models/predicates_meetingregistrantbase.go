package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeetingRegistrantBaseOperationPredicate struct {
	Id         *string
	JoinWebUrl *string
	ODataType  *string
}

func (p MeetingRegistrantBaseOperationPredicate) Matches(input MeetingRegistrantBase) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.JoinWebUrl != nil && (input.JoinWebUrl == nil || *p.JoinWebUrl != *input.JoinWebUrl) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
