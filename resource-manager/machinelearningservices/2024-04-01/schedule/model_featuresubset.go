package schedule

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MonitoringFeatureFilterBase = FeatureSubset{}

type FeatureSubset struct {
	Features []string `json:"features"`

	// Fields inherited from MonitoringFeatureFilterBase
}

var _ json.Marshaler = FeatureSubset{}

func (s FeatureSubset) MarshalJSON() ([]byte, error) {
	type wrapper FeatureSubset
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling FeatureSubset: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling FeatureSubset: %+v", err)
	}
	decoded["filterType"] = "FeatureSubset"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling FeatureSubset: %+v", err)
	}

	return encoded, nil
}
