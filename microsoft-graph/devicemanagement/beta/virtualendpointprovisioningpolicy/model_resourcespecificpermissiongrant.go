package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourceSpecificPermissionGrant struct {
	ClientAppId     *string `json:"clientAppId,omitempty"`
	ClientId        *string `json:"clientId,omitempty"`
	DeletedDateTime *string `json:"deletedDateTime,omitempty"`
	Id              *string `json:"id,omitempty"`
	ODataType       *string `json:"@odata.type,omitempty"`
	Permission      *string `json:"permission,omitempty"`
	PermissionType  *string `json:"permissionType,omitempty"`
	ResourceAppId   *string `json:"resourceAppId,omitempty"`
}
