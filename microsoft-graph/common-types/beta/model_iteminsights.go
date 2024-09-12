package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ OfficeGraphInsights = ItemInsights{}

type ItemInsights struct {

	// Fields inherited from OfficeGraphInsights

	// Access this property from the derived type itemInsights.
	Shared *[]SharedInsight `json:"shared,omitempty"`

	// Access this property from the derived type itemInsights.
	Trending *[]Trending `json:"trending,omitempty"`

	// Access this property from the derived type itemInsights.
	Used *[]UsedInsight `json:"used,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s ItemInsights) OfficeGraphInsights() BaseOfficeGraphInsightsImpl {
	return BaseOfficeGraphInsightsImpl{
		Shared:    s.Shared,
		Trending:  s.Trending,
		Used:      s.Used,
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s ItemInsights) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ItemInsights{}

func (s ItemInsights) MarshalJSON() ([]byte, error) {
	type wrapper ItemInsights
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ItemInsights: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ItemInsights: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.itemInsights"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ItemInsights: %+v", err)
	}

	return encoded, nil
}
