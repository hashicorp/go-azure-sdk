package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultUserRoleOverride struct {
	Id              *string                  `json:"id,omitempty"`
	IsDefault       *bool                    `json:"isDefault,omitempty"`
	ODataType       *string                  `json:"@odata.type,omitempty"`
	RolePermissions *[]UnifiedRolePermission `json:"rolePermissions,omitempty"`
}
