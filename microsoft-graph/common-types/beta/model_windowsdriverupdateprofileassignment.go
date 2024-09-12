package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = WindowsDriverUpdateProfileAssignment{}

type WindowsDriverUpdateProfileAssignment struct {
	// Base type for assignment targets.
	Target DeviceAndAppManagementAssignmentTarget `json:"target"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s WindowsDriverUpdateProfileAssignment) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsDriverUpdateProfileAssignment{}

func (s WindowsDriverUpdateProfileAssignment) MarshalJSON() ([]byte, error) {
	type wrapper WindowsDriverUpdateProfileAssignment
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsDriverUpdateProfileAssignment: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsDriverUpdateProfileAssignment: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.windowsDriverUpdateProfileAssignment"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsDriverUpdateProfileAssignment: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &WindowsDriverUpdateProfileAssignment{}

func (s *WindowsDriverUpdateProfileAssignment) UnmarshalJSON(bytes []byte) error {
	type alias WindowsDriverUpdateProfileAssignment
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into WindowsDriverUpdateProfileAssignment: %+v", err)
	}

	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling WindowsDriverUpdateProfileAssignment into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["target"]; ok {
		impl, err := UnmarshalDeviceAndAppManagementAssignmentTargetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Target' for 'WindowsDriverUpdateProfileAssignment': %+v", err)
		}
		s.Target = impl
	}
	return nil
}
