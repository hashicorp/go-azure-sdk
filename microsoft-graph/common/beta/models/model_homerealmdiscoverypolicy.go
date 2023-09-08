package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HomeRealmDiscoveryPolicy struct {
	AppliesTo             *[]DirectoryObject `json:"appliesTo,omitempty"`
	Definition            *[]string          `json:"definition,omitempty"`
	DeletedDateTime       *string            `json:"deletedDateTime,omitempty"`
	Description           *string            `json:"description,omitempty"`
	DisplayName           *string            `json:"displayName,omitempty"`
	Id                    *string            `json:"id,omitempty"`
	IsOrganizationDefault *bool              `json:"isOrganizationDefault,omitempty"`
	ODataType             *string            `json:"@odata.type,omitempty"`
}
