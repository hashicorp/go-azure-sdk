package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ManagedDeviceMobileAppConfigurationAssignment{}

type ManagedDeviceMobileAppConfigurationAssignment struct {
	// Assignment target that the T&C policy is assigned to.
	Target DeviceAndAppManagementAssignmentTarget `json:"target"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s ManagedDeviceMobileAppConfigurationAssignment) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ManagedDeviceMobileAppConfigurationAssignment{}

func (s ManagedDeviceMobileAppConfigurationAssignment) MarshalJSON() ([]byte, error) {
	type wrapper ManagedDeviceMobileAppConfigurationAssignment
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedDeviceMobileAppConfigurationAssignment: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedDeviceMobileAppConfigurationAssignment: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.managedDeviceMobileAppConfigurationAssignment"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedDeviceMobileAppConfigurationAssignment: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &ManagedDeviceMobileAppConfigurationAssignment{}

func (s *ManagedDeviceMobileAppConfigurationAssignment) UnmarshalJSON(bytes []byte) error {
	type alias ManagedDeviceMobileAppConfigurationAssignment
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into ManagedDeviceMobileAppConfigurationAssignment: %+v", err)
	}

	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ManagedDeviceMobileAppConfigurationAssignment into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["target"]; ok {
		impl, err := UnmarshalDeviceAndAppManagementAssignmentTargetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Target' for 'ManagedDeviceMobileAppConfigurationAssignment': %+v", err)
		}
		s.Target = impl
	}
	return nil
}
