package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatViewpointOperationPredicate struct {
	IsHidden                *bool
	LastMessageReadDateTime *string
	ODataType               *string
}

func (p ChatViewpointOperationPredicate) Matches(input ChatViewpoint) bool {

	if p.IsHidden != nil && (input.IsHidden == nil || *p.IsHidden != *input.IsHidden) {
		return false
	}

	if p.LastMessageReadDateTime != nil && (input.LastMessageReadDateTime == nil || *p.LastMessageReadDateTime != *input.LastMessageReadDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
