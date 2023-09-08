package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatMessageInfoOperationPredicate struct {
	CreatedDateTime *string
	Id              *string
	IsDeleted       *bool
	ODataType       *string
}

func (p ChatMessageInfoOperationPredicate) Matches(input ChatMessageInfo) bool {

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsDeleted != nil && (input.IsDeleted == nil || *p.IsDeleted != *input.IsDeleted) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
