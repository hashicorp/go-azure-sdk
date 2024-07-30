package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GcpRole struct {
	DisplayName *string             `json:"displayName,omitempty"`
	ExternalId  *string             `json:"externalId,omitempty"`
	GcpRoleType *GcpRoleGcpRoleType `json:"gcpRoleType,omitempty"`
	Id          *string             `json:"id,omitempty"`
	ODataType   *string             `json:"@odata.type,omitempty"`
	Scopes      *[]GcpScope         `json:"scopes,omitempty"`
}
