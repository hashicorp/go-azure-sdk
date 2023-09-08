package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConfigManagerCollection struct {
	CollectionIdentifier *string `json:"collectionIdentifier,omitempty"`
	CreatedDateTime      *string `json:"createdDateTime,omitempty"`
	DisplayName          *string `json:"displayName,omitempty"`
	HierarchyIdentifier  *string `json:"hierarchyIdentifier,omitempty"`
	HierarchyName        *string `json:"hierarchyName,omitempty"`
	Id                   *string `json:"id,omitempty"`
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string `json:"@odata.type,omitempty"`
}
