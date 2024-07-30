package lifecycleworkflowworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatMessageMentionedIdentitySet struct {
	Application  *Identity                     `json:"application,omitempty"`
	Conversation *TeamworkConversationIdentity `json:"conversation,omitempty"`
	Device       *Identity                     `json:"device,omitempty"`
	ODataType    *string                       `json:"@odata.type,omitempty"`
	User         *Identity                     `json:"user,omitempty"`
}
