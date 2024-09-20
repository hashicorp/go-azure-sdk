package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ BaseCollectionPaginationCountResponse = WindowsInformationProtectionCollectionResponse{}

type WindowsInformationProtectionCollectionResponse struct {
	Value *[]WindowsInformationProtection `json:"value,omitempty"`

	// Fields inherited from BaseCollectionPaginationCountResponse

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	ODataNextLink nullable.Type[string] `json:"@odata.nextLink,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s WindowsInformationProtectionCollectionResponse) BaseCollectionPaginationCountResponse() BaseBaseCollectionPaginationCountResponseImpl {
	return BaseBaseCollectionPaginationCountResponseImpl{
		ODataId:       s.ODataId,
		ODataNextLink: s.ODataNextLink,
		ODataType:     s.ODataType,
	}
}

var _ json.Marshaler = WindowsInformationProtectionCollectionResponse{}

func (s WindowsInformationProtectionCollectionResponse) MarshalJSON() ([]byte, error) {
	type wrapper WindowsInformationProtectionCollectionResponse
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsInformationProtectionCollectionResponse: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsInformationProtectionCollectionResponse: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.windowsInformationProtectionCollectionResponse"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsInformationProtectionCollectionResponse: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &WindowsInformationProtectionCollectionResponse{}

func (s *WindowsInformationProtectionCollectionResponse) UnmarshalJSON(bytes []byte) error {

	var decoded struct {
		Value         *[]WindowsInformationProtection `json:"value,omitempty"`
		ODataId       *string                         `json:"@odata.id,omitempty"`
		ODataNextLink nullable.Type[string]           `json:"@odata.nextLink,omitempty"`
		ODataType     *string                         `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ODataId = decoded.ODataId
	s.ODataNextLink = decoded.ODataNextLink
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling WindowsInformationProtectionCollectionResponse into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["value"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Value into list []json.RawMessage: %+v", err)
		}

		output := make([]WindowsInformationProtection, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalWindowsInformationProtectionImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Value' for 'WindowsInformationProtectionCollectionResponse': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Value = &output
	}

	return nil
}
