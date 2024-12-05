package schedule

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MonitoringFeatureFilterBase = TopNFeaturesByAttribution{}

type TopNFeaturesByAttribution struct {
	Top *int64 `json:"top,omitempty"`

	// Fields inherited from MonitoringFeatureFilterBase

	FilterType MonitoringFeatureFilterType `json:"filterType"`
}

func (s TopNFeaturesByAttribution) MonitoringFeatureFilterBase() BaseMonitoringFeatureFilterBaseImpl {
	return BaseMonitoringFeatureFilterBaseImpl{
		FilterType: s.FilterType,
	}
}

var _ json.Marshaler = TopNFeaturesByAttribution{}

func (s TopNFeaturesByAttribution) MarshalJSON() ([]byte, error) {
	type wrapper TopNFeaturesByAttribution
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TopNFeaturesByAttribution: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TopNFeaturesByAttribution: %+v", err)
	}

	decoded["filterType"] = "TopNByAttribution"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TopNFeaturesByAttribution: %+v", err)
	}

	return encoded, nil
}
