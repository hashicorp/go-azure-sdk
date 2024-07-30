package sitelistitem

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateSiteListItemCreateLinkRequest struct {
	ExpirationDateTime         *string           `json:"expirationDateTime,omitempty"`
	Message                    *string           `json:"message,omitempty"`
	Password                   *string           `json:"password,omitempty"`
	Recipients                 *[]DriveRecipient `json:"recipients,omitempty"`
	RetainInheritedPermissions *bool             `json:"retainInheritedPermissions,omitempty"`
	Scope                      *string           `json:"scope,omitempty"`
	SendNotification           *bool             `json:"sendNotification,omitempty"`
	Type                       *string           `json:"type,omitempty"`
}
