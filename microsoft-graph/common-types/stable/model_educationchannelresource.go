package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EducationResource = EducationChannelResource{}

type EducationChannelResource struct {
	// URL of the channel resource.
	Url nullable.Type[string] `json:"url,omitempty"`

	// Fields inherited from EducationResource

	// The individual who created the resource.
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`

	// Moment in time when the resource was created. The Timestamp type represents date and time information using ISO 8601
	// format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Display name of resource.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The last user to modify the resource.
	LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`

	// Moment in time when the resource was last modified. The Timestamp type represents date and time information using ISO
	// 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s EducationChannelResource) EducationResource() BaseEducationResourceImpl {
	return BaseEducationResourceImpl{
		CreatedBy:            s.CreatedBy,
		CreatedDateTime:      s.CreatedDateTime,
		DisplayName:          s.DisplayName,
		LastModifiedBy:       s.LastModifiedBy,
		LastModifiedDateTime: s.LastModifiedDateTime,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

var _ json.Marshaler = EducationChannelResource{}

func (s EducationChannelResource) MarshalJSON() ([]byte, error) {
	type wrapper EducationChannelResource
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EducationChannelResource: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EducationChannelResource: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.educationChannelResource"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EducationChannelResource: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &EducationChannelResource{}

func (s *EducationChannelResource) UnmarshalJSON(bytes []byte) error {
	type alias EducationChannelResource
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into EducationChannelResource: %+v", err)
	}

	s.CreatedDateTime = decoded.CreatedDateTime
	s.DisplayName = decoded.DisplayName
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Url = decoded.Url

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling EducationChannelResource into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'EducationChannelResource': %+v", err)
		}
		s.CreatedBy = &impl
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'EducationChannelResource': %+v", err)
		}
		s.LastModifiedBy = &impl
	}
	return nil
}
