package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryObjectPartnerReference struct {
	DeletedDateTime         *string `json:"deletedDateTime,omitempty"`
	Description             *string `json:"description,omitempty"`
	DisplayName             *string `json:"displayName,omitempty"`
	ExternalPartnerTenantId *string `json:"externalPartnerTenantId,omitempty"`
	Id                      *string `json:"id,omitempty"`
	ODataType               *string `json:"@odata.type,omitempty"`
	ObjectType              *string `json:"objectType,omitempty"`
}
