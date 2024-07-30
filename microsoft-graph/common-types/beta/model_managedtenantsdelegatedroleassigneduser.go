package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsDelegatedRoleAssignedUser struct {
	DisplayName       *string `json:"displayName,omitempty"`
	ODataType         *string `json:"@odata.type,omitempty"`
	UserEntityId      *string `json:"userEntityId,omitempty"`
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
}
