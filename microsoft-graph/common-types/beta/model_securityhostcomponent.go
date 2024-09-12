package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SecurityArtifact = SecurityHostComponent{}

type SecurityHostComponent struct {
	// The type of component that was detected (for example, Operating System, Framework, Remote Access, or Server).
	Category nullable.Type[string] `json:"category,omitempty"`

	// The first date and time when this web component was observed by Microsoft Defender Threat Intelligence. The Timestamp
	// type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC
	// on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	FirstSeenDateTime *string `json:"firstSeenDateTime,omitempty"`

	Host *SecurityHost `json:"host,omitempty"`

	// The most recent date and time when this web component was observed by Microsoft Defender Threat Intelligence. The
	// Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	LastSeenDateTime *string `json:"lastSeenDateTime,omitempty"`

	// A name running on the artifact, for example, Microsoft IIS.
	Name *string `json:"name,omitempty"`

	// The component version running on the artifact, for example, v8.5. This shouldn't be assumed to be strictly numerical.
	Version nullable.Type[string] `json:"version,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s SecurityHostComponent) SecurityArtifact() BaseSecurityArtifactImpl {
	return BaseSecurityArtifactImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s SecurityHostComponent) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityHostComponent{}

func (s SecurityHostComponent) MarshalJSON() ([]byte, error) {
	type wrapper SecurityHostComponent
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityHostComponent: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityHostComponent: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.security.hostComponent"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityHostComponent: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &SecurityHostComponent{}

func (s *SecurityHostComponent) UnmarshalJSON(bytes []byte) error {
	type alias SecurityHostComponent
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into SecurityHostComponent: %+v", err)
	}

	s.Category = decoded.Category
	s.FirstSeenDateTime = decoded.FirstSeenDateTime
	s.Id = decoded.Id
	s.LastSeenDateTime = decoded.LastSeenDateTime
	s.Name = decoded.Name
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Version = decoded.Version

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SecurityHostComponent into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["host"]; ok {
		impl, err := UnmarshalSecurityHostImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Host' for 'SecurityHostComponent': %+v", err)
		}
		s.Host = &impl
	}
	return nil
}
