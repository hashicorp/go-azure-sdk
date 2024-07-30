package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppRoleAssignment struct {
	AppRoleId            *string `json:"appRoleId,omitempty"`
	CreationTimestamp    *string `json:"creationTimestamp,omitempty"`
	Id                   *string `json:"id,omitempty"`
	ODataType            *string `json:"@odata.type,omitempty"`
	PrincipalDisplayName *string `json:"principalDisplayName,omitempty"`
	PrincipalId          *string `json:"principalId,omitempty"`
	PrincipalType        *string `json:"principalType,omitempty"`
	ResourceDisplayName  *string `json:"resourceDisplayName,omitempty"`
	ResourceId           *string `json:"resourceId,omitempty"`
}
