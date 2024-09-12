package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ BaseCollectionPaginationCountResponse = Fido2AuthenticationMethodCollectionResponse{}

type Fido2AuthenticationMethodCollectionResponse struct {
	Value *[]Fido2AuthenticationMethod `json:"value,omitempty"`

	// Fields inherited from BaseCollectionPaginationCountResponse

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	ODataNextLink nullable.Type[string] `json:"@odata.nextLink,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s Fido2AuthenticationMethodCollectionResponse) BaseCollectionPaginationCountResponse() BaseBaseCollectionPaginationCountResponseImpl {
	return BaseBaseCollectionPaginationCountResponseImpl{
		ODataId:       s.ODataId,
		ODataNextLink: s.ODataNextLink,
		ODataType:     s.ODataType,
	}
}

var _ json.Marshaler = Fido2AuthenticationMethodCollectionResponse{}

func (s Fido2AuthenticationMethodCollectionResponse) MarshalJSON() ([]byte, error) {
	type wrapper Fido2AuthenticationMethodCollectionResponse
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Fido2AuthenticationMethodCollectionResponse: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Fido2AuthenticationMethodCollectionResponse: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.fido2AuthenticationMethodCollectionResponse"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Fido2AuthenticationMethodCollectionResponse: %+v", err)
	}

	return encoded, nil
}
