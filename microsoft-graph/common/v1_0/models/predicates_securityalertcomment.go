package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAlertCommentOperationPredicate struct {
	Comment              *string
	CreatedByDisplayName *string
	CreatedDateTime      *string
	ODataType            *string
}

func (p SecurityAlertCommentOperationPredicate) Matches(input SecurityAlertComment) bool {

	if p.Comment != nil && (input.Comment == nil || *p.Comment != *input.Comment) {
		return false
	}

	if p.CreatedByDisplayName != nil && (input.CreatedByDisplayName == nil || *p.CreatedByDisplayName != *input.CreatedByDisplayName) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
