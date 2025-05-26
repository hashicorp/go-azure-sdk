package schedule

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MonitoringFeatureFilterBase = AllFeatures{}

type AllFeatures struct {

	// Fields inherited from MonitoringFeatureFilterBase

	FilterType MonitoringFeatureFilterType `json:"filterType"`
}

func (s AllFeatures) MonitoringFeatureFilterBase() BaseMonitoringFeatureFilterBaseImpl {
	return BaseMonitoringFeatureFilterBaseImpl{
		FilterType: s.FilterType,
	}
}

var _ json.Marshaler = AllFeatures{}

func (s AllFeatures) MarshalJSON() ([]byte, error) {
	type wrapper AllFeatures
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AllFeatures: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AllFeatures: %+v", err)
	}

	decoded["filterType"] = "AllFeatures"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AllFeatures: %+v", err)
	}

	return encoded, nil
}
