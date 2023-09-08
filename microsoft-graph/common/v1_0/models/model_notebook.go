package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Notebook struct {
	CreatedBy            *IdentitySet      `json:"createdBy,omitempty"`
	CreatedDateTime      *string           `json:"createdDateTime,omitempty"`
	DisplayName          *string           `json:"displayName,omitempty"`
	Id                   *string           `json:"id,omitempty"`
	IsDefault            *bool             `json:"isDefault,omitempty"`
	IsShared             *bool             `json:"isShared,omitempty"`
	LastModifiedBy       *IdentitySet      `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime *string           `json:"lastModifiedDateTime,omitempty"`
	Links                *NotebookLinks    `json:"links,omitempty"`
	ODataType            *string           `json:"@odata.type,omitempty"`
	SectionGroups        *[]SectionGroup   `json:"sectionGroups,omitempty"`
	SectionGroupsUrl     *string           `json:"sectionGroupsUrl,omitempty"`
	Sections             *[]OnenoteSection `json:"sections,omitempty"`
	SectionsUrl          *string           `json:"sectionsUrl,omitempty"`
	Self                 *string           `json:"self,omitempty"`
	UserRole             *NotebookUserRole `json:"userRole,omitempty"`
}
