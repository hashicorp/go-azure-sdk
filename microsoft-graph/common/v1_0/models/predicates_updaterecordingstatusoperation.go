package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateRecordingStatusOperationOperationPredicate struct {
	ClientContext *string
	Id            *string
	ODataType     *string
}

func (p UpdateRecordingStatusOperationOperationPredicate) Matches(input UpdateRecordingStatusOperation) bool {

	if p.ClientContext != nil && (input.ClientContext == nil || *p.ClientContext != *input.ClientContext) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
