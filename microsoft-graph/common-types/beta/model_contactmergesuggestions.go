package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ContactMergeSuggestions{}

type ContactMergeSuggestions struct {
	// true if the duplicate contact merge suggestions feature is enabled for the user; false if the feature is disabled.
	// Default value is true.
	IsEnabled nullable.Type[bool] `json:"isEnabled,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s ContactMergeSuggestions) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ContactMergeSuggestions{}

func (s ContactMergeSuggestions) MarshalJSON() ([]byte, error) {
	type wrapper ContactMergeSuggestions
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ContactMergeSuggestions: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ContactMergeSuggestions: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.contactMergeSuggestions"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ContactMergeSuggestions: %+v", err)
	}

	return encoded, nil
}
