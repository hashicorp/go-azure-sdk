package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Command struct {
	AppServiceName    *string          `json:"appServiceName,omitempty"`
	Error             *string          `json:"error,omitempty"`
	Id                *string          `json:"id,omitempty"`
	ODataType         *string          `json:"@odata.type,omitempty"`
	PackageFamilyName *string          `json:"packageFamilyName,omitempty"`
	Payload           *PayloadRequest  `json:"payload,omitempty"`
	PermissionTicket  *string          `json:"permissionTicket,omitempty"`
	PostBackUri       *string          `json:"postBackUri,omitempty"`
	Responsepayload   *PayloadResponse `json:"responsepayload,omitempty"`
	Status            *string          `json:"status,omitempty"`
	Type              *string          `json:"type,omitempty"`
}
