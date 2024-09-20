package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = WindowsQualityUpdateProfileAssignment{}

type WindowsQualityUpdateProfileAssignment struct {
	// The assignment target that the quality update profile is assigned to.
	Target DeviceAndAppManagementAssignmentTarget `json:"target"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s WindowsQualityUpdateProfileAssignment) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsQualityUpdateProfileAssignment{}

func (s WindowsQualityUpdateProfileAssignment) MarshalJSON() ([]byte, error) {
	type wrapper WindowsQualityUpdateProfileAssignment
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsQualityUpdateProfileAssignment: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsQualityUpdateProfileAssignment: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.windowsQualityUpdateProfileAssignment"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsQualityUpdateProfileAssignment: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &WindowsQualityUpdateProfileAssignment{}

func (s *WindowsQualityUpdateProfileAssignment) UnmarshalJSON(bytes []byte) error {

	var decoded struct {
		Id        *string `json:"id,omitempty"`
		ODataId   *string `json:"@odata.id,omitempty"`
		ODataType *string `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling WindowsQualityUpdateProfileAssignment into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["target"]; ok {
		impl, err := UnmarshalDeviceAndAppManagementAssignmentTargetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Target' for 'WindowsQualityUpdateProfileAssignment': %+v", err)
		}
		s.Target = impl
	}

	return nil
}
