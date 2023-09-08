package memailfoldermessage

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateMeMailFolderByIdMessageByIdCreateReplyRequest struct {
	Comment *string  `json:"Comment,omitempty"`
	Message *Message `json:"Message,omitempty"`
}
