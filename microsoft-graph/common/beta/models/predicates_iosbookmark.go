package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosBookmarkOperationPredicate struct {
	BookmarkFolder *string
	DisplayName    *string
	ODataType      *string
	Url            *string
}

func (p IosBookmarkOperationPredicate) Matches(input IosBookmark) bool {

	if p.BookmarkFolder != nil && (input.BookmarkFolder == nil || *p.BookmarkFolder != *input.BookmarkFolder) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Url != nil && (input.Url == nil || *p.Url != *input.Url) {
		return false
	}

	return true
}
