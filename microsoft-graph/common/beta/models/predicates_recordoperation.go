package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecordOperationOperationPredicate struct {
	ClientContext        *string
	Id                   *string
	ODataType            *string
	RecordingAccessToken *string
	RecordingLocation    *string
}

func (p RecordOperationOperationPredicate) Matches(input RecordOperation) bool {

	if p.ClientContext != nil && (input.ClientContext == nil || *p.ClientContext != *input.ClientContext) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RecordingAccessToken != nil && (input.RecordingAccessToken == nil || *p.RecordingAccessToken != *input.RecordingAccessToken) {
		return false
	}

	if p.RecordingLocation != nil && (input.RecordingLocation == nil || *p.RecordingLocation != *input.RecordingLocation) {
		return false
	}

	return true
}
