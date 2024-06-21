package apiissue

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IssueUpdateContractProperties struct {
	ApiId       *string `json:"apiId,omitempty"`
	CreatedDate *string `json:"createdDate,omitempty"`
	Description *string `json:"description,omitempty"`
	State       *State  `json:"state,omitempty"`
	Title       *string `json:"title,omitempty"`
	UserId      *string `json:"userId,omitempty"`
}
