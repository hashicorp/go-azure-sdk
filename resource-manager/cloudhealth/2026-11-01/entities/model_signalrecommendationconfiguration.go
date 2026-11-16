package entities

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignalRecommendationConfiguration interface {
	SignalRecommendationConfiguration() BaseSignalRecommendationConfigurationImpl
}

var _ SignalRecommendationConfiguration = BaseSignalRecommendationConfigurationImpl{}

type BaseSignalRecommendationConfigurationImpl struct {
	SignalKind SignalRecommendationKind `json:"signalKind"`
}

func (s BaseSignalRecommendationConfigurationImpl) SignalRecommendationConfiguration() BaseSignalRecommendationConfigurationImpl {
	return s
}

var _ SignalRecommendationConfiguration = RawSignalRecommendationConfigurationImpl{}

// RawSignalRecommendationConfigurationImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawSignalRecommendationConfigurationImpl struct {
	signalRecommendationConfiguration BaseSignalRecommendationConfigurationImpl
	Type                              string
	Values                            map[string]interface{}
}

func (s RawSignalRecommendationConfigurationImpl) SignalRecommendationConfiguration() BaseSignalRecommendationConfigurationImpl {
	return s.signalRecommendationConfiguration
}

func (s RawSignalRecommendationConfigurationImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalSignalRecommendationConfigurationImplementation(input []byte) (SignalRecommendationConfiguration, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling SignalRecommendationConfiguration into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["signalKind"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "AzureResourceMetric") {
		var out AzureResourceMetricRecommendationConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureResourceMetricRecommendationConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "LogAnalyticsQuery") {
		var out LogAnalyticsQueryRecommendationConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into LogAnalyticsQueryRecommendationConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "PrometheusMetricsQuery") {
		var out PrometheusMetricsRecommendationConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrometheusMetricsRecommendationConfiguration: %+v", err)
		}
		return out, nil
	}

	var parent BaseSignalRecommendationConfigurationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseSignalRecommendationConfigurationImpl: %+v", err)
	}

	return RawSignalRecommendationConfigurationImpl{
		signalRecommendationConfiguration: parent,
		Type:                              value,
		Values:                            temp,
	}, nil

}
