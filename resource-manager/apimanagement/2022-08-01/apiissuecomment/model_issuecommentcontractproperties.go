package apiissuecomment

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IssueCommentContractProperties struct {
	CreatedDate *string `json:"createdDate,omitempty"`
	Text        string  `json:"text"`
	UserId      string  `json:"userId"`
}
