package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CopyNotebookModel struct {
	CreatedBy              *string                    `json:"createdBy,omitempty"`
	CreatedByIdentity      *IdentitySet               `json:"createdByIdentity,omitempty"`
	CreatedTime            *string                    `json:"createdTime,omitempty"`
	Id                     *string                    `json:"id,omitempty"`
	IsDefault              *bool                      `json:"isDefault,omitempty"`
	IsShared               *bool                      `json:"isShared,omitempty"`
	LastModifiedBy         *string                    `json:"lastModifiedBy,omitempty"`
	LastModifiedByIdentity *IdentitySet               `json:"lastModifiedByIdentity,omitempty"`
	LastModifiedTime       *string                    `json:"lastModifiedTime,omitempty"`
	Links                  *NotebookLinks             `json:"links,omitempty"`
	Name                   *string                    `json:"name,omitempty"`
	ODataType              *string                    `json:"@odata.type,omitempty"`
	SectionGroupsUrl       *string                    `json:"sectionGroupsUrl,omitempty"`
	SectionsUrl            *string                    `json:"sectionsUrl,omitempty"`
	Self                   *string                    `json:"self,omitempty"`
	UserRole               *CopyNotebookModelUserRole `json:"userRole,omitempty"`
}
