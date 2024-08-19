package listtenantconfigurationviolationsoperations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Violation struct {
	ErrorMessage *string `json:"errorMessage,omitempty"`
	Id           *string `json:"id,omitempty"`
	UserId       *string `json:"userId,omitempty"`
}
