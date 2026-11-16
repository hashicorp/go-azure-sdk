package entities

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SignalRecommendationConfiguration = AzureResourceMetricRecommendationConfiguration{}

type AzureResourceMetricRecommendationConfiguration struct {
	AggregationType MetricAggregationType `json:"aggregationType"`
	DimensionFilter *string               `json:"dimensionFilter,omitempty"`
	MetricName      string                `json:"metricName"`
	MetricNamespace string                `json:"metricNamespace"`
	TimeGrain       string                `json:"timeGrain"`

	// Fields inherited from SignalRecommendationConfiguration

	SignalKind SignalRecommendationKind `json:"signalKind"`
}

func (s AzureResourceMetricRecommendationConfiguration) SignalRecommendationConfiguration() BaseSignalRecommendationConfigurationImpl {
	return BaseSignalRecommendationConfigurationImpl{
		SignalKind: s.SignalKind,
	}
}

var _ json.Marshaler = AzureResourceMetricRecommendationConfiguration{}

func (s AzureResourceMetricRecommendationConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper AzureResourceMetricRecommendationConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzureResourceMetricRecommendationConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureResourceMetricRecommendationConfiguration: %+v", err)
	}

	decoded["signalKind"] = "AzureResourceMetric"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzureResourceMetricRecommendationConfiguration: %+v", err)
	}

	return encoded, nil
}
