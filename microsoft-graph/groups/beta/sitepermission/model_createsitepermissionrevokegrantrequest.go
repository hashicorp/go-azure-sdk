package sitepermission

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateSitePermissionRevokeGrantRequest struct {
	Grantees *[]DriveRecipient `json:"grantees,omitempty"`
}
