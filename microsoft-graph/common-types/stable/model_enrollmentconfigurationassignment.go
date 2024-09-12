package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = EnrollmentConfigurationAssignment{}

type EnrollmentConfigurationAssignment struct {
	// Represents an assignment to managed devices in the tenant
	Target DeviceAndAppManagementAssignmentTarget `json:"target"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s EnrollmentConfigurationAssignment) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EnrollmentConfigurationAssignment{}

func (s EnrollmentConfigurationAssignment) MarshalJSON() ([]byte, error) {
	type wrapper EnrollmentConfigurationAssignment
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EnrollmentConfigurationAssignment: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EnrollmentConfigurationAssignment: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.enrollmentConfigurationAssignment"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EnrollmentConfigurationAssignment: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &EnrollmentConfigurationAssignment{}

func (s *EnrollmentConfigurationAssignment) UnmarshalJSON(bytes []byte) error {
	type alias EnrollmentConfigurationAssignment
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into EnrollmentConfigurationAssignment: %+v", err)
	}

	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling EnrollmentConfigurationAssignment into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["target"]; ok {
		impl, err := UnmarshalDeviceAndAppManagementAssignmentTargetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Target' for 'EnrollmentConfigurationAssignment': %+v", err)
		}
		s.Target = impl
	}
	return nil
}
