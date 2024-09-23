package driverootlistitem

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateDriveRootListItemLinkRequest struct {
	ExpirationDateTime         nullable.Type[string]  `json:"expirationDateTime,omitempty"`
	Message                    nullable.Type[string]  `json:"message,omitempty"`
	Password                   nullable.Type[string]  `json:"password,omitempty"`
	Recipients                 *[]beta.DriveRecipient `json:"recipients,omitempty"`
	RetainInheritedPermissions nullable.Type[bool]    `json:"retainInheritedPermissions,omitempty"`
	Scope                      nullable.Type[string]  `json:"scope,omitempty"`
	SendNotification           nullable.Type[bool]    `json:"sendNotification,omitempty"`
	Type                       nullable.Type[string]  `json:"type,omitempty"`
}
