package chat

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnhideChatForUserRequest struct {
	TenantId nullable.Type[string]      `json:"tenantId,omitempty"`
	User     *beta.TeamworkUserIdentity `json:"user,omitempty"`
}
