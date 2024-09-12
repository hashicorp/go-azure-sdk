package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ManagedEBookAssignment = IosVppEBookAssignment{}

type IosVppEBookAssignment struct {

	// Fields inherited from ManagedEBookAssignment

	// Possible values for the install intent chosen by the admin.
	InstallIntent *InstallIntent `json:"installIntent,omitempty"`

	// The assignment target for eBook.
	Target DeviceAndAppManagementAssignmentTarget `json:"target"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s IosVppEBookAssignment) ManagedEBookAssignment() BaseManagedEBookAssignmentImpl {
	return BaseManagedEBookAssignmentImpl{
		InstallIntent: s.InstallIntent,
		Target:        s.Target,
		Id:            s.Id,
		ODataId:       s.ODataId,
		ODataType:     s.ODataType,
	}
}

func (s IosVppEBookAssignment) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IosVppEBookAssignment{}

func (s IosVppEBookAssignment) MarshalJSON() ([]byte, error) {
	type wrapper IosVppEBookAssignment
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IosVppEBookAssignment: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IosVppEBookAssignment: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.iosVppEBookAssignment"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IosVppEBookAssignment: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &IosVppEBookAssignment{}

func (s *IosVppEBookAssignment) UnmarshalJSON(bytes []byte) error {
	type alias IosVppEBookAssignment
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into IosVppEBookAssignment: %+v", err)
	}

	s.Id = decoded.Id
	s.InstallIntent = decoded.InstallIntent
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling IosVppEBookAssignment into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["target"]; ok {
		impl, err := UnmarshalDeviceAndAppManagementAssignmentTargetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Target' for 'IosVppEBookAssignment': %+v", err)
		}
		s.Target = impl
	}
	return nil
}
