package mechat

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateMeChatByIdHideForUserRequest struct {
	TenantId *string               `json:"tenantId,omitempty"`
	User     *TeamworkUserIdentity `json:"user,omitempty"`
}
