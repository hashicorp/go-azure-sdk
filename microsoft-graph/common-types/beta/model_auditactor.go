package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuditActor struct {
	ApplicationDisplayName *string             `json:"applicationDisplayName,omitempty"`
	ApplicationId          *string             `json:"applicationId,omitempty"`
	AuditActorType         *string             `json:"auditActorType,omitempty"`
	IpAddress              *string             `json:"ipAddress,omitempty"`
	ODataType              *string             `json:"@odata.type,omitempty"`
	RemoteTenantId         *string             `json:"remoteTenantId,omitempty"`
	RemoteUserId           *string             `json:"remoteUserId,omitempty"`
	ServicePrincipalName   *string             `json:"servicePrincipalName,omitempty"`
	Type                   *string             `json:"type,omitempty"`
	UserId                 *string             `json:"userId,omitempty"`
	UserPermissions        *[]string           `json:"userPermissions,omitempty"`
	UserPrincipalName      *string             `json:"userPrincipalName,omitempty"`
	UserRoleScopeTags      *[]RoleScopeTagInfo `json:"userRoleScopeTags,omitempty"`
}
