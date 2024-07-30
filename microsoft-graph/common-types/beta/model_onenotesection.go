package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenoteSection struct {
	CreatedBy            *IdentitySet   `json:"createdBy,omitempty"`
	CreatedDateTime      *string        `json:"createdDateTime,omitempty"`
	DisplayName          *string        `json:"displayName,omitempty"`
	Id                   *string        `json:"id,omitempty"`
	IsDefault            *bool          `json:"isDefault,omitempty"`
	LastModifiedBy       *IdentitySet   `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime *string        `json:"lastModifiedDateTime,omitempty"`
	Links                *SectionLinks  `json:"links,omitempty"`
	ODataType            *string        `json:"@odata.type,omitempty"`
	Pages                *[]OnenotePage `json:"pages,omitempty"`
	PagesUrl             *string        `json:"pagesUrl,omitempty"`
	ParentNotebook       *Notebook      `json:"parentNotebook,omitempty"`
	ParentSectionGroup   *SectionGroup  `json:"parentSectionGroup,omitempty"`
	Self                 *string        `json:"self,omitempty"`
}
