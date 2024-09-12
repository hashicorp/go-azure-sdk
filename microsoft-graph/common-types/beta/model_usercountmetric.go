package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UserCountMetric{}

type UserCountMetric struct {
	// The total number of users in the tenant over time.
	Count *int64 `json:"count,omitempty"`

	// The date of the insight.
	FactDate *string `json:"factDate,omitempty"`

	Language nullable.Type[string] `json:"language,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s UserCountMetric) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UserCountMetric{}

func (s UserCountMetric) MarshalJSON() ([]byte, error) {
	type wrapper UserCountMetric
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UserCountMetric: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UserCountMetric: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.userCountMetric"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UserCountMetric: %+v", err)
	}

	return encoded, nil
}
