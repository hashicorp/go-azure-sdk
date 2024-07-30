package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScopedRoleMembership struct {
	AdministrativeUnitId *string   `json:"administrativeUnitId,omitempty"`
	Id                   *string   `json:"id,omitempty"`
	ODataType            *string   `json:"@odata.type,omitempty"`
	RoleId               *string   `json:"roleId,omitempty"`
	RoleMemberInfo       *Identity `json:"roleMemberInfo,omitempty"`
}
