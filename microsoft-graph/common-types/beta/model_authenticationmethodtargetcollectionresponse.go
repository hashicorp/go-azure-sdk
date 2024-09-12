package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ BaseCollectionPaginationCountResponse = AuthenticationMethodTargetCollectionResponse{}

type AuthenticationMethodTargetCollectionResponse struct {
	Value *[]AuthenticationMethodTarget `json:"value,omitempty"`

	// Fields inherited from BaseCollectionPaginationCountResponse

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	ODataNextLink nullable.Type[string] `json:"@odata.nextLink,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s AuthenticationMethodTargetCollectionResponse) BaseCollectionPaginationCountResponse() BaseBaseCollectionPaginationCountResponseImpl {
	return BaseBaseCollectionPaginationCountResponseImpl{
		ODataId:       s.ODataId,
		ODataNextLink: s.ODataNextLink,
		ODataType:     s.ODataType,
	}
}

var _ json.Marshaler = AuthenticationMethodTargetCollectionResponse{}

func (s AuthenticationMethodTargetCollectionResponse) MarshalJSON() ([]byte, error) {
	type wrapper AuthenticationMethodTargetCollectionResponse
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AuthenticationMethodTargetCollectionResponse: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AuthenticationMethodTargetCollectionResponse: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.authenticationMethodTargetCollectionResponse"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AuthenticationMethodTargetCollectionResponse: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &AuthenticationMethodTargetCollectionResponse{}

func (s *AuthenticationMethodTargetCollectionResponse) UnmarshalJSON(bytes []byte) error {
	type alias AuthenticationMethodTargetCollectionResponse
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into AuthenticationMethodTargetCollectionResponse: %+v", err)
	}

	s.ODataId = decoded.ODataId
	s.ODataNextLink = decoded.ODataNextLink
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AuthenticationMethodTargetCollectionResponse into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["value"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Value into list []json.RawMessage: %+v", err)
		}

		output := make([]AuthenticationMethodTarget, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalAuthenticationMethodTargetImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Value' for 'AuthenticationMethodTargetCollectionResponse': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Value = &output
	}
	return nil
}
