package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CrossTenantAccessPolicyTenantRestrictions struct {
	Applications   *CrossTenantAccessPolicyTargetConfiguration `json:"applications,omitempty"`
	Devices        *DevicesFilter                              `json:"devices,omitempty"`
	ODataType      *string                                     `json:"@odata.type,omitempty"`
	UsersAndGroups *CrossTenantAccessPolicyTargetConfiguration `json:"usersAndGroups,omitempty"`
}
