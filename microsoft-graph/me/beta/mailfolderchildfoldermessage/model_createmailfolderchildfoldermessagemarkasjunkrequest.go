package mailfolderchildfoldermessage

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateMailFolderChildFolderMessageMarkAsJunkRequest struct {
	MoveToJunk *bool `json:"MoveToJunk,omitempty"`
}
