package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReactionsFacet struct {
	CommentCount *int64  `json:"commentCount,omitempty"`
	LikeCount    *int64  `json:"likeCount,omitempty"`
	ODataType    *string `json:"@odata.type,omitempty"`
	ShareCount   *int64  `json:"shareCount,omitempty"`
}
