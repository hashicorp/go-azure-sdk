package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ BaseCollectionPaginationCountResponse = TrustedCertificateAuthorityAsEntityBaseCollectionResponse{}

type TrustedCertificateAuthorityAsEntityBaseCollectionResponse struct {
	Value *[]TrustedCertificateAuthorityAsEntityBase `json:"value,omitempty"`

	// Fields inherited from BaseCollectionPaginationCountResponse

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	ODataNextLink nullable.Type[string] `json:"@odata.nextLink,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s TrustedCertificateAuthorityAsEntityBaseCollectionResponse) BaseCollectionPaginationCountResponse() BaseBaseCollectionPaginationCountResponseImpl {
	return BaseBaseCollectionPaginationCountResponseImpl{
		ODataId:       s.ODataId,
		ODataNextLink: s.ODataNextLink,
		ODataType:     s.ODataType,
	}
}

var _ json.Marshaler = TrustedCertificateAuthorityAsEntityBaseCollectionResponse{}

func (s TrustedCertificateAuthorityAsEntityBaseCollectionResponse) MarshalJSON() ([]byte, error) {
	type wrapper TrustedCertificateAuthorityAsEntityBaseCollectionResponse
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TrustedCertificateAuthorityAsEntityBaseCollectionResponse: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TrustedCertificateAuthorityAsEntityBaseCollectionResponse: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.trustedCertificateAuthorityAsEntityBaseCollectionResponse"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TrustedCertificateAuthorityAsEntityBaseCollectionResponse: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &TrustedCertificateAuthorityAsEntityBaseCollectionResponse{}

func (s *TrustedCertificateAuthorityAsEntityBaseCollectionResponse) UnmarshalJSON(bytes []byte) error {

	var decoded struct {
		Value         *[]TrustedCertificateAuthorityAsEntityBase `json:"value,omitempty"`
		ODataId       *string                                    `json:"@odata.id,omitempty"`
		ODataNextLink nullable.Type[string]                      `json:"@odata.nextLink,omitempty"`
		ODataType     *string                                    `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ODataId = decoded.ODataId
	s.ODataNextLink = decoded.ODataNextLink
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling TrustedCertificateAuthorityAsEntityBaseCollectionResponse into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["value"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Value into list []json.RawMessage: %+v", err)
		}

		output := make([]TrustedCertificateAuthorityAsEntityBase, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalTrustedCertificateAuthorityAsEntityBaseImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Value' for 'TrustedCertificateAuthorityAsEntityBaseCollectionResponse': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Value = &output
	}

	return nil
}
