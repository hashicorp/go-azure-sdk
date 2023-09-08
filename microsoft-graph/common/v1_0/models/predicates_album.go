package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlbumOperationPredicate struct {
	CoverImageItemId *string
	ODataType        *string
}

func (p AlbumOperationPredicate) Matches(input Album) bool {

	if p.CoverImageItemId != nil && (input.CoverImageItemId == nil || *p.CoverImageItemId != *input.CoverImageItemId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
