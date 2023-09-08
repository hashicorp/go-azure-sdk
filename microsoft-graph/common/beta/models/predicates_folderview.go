package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FolderViewOperationPredicate struct {
	ODataType *string
	SortBy    *string
	SortOrder *string
	ViewType  *string
}

func (p FolderViewOperationPredicate) Matches(input FolderView) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.SortBy != nil && (input.SortBy == nil || *p.SortBy != *input.SortBy) {
		return false
	}

	if p.SortOrder != nil && (input.SortOrder == nil || *p.SortOrder != *input.SortOrder) {
		return false
	}

	if p.ViewType != nil && (input.ViewType == nil || *p.ViewType != *input.ViewType) {
		return false
	}

	return true
}
