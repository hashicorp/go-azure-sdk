package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarPermission struct {
	AllowedRoles         *CalendarPermissionAllowedRoles `json:"allowedRoles,omitempty"`
	EmailAddress         *EmailAddress                   `json:"emailAddress,omitempty"`
	Id                   *string                         `json:"id,omitempty"`
	IsInsideOrganization *bool                           `json:"isInsideOrganization,omitempty"`
	IsRemovable          *bool                           `json:"isRemovable,omitempty"`
	ODataType            *string                         `json:"@odata.type,omitempty"`
	Role                 *CalendarPermissionRole         `json:"role,omitempty"`
}
