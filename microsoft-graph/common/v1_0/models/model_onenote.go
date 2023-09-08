package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Onenote struct {
	Id            *string             `json:"id,omitempty"`
	Notebooks     *[]Notebook         `json:"notebooks,omitempty"`
	ODataType     *string             `json:"@odata.type,omitempty"`
	Operations    *[]OnenoteOperation `json:"operations,omitempty"`
	Pages         *[]OnenotePage      `json:"pages,omitempty"`
	Resources     *[]OnenoteResource  `json:"resources,omitempty"`
	SectionGroups *[]SectionGroup     `json:"sectionGroups,omitempty"`
	Sections      *[]OnenoteSection   `json:"sections,omitempty"`
}
