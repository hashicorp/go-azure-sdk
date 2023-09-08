package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReactionsFacetOperationPredicate struct {
	CommentCount *int64
	LikeCount    *int64
	ODataType    *string
	ShareCount   *int64
}

func (p ReactionsFacetOperationPredicate) Matches(input ReactionsFacet) bool {

	if p.CommentCount != nil && (input.CommentCount == nil || *p.CommentCount != *input.CommentCount) {
		return false
	}

	if p.LikeCount != nil && (input.LikeCount == nil || *p.LikeCount != *input.LikeCount) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ShareCount != nil && (input.ShareCount == nil || *p.ShareCount != *input.ShareCount) {
		return false
	}

	return true
}
