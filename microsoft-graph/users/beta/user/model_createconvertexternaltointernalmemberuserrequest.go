package user

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateConvertExternalToInternalMemberUserRequest struct {
	Mail              nullable.Type[string] `json:"mail,omitempty"`
	PasswordProfile   *beta.PasswordProfile `json:"passwordProfile,omitempty"`
	UserPrincipalName nullable.Type[string] `json:"userPrincipalName,omitempty"`
}
