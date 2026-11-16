package entities

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SignalRecommendationConfiguration = PrometheusMetricsRecommendationConfiguration{}

type PrometheusMetricsRecommendationConfiguration struct {
	QueryText             string    `json:"queryText"`
	RequiredMetrics       *[]string `json:"requiredMetrics,omitempty"`
	RequiredScrapeTargets *[]string `json:"requiredScrapeTargets,omitempty"`
	TimeGrain             *string   `json:"timeGrain,omitempty"`

	// Fields inherited from SignalRecommendationConfiguration

	SignalKind SignalRecommendationKind `json:"signalKind"`
}

func (s PrometheusMetricsRecommendationConfiguration) SignalRecommendationConfiguration() BaseSignalRecommendationConfigurationImpl {
	return BaseSignalRecommendationConfigurationImpl{
		SignalKind: s.SignalKind,
	}
}

var _ json.Marshaler = PrometheusMetricsRecommendationConfiguration{}

func (s PrometheusMetricsRecommendationConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper PrometheusMetricsRecommendationConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PrometheusMetricsRecommendationConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PrometheusMetricsRecommendationConfiguration: %+v", err)
	}

	decoded["signalKind"] = "PrometheusMetricsQuery"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PrometheusMetricsRecommendationConfiguration: %+v", err)
	}

	return encoded, nil
}
