package mechat

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateMeChatByIdMarkChatUnreadForUserRequest struct {
	LastMessageReadDateTime *string               `json:"lastMessageReadDateTime,omitempty"`
	TenantId                *string               `json:"tenantId,omitempty"`
	User                    *TeamworkUserIdentity `json:"user,omitempty"`
}
