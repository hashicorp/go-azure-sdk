package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ OnenoteEntityHierarchyModel = OnenoteSection{}

type OnenoteSection struct {
	// Indicates whether this is the user's default section. Read-only.
	IsDefault nullable.Type[bool] `json:"isDefault,omitempty"`

	// Links for opening the section. The oneNoteClientURL link opens the section in the OneNote native client if it's
	// installed. The oneNoteWebURL link opens the section in OneNote on the web.
	Links *SectionLinks `json:"links,omitempty"`

	// The collection of pages in the section. Read-only. Nullable.
	Pages *[]OnenotePage `json:"pages,omitempty"`

	// The pages endpoint where you can get details for all the pages in the section. Read-only.
	PagesUrl nullable.Type[string] `json:"pagesUrl,omitempty"`

	// The notebook that contains the section. Read-only.
	ParentNotebook *Notebook `json:"parentNotebook,omitempty"`

	// The section group that contains the section. Read-only.
	ParentSectionGroup *SectionGroup `json:"parentSectionGroup,omitempty"`

	// Fields inherited from OnenoteEntityHierarchyModel

	// Identity of the user, device, and application that created the item. Read-only.
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`

	// The name of the notebook.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Identity of the user, device, and application that created the item. Read-only.
	LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`

	// The date and time when the notebook was last modified. The timestamp represents date and time information using ISO
	// 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// Fields inherited from OnenoteEntitySchemaObjectModel

	// The date and time when the page was created. The timestamp represents date and time information using ISO 8601 format
	// and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Fields inherited from OnenoteEntityBaseModel

	// The endpoint where you can get details about the page. Read-only.
	Self nullable.Type[string] `json:"self,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s OnenoteSection) OnenoteEntityHierarchyModel() BaseOnenoteEntityHierarchyModelImpl {
	return BaseOnenoteEntityHierarchyModelImpl{
		CreatedBy:            s.CreatedBy,
		DisplayName:          s.DisplayName,
		LastModifiedBy:       s.LastModifiedBy,
		LastModifiedDateTime: s.LastModifiedDateTime,
		CreatedDateTime:      s.CreatedDateTime,
		Self:                 s.Self,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s OnenoteSection) OnenoteEntitySchemaObjectModel() BaseOnenoteEntitySchemaObjectModelImpl {
	return BaseOnenoteEntitySchemaObjectModelImpl{
		CreatedDateTime: s.CreatedDateTime,
		Self:            s.Self,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s OnenoteSection) OnenoteEntityBaseModel() BaseOnenoteEntityBaseModelImpl {
	return BaseOnenoteEntityBaseModelImpl{
		Self:      s.Self,
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s OnenoteSection) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = OnenoteSection{}

func (s OnenoteSection) MarshalJSON() ([]byte, error) {
	type wrapper OnenoteSection
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling OnenoteSection: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling OnenoteSection: %+v", err)
	}

	delete(decoded, "isDefault")
	delete(decoded, "pages")
	delete(decoded, "pagesUrl")
	delete(decoded, "parentNotebook")
	delete(decoded, "parentSectionGroup")
	decoded["@odata.type"] = "#microsoft.graph.onenoteSection"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling OnenoteSection: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &OnenoteSection{}

func (s *OnenoteSection) UnmarshalJSON(bytes []byte) error {

	var decoded struct {
		IsDefault            nullable.Type[bool]   `json:"isDefault,omitempty"`
		Links                *SectionLinks         `json:"links,omitempty"`
		Pages                *[]OnenotePage        `json:"pages,omitempty"`
		PagesUrl             nullable.Type[string] `json:"pagesUrl,omitempty"`
		ParentNotebook       *Notebook             `json:"parentNotebook,omitempty"`
		ParentSectionGroup   *SectionGroup         `json:"parentSectionGroup,omitempty"`
		CreatedBy            *IdentitySet          `json:"createdBy,omitempty"`
		DisplayName          nullable.Type[string] `json:"displayName,omitempty"`
		LastModifiedBy       *IdentitySet          `json:"lastModifiedBy,omitempty"`
		LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`
		CreatedDateTime      nullable.Type[string] `json:"createdDateTime,omitempty"`
		Self                 nullable.Type[string] `json:"self,omitempty"`
		Id                   *string               `json:"id,omitempty"`
		ODataId              *string               `json:"@odata.id,omitempty"`
		ODataType            *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.IsDefault = decoded.IsDefault
	s.Links = decoded.Links
	s.Pages = decoded.Pages
	s.PagesUrl = decoded.PagesUrl
	s.ParentNotebook = decoded.ParentNotebook
	s.ParentSectionGroup = decoded.ParentSectionGroup
	s.CreatedDateTime = decoded.CreatedDateTime
	s.DisplayName = decoded.DisplayName
	s.Id = decoded.Id
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Self = decoded.Self

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling OnenoteSection into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'OnenoteSection': %+v", err)
		}
		s.CreatedBy = &impl
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'OnenoteSection': %+v", err)
		}
		s.LastModifiedBy = &impl
	}

	return nil
}
