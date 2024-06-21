package usersession

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserSessionProperties struct {
	ActiveDirectoryUserName *string          `json:"activeDirectoryUserName,omitempty"`
	ApplicationType         *ApplicationType `json:"applicationType,omitempty"`
	CreateTime              *string          `json:"createTime,omitempty"`
	ObjectId                *string          `json:"objectId,omitempty"`
	SessionState            *SessionState    `json:"sessionState,omitempty"`
	UserPrincipalName       *string          `json:"userPrincipalName,omitempty"`
}
