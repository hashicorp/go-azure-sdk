package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DocumentSetVersionItemOperationPredicate struct {
	ItemId    *string
	ODataType *string
	Title     *string
	VersionId *string
}

func (p DocumentSetVersionItemOperationPredicate) Matches(input DocumentSetVersionItem) bool {

	if p.ItemId != nil && (input.ItemId == nil || *p.ItemId != *input.ItemId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Title != nil && (input.Title == nil || *p.Title != *input.Title) {
		return false
	}

	if p.VersionId != nil && (input.VersionId == nil || *p.VersionId != *input.VersionId) {
		return false
	}

	return true
}
