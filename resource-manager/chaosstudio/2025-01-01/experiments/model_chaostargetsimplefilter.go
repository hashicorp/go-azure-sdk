package experiments

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ChaosTargetFilter = ChaosTargetSimpleFilter{}

type ChaosTargetSimpleFilter struct {
	Parameters *ChaosTargetSimpleFilterParameters `json:"parameters,omitempty"`

	// Fields inherited from ChaosTargetFilter

	Type FilterType `json:"type"`
}

func (s ChaosTargetSimpleFilter) ChaosTargetFilter() BaseChaosTargetFilterImpl {
	return BaseChaosTargetFilterImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = ChaosTargetSimpleFilter{}

func (s ChaosTargetSimpleFilter) MarshalJSON() ([]byte, error) {
	type wrapper ChaosTargetSimpleFilter
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ChaosTargetSimpleFilter: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ChaosTargetSimpleFilter: %+v", err)
	}

	decoded["type"] = "Simple"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ChaosTargetSimpleFilter: %+v", err)
	}

	return encoded, nil
}
