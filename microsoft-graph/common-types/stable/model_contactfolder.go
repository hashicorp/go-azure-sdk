package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContactFolder struct {
	ChildFolders                  *[]ContactFolder                     `json:"childFolders,omitempty"`
	Contacts                      *[]Contact                           `json:"contacts,omitempty"`
	DisplayName                   *string                              `json:"displayName,omitempty"`
	Id                            *string                              `json:"id,omitempty"`
	MultiValueExtendedProperties  *[]MultiValueLegacyExtendedProperty  `json:"multiValueExtendedProperties,omitempty"`
	ODataType                     *string                              `json:"@odata.type,omitempty"`
	ParentFolderId                *string                              `json:"parentFolderId,omitempty"`
	SingleValueExtendedProperties *[]SingleValueLegacyExtendedProperty `json:"singleValueExtendedProperties,omitempty"`
}
