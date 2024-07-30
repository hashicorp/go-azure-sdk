package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConversationMemberRoleUpdatedEventMessageDetail struct {
	ConversationMemberRoles *[]string             `json:"conversationMemberRoles,omitempty"`
	ConversationMemberUser  *TeamworkUserIdentity `json:"conversationMemberUser,omitempty"`
	Initiator               *IdentitySet          `json:"initiator,omitempty"`
	ODataType               *string               `json:"@odata.type,omitempty"`
}
