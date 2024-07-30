package lifecycleworkflowdeleteditemworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SectionGroup struct {
	CreatedBy            *IdentitySet      `json:"createdBy,omitempty"`
	CreatedDateTime      *string           `json:"createdDateTime,omitempty"`
	DisplayName          *string           `json:"displayName,omitempty"`
	Id                   *string           `json:"id,omitempty"`
	LastModifiedBy       *IdentitySet      `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime *string           `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string           `json:"@odata.type,omitempty"`
	ParentNotebook       *Notebook         `json:"parentNotebook,omitempty"`
	ParentSectionGroup   *SectionGroup     `json:"parentSectionGroup,omitempty"`
	SectionGroups        *[]SectionGroup   `json:"sectionGroups,omitempty"`
	SectionGroupsUrl     *string           `json:"sectionGroupsUrl,omitempty"`
	Sections             *[]OnenoteSection `json:"sections,omitempty"`
	SectionsUrl          *string           `json:"sectionsUrl,omitempty"`
	Self                 *string           `json:"self,omitempty"`
}
