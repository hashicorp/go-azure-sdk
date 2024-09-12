package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ BaseCollectionPaginationCountResponse = AccessPackageAnswerChoiceCollectionResponse{}

type AccessPackageAnswerChoiceCollectionResponse struct {
	Value *[]AccessPackageAnswerChoice `json:"value,omitempty"`

	// Fields inherited from BaseCollectionPaginationCountResponse

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	ODataNextLink nullable.Type[string] `json:"@odata.nextLink,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s AccessPackageAnswerChoiceCollectionResponse) BaseCollectionPaginationCountResponse() BaseBaseCollectionPaginationCountResponseImpl {
	return BaseBaseCollectionPaginationCountResponseImpl{
		ODataId:       s.ODataId,
		ODataNextLink: s.ODataNextLink,
		ODataType:     s.ODataType,
	}
}

var _ json.Marshaler = AccessPackageAnswerChoiceCollectionResponse{}

func (s AccessPackageAnswerChoiceCollectionResponse) MarshalJSON() ([]byte, error) {
	type wrapper AccessPackageAnswerChoiceCollectionResponse
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AccessPackageAnswerChoiceCollectionResponse: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AccessPackageAnswerChoiceCollectionResponse: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.accessPackageAnswerChoiceCollectionResponse"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AccessPackageAnswerChoiceCollectionResponse: %+v", err)
	}

	return encoded, nil
}
