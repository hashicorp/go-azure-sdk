package groupsitepermission

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListGroupByIdSiteByIdPermissionByIdGrantsRequest struct {
	Recipients *[]DriveRecipient `json:"recipients,omitempty"`
	Roles      *[]string         `json:"roles,omitempty"`
}
