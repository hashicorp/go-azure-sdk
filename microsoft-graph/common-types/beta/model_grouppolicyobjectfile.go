package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyObjectFile struct {
	Content              *string   `json:"content,omitempty"`
	CreatedDateTime      *string   `json:"createdDateTime,omitempty"`
	GroupPolicyObjectId  *string   `json:"groupPolicyObjectId,omitempty"`
	Id                   *string   `json:"id,omitempty"`
	LastModifiedDateTime *string   `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string   `json:"@odata.type,omitempty"`
	OuDistinguishedName  *string   `json:"ouDistinguishedName,omitempty"`
	RoleScopeTagIds      *[]string `json:"roleScopeTagIds,omitempty"`
}
