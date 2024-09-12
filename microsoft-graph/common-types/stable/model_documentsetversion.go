package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ListItemVersion = DocumentSetVersion{}

type DocumentSetVersion struct {
	// Comment about the captured version.
	Comment nullable.Type[string] `json:"comment,omitempty"`

	// User who captured the version.
	CreatedBy IdentitySet `json:"createdBy"`

	// Date and time when this version was created.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Items within the document set that are captured as part of this version.
	Items *[]DocumentSetVersionItem `json:"items,omitempty"`

	// If true, minor versions of items are also captured; otherwise, only major versions are captured. The default value is
	// false.
	ShouldCaptureMinorVersion nullable.Type[bool] `json:"shouldCaptureMinorVersion,omitempty"`

	// Fields inherited from ListItemVersion

	// A collection of the fields and values for this version of the list item.
	Fields *FieldValueSet `json:"fields,omitempty"`

	// Fields inherited from BaseItemVersion

	// Identity of the user which last modified the version. Read-only.
	LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`

	// Date and time the version was last modified. Read-only.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// Indicates the publication status of this particular version. Read-only.
	Publication *PublicationFacet `json:"publication,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s DocumentSetVersion) ListItemVersion() BaseListItemVersionImpl {
	return BaseListItemVersionImpl{
		Fields:               s.Fields,
		LastModifiedBy:       s.LastModifiedBy,
		LastModifiedDateTime: s.LastModifiedDateTime,
		Publication:          s.Publication,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s DocumentSetVersion) BaseItemVersion() BaseBaseItemVersionImpl {
	return BaseBaseItemVersionImpl{
		LastModifiedBy:       s.LastModifiedBy,
		LastModifiedDateTime: s.LastModifiedDateTime,
		Publication:          s.Publication,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s DocumentSetVersion) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DocumentSetVersion{}

func (s DocumentSetVersion) MarshalJSON() ([]byte, error) {
	type wrapper DocumentSetVersion
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DocumentSetVersion: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DocumentSetVersion: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.documentSetVersion"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DocumentSetVersion: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &DocumentSetVersion{}

func (s *DocumentSetVersion) UnmarshalJSON(bytes []byte) error {
	type alias DocumentSetVersion
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into DocumentSetVersion: %+v", err)
	}

	s.Comment = decoded.Comment
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Fields = decoded.Fields
	s.Id = decoded.Id
	s.Items = decoded.Items
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Publication = decoded.Publication
	s.ShouldCaptureMinorVersion = decoded.ShouldCaptureMinorVersion

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling DocumentSetVersion into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'DocumentSetVersion': %+v", err)
		}
		s.CreatedBy = impl
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'DocumentSetVersion': %+v", err)
		}
		s.LastModifiedBy = &impl
	}
	return nil
}
