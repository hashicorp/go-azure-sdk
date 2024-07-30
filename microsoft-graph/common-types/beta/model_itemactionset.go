package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ItemActionSet struct {
	Comment   *CommentAction `json:"comment,omitempty"`
	Create    *CreateAction  `json:"create,omitempty"`
	Delete    *DeleteAction  `json:"delete,omitempty"`
	Edit      *EditAction    `json:"edit,omitempty"`
	Mention   *MentionAction `json:"mention,omitempty"`
	Move      *MoveAction    `json:"move,omitempty"`
	ODataType *string        `json:"@odata.type,omitempty"`
	Rename    *RenameAction  `json:"rename,omitempty"`
	Restore   *RestoreAction `json:"restore,omitempty"`
	Share     *ShareAction   `json:"share,omitempty"`
	Version   *VersionAction `json:"version,omitempty"`
}
