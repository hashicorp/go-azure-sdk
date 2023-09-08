package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceUpdateMessageViewpointOperationPredicate struct {
	IsArchived  *bool
	IsFavorited *bool
	IsRead      *bool
	ODataType   *string
}

func (p ServiceUpdateMessageViewpointOperationPredicate) Matches(input ServiceUpdateMessageViewpoint) bool {

	if p.IsArchived != nil && (input.IsArchived == nil || *p.IsArchived != *input.IsArchived) {
		return false
	}

	if p.IsFavorited != nil && (input.IsFavorited == nil || *p.IsFavorited != *input.IsFavorited) {
		return false
	}

	if p.IsRead != nil && (input.IsRead == nil || *p.IsRead != *input.IsRead) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
