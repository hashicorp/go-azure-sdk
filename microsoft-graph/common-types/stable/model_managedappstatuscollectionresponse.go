package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ BaseCollectionPaginationCountResponse = ManagedAppStatusCollectionResponse{}

type ManagedAppStatusCollectionResponse struct {
	Value *[]ManagedAppStatus `json:"value,omitempty"`

	// Fields inherited from BaseCollectionPaginationCountResponse

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	ODataNextLink nullable.Type[string] `json:"@odata.nextLink,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s ManagedAppStatusCollectionResponse) BaseCollectionPaginationCountResponse() BaseBaseCollectionPaginationCountResponseImpl {
	return BaseBaseCollectionPaginationCountResponseImpl{
		ODataId:       s.ODataId,
		ODataNextLink: s.ODataNextLink,
		ODataType:     s.ODataType,
	}
}

var _ json.Marshaler = ManagedAppStatusCollectionResponse{}

func (s ManagedAppStatusCollectionResponse) MarshalJSON() ([]byte, error) {
	type wrapper ManagedAppStatusCollectionResponse
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedAppStatusCollectionResponse: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedAppStatusCollectionResponse: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.managedAppStatusCollectionResponse"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedAppStatusCollectionResponse: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &ManagedAppStatusCollectionResponse{}

func (s *ManagedAppStatusCollectionResponse) UnmarshalJSON(bytes []byte) error {
	type alias ManagedAppStatusCollectionResponse
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into ManagedAppStatusCollectionResponse: %+v", err)
	}

	s.ODataId = decoded.ODataId
	s.ODataNextLink = decoded.ODataNextLink
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ManagedAppStatusCollectionResponse into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["value"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Value into list []json.RawMessage: %+v", err)
		}

		output := make([]ManagedAppStatus, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalManagedAppStatusImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Value' for 'ManagedAppStatusCollectionResponse': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Value = &output
	}
	return nil
}
