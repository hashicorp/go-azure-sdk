package applicationgroup

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ApplicationGroupPolicy = ThrottlingPolicy{}

type ThrottlingPolicy struct {
	MetricId           MetricId `json:"metricId"`
	RateLimitThreshold int64    `json:"rateLimitThreshold"`

	// Fields inherited from ApplicationGroupPolicy
	Name string `json:"name"`
}

var _ json.Marshaler = ThrottlingPolicy{}

func (s ThrottlingPolicy) MarshalJSON() ([]byte, error) {
	type wrapper ThrottlingPolicy
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ThrottlingPolicy: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ThrottlingPolicy: %+v", err)
	}
	decoded["type"] = "ThrottlingPolicy"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ThrottlingPolicy: %+v", err)
	}

	return encoded, nil
}
