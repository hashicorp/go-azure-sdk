package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultUserRolePermissions struct {
	AllowedToCreateApps                      *bool     `json:"allowedToCreateApps,omitempty"`
	AllowedToCreateSecurityGroups            *bool     `json:"allowedToCreateSecurityGroups,omitempty"`
	AllowedToCreateTenants                   *bool     `json:"allowedToCreateTenants,omitempty"`
	AllowedToReadBitlockerKeysForOwnedDevice *bool     `json:"allowedToReadBitlockerKeysForOwnedDevice,omitempty"`
	AllowedToReadOtherUsers                  *bool     `json:"allowedToReadOtherUsers,omitempty"`
	ODataType                                *string   `json:"@odata.type,omitempty"`
	PermissionGrantPoliciesAssigned          *[]string `json:"permissionGrantPoliciesAssigned,omitempty"`
}
