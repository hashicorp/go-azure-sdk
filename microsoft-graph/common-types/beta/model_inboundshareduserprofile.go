package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InboundSharedUserProfile struct {
	DisplayName       *string `json:"displayName,omitempty"`
	HomeTenantId      *string `json:"homeTenantId,omitempty"`
	ODataType         *string `json:"@odata.type,omitempty"`
	UserId            *string `json:"userId,omitempty"`
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
}
