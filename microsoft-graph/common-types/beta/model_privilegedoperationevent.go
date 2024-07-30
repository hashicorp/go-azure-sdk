package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedOperationEvent struct {
	AdditionalInformation *string `json:"additionalInformation,omitempty"`
	CreationDateTime      *string `json:"creationDateTime,omitempty"`
	ExpirationDateTime    *string `json:"expirationDateTime,omitempty"`
	Id                    *string `json:"id,omitempty"`
	ODataType             *string `json:"@odata.type,omitempty"`
	ReferenceKey          *string `json:"referenceKey,omitempty"`
	ReferenceSystem       *string `json:"referenceSystem,omitempty"`
	RequestType           *string `json:"requestType,omitempty"`
	RequestorId           *string `json:"requestorId,omitempty"`
	RequestorName         *string `json:"requestorName,omitempty"`
	RoleId                *string `json:"roleId,omitempty"`
	RoleName              *string `json:"roleName,omitempty"`
	TenantId              *string `json:"tenantId,omitempty"`
	UserId                *string `json:"userId,omitempty"`
	UserMail              *string `json:"userMail,omitempty"`
	UserName              *string `json:"userName,omitempty"`
}
