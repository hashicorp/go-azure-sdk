package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ BaseCollectionPaginationCountResponse = OverprovisionedServerlessFunctionFindingCollectionResponse{}

type OverprovisionedServerlessFunctionFindingCollectionResponse struct {
	Value *[]OverprovisionedServerlessFunctionFinding `json:"value,omitempty"`

	// Fields inherited from BaseCollectionPaginationCountResponse

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	ODataNextLink nullable.Type[string] `json:"@odata.nextLink,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s OverprovisionedServerlessFunctionFindingCollectionResponse) BaseCollectionPaginationCountResponse() BaseBaseCollectionPaginationCountResponseImpl {
	return BaseBaseCollectionPaginationCountResponseImpl{
		ODataId:       s.ODataId,
		ODataNextLink: s.ODataNextLink,
		ODataType:     s.ODataType,
	}
}

var _ json.Marshaler = OverprovisionedServerlessFunctionFindingCollectionResponse{}

func (s OverprovisionedServerlessFunctionFindingCollectionResponse) MarshalJSON() ([]byte, error) {
	type wrapper OverprovisionedServerlessFunctionFindingCollectionResponse
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling OverprovisionedServerlessFunctionFindingCollectionResponse: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling OverprovisionedServerlessFunctionFindingCollectionResponse: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.overprovisionedServerlessFunctionFindingCollectionResponse"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling OverprovisionedServerlessFunctionFindingCollectionResponse: %+v", err)
	}

	return encoded, nil
}
