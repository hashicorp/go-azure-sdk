package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConvertIdResultOperationPredicate struct {
	ODataType *string
	SourceId  *string
	TargetId  *string
}

func (p ConvertIdResultOperationPredicate) Matches(input ConvertIdResult) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.SourceId != nil && (input.SourceId == nil || *p.SourceId != *input.SourceId) {
		return false
	}

	if p.TargetId != nil && (input.TargetId == nil || *p.TargetId != *input.TargetId) {
		return false
	}

	return true
}
