package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ BaseCollectionPaginationCountResponse = ServiceProvisioningErrorCollectionResponse{}

type ServiceProvisioningErrorCollectionResponse struct {
	Value *[]ServiceProvisioningError `json:"value,omitempty"`

	// Fields inherited from BaseCollectionPaginationCountResponse

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	ODataNextLink nullable.Type[string] `json:"@odata.nextLink,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s ServiceProvisioningErrorCollectionResponse) BaseCollectionPaginationCountResponse() BaseBaseCollectionPaginationCountResponseImpl {
	return BaseBaseCollectionPaginationCountResponseImpl{
		ODataId:       s.ODataId,
		ODataNextLink: s.ODataNextLink,
		ODataType:     s.ODataType,
	}
}

var _ json.Marshaler = ServiceProvisioningErrorCollectionResponse{}

func (s ServiceProvisioningErrorCollectionResponse) MarshalJSON() ([]byte, error) {
	type wrapper ServiceProvisioningErrorCollectionResponse
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ServiceProvisioningErrorCollectionResponse: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ServiceProvisioningErrorCollectionResponse: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.serviceProvisioningErrorCollectionResponse"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ServiceProvisioningErrorCollectionResponse: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &ServiceProvisioningErrorCollectionResponse{}

func (s *ServiceProvisioningErrorCollectionResponse) UnmarshalJSON(bytes []byte) error {
	type alias ServiceProvisioningErrorCollectionResponse
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into ServiceProvisioningErrorCollectionResponse: %+v", err)
	}

	s.ODataId = decoded.ODataId
	s.ODataNextLink = decoded.ODataNextLink
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ServiceProvisioningErrorCollectionResponse into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["value"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Value into list []json.RawMessage: %+v", err)
		}

		output := make([]ServiceProvisioningError, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalServiceProvisioningErrorImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Value' for 'ServiceProvisioningErrorCollectionResponse': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Value = &output
	}
	return nil
}
