package driveroot

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListDriveRootInvitesRequest struct {
	ExpirationDateTime         nullable.Type[string]    `json:"expirationDateTime,omitempty"`
	Message                    nullable.Type[string]    `json:"message,omitempty"`
	Password                   nullable.Type[string]    `json:"password,omitempty"`
	Recipients                 *[]stable.DriveRecipient `json:"recipients,omitempty"`
	RequireSignIn              nullable.Type[bool]      `json:"requireSignIn,omitempty"`
	RetainInheritedPermissions nullable.Type[bool]      `json:"retainInheritedPermissions,omitempty"`
	Roles                      *[]string                `json:"roles,omitempty"`
	SendInvitation             nullable.Type[bool]      `json:"sendInvitation,omitempty"`
}
