package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommentAction struct {
	IsReply      *bool          `json:"isReply,omitempty"`
	ODataType    *string        `json:"@odata.type,omitempty"`
	ParentAuthor *IdentitySet   `json:"parentAuthor,omitempty"`
	Participants *[]IdentitySet `json:"participants,omitempty"`
}
