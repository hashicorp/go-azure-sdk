package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatRenamedEventMessageDetail struct {
	ChatDisplayName *string      `json:"chatDisplayName,omitempty"`
	ChatId          *string      `json:"chatId,omitempty"`
	Initiator       *IdentitySet `json:"initiator,omitempty"`
	ODataType       *string      `json:"@odata.type,omitempty"`
}
