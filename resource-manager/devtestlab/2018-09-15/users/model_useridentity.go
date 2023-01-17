package users

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserIdentity struct {
	AppId         *string `json:"appId,omitempty"`
	ObjectId      *string `json:"objectId,omitempty"`
	PrincipalId   *string `json:"principalId,omitempty"`
	PrincipalName *string `json:"principalName,omitempty"`
	TenantId      *string `json:"tenantId,omitempty"`
}
