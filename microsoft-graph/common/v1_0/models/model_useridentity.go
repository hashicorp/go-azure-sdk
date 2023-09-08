package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserIdentity struct {
	DisplayName       *string `json:"displayName,omitempty"`
	Id                *string `json:"id,omitempty"`
	IpAddress         *string `json:"ipAddress,omitempty"`
	ODataType         *string `json:"@odata.type,omitempty"`
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
}
