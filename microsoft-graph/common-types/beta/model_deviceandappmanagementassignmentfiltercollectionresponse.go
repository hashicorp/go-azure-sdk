package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ BaseCollectionPaginationCountResponse = DeviceAndAppManagementAssignmentFilterCollectionResponse{}

type DeviceAndAppManagementAssignmentFilterCollectionResponse struct {
	Value *[]DeviceAndAppManagementAssignmentFilter `json:"value,omitempty"`

	// Fields inherited from BaseCollectionPaginationCountResponse

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	ODataNextLink nullable.Type[string] `json:"@odata.nextLink,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s DeviceAndAppManagementAssignmentFilterCollectionResponse) BaseCollectionPaginationCountResponse() BaseBaseCollectionPaginationCountResponseImpl {
	return BaseBaseCollectionPaginationCountResponseImpl{
		ODataId:       s.ODataId,
		ODataNextLink: s.ODataNextLink,
		ODataType:     s.ODataType,
	}
}

var _ json.Marshaler = DeviceAndAppManagementAssignmentFilterCollectionResponse{}

func (s DeviceAndAppManagementAssignmentFilterCollectionResponse) MarshalJSON() ([]byte, error) {
	type wrapper DeviceAndAppManagementAssignmentFilterCollectionResponse
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceAndAppManagementAssignmentFilterCollectionResponse: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceAndAppManagementAssignmentFilterCollectionResponse: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.deviceAndAppManagementAssignmentFilterCollectionResponse"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceAndAppManagementAssignmentFilterCollectionResponse: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &DeviceAndAppManagementAssignmentFilterCollectionResponse{}

func (s *DeviceAndAppManagementAssignmentFilterCollectionResponse) UnmarshalJSON(bytes []byte) error {
	type alias DeviceAndAppManagementAssignmentFilterCollectionResponse
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into DeviceAndAppManagementAssignmentFilterCollectionResponse: %+v", err)
	}

	s.ODataId = decoded.ODataId
	s.ODataNextLink = decoded.ODataNextLink
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling DeviceAndAppManagementAssignmentFilterCollectionResponse into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["value"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Value into list []json.RawMessage: %+v", err)
		}

		output := make([]DeviceAndAppManagementAssignmentFilter, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDeviceAndAppManagementAssignmentFilterImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Value' for 'DeviceAndAppManagementAssignmentFilterCollectionResponse': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Value = &output
	}
	return nil
}
