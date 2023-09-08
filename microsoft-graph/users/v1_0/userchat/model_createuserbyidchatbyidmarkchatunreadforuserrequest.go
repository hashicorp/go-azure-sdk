package userchat

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateUserByIdChatByIdMarkChatUnreadForUserRequest struct {
	LastMessageReadDateTime *string               `json:"lastMessageReadDateTime,omitempty"`
	User                    *TeamworkUserIdentity `json:"user,omitempty"`
}
