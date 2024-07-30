package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuditUserIdentity struct {
	DisplayName       *string `json:"displayName,omitempty"`
	HomeTenantId      *string `json:"homeTenantId,omitempty"`
	HomeTenantName    *string `json:"homeTenantName,omitempty"`
	Id                *string `json:"id,omitempty"`
	IpAddress         *string `json:"ipAddress,omitempty"`
	ODataType         *string `json:"@odata.type,omitempty"`
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
}
