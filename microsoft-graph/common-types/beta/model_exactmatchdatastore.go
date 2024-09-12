package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ExactMatchDataStoreBase = ExactMatchDataStore{}

type ExactMatchDataStore struct {
	Sessions *[]ExactMatchSession `json:"sessions,omitempty"`

	// Fields inherited from ExactMatchDataStoreBase

	Columns                 *[]ExactDataMatchStoreColumn `json:"columns,omitempty"`
	DataLastUpdatedDateTime nullable.Type[string]        `json:"dataLastUpdatedDateTime,omitempty"`
	Description             nullable.Type[string]        `json:"description,omitempty"`
	DisplayName             nullable.Type[string]        `json:"displayName,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s ExactMatchDataStore) ExactMatchDataStoreBase() BaseExactMatchDataStoreBaseImpl {
	return BaseExactMatchDataStoreBaseImpl{
		Columns:                 s.Columns,
		DataLastUpdatedDateTime: s.DataLastUpdatedDateTime,
		Description:             s.Description,
		DisplayName:             s.DisplayName,
		Id:                      s.Id,
		ODataId:                 s.ODataId,
		ODataType:               s.ODataType,
	}
}

func (s ExactMatchDataStore) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ExactMatchDataStore{}

func (s ExactMatchDataStore) MarshalJSON() ([]byte, error) {
	type wrapper ExactMatchDataStore
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ExactMatchDataStore: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ExactMatchDataStore: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.exactMatchDataStore"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ExactMatchDataStore: %+v", err)
	}

	return encoded, nil
}
