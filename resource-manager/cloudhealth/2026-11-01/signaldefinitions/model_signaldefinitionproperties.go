package signaldefinitions

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignalDefinitionProperties interface {
	SignalDefinitionProperties() BaseSignalDefinitionPropertiesImpl
}

var _ SignalDefinitionProperties = BaseSignalDefinitionPropertiesImpl{}

type BaseSignalDefinitionPropertiesImpl struct {
	DataUnit          *string                       `json:"dataUnit,omitempty"`
	DisplayName       *string                       `json:"displayName,omitempty"`
	EvaluationRules   EvaluationRule                `json:"evaluationRules"`
	ProvisioningState *HealthModelProvisioningState `json:"provisioningState,omitempty"`
	RefreshInterval   *RefreshInterval              `json:"refreshInterval,omitempty"`
	SignalKind        SignalKind                    `json:"signalKind"`
	Tags              *map[string]string            `json:"tags,omitempty"`
}

func (s BaseSignalDefinitionPropertiesImpl) SignalDefinitionProperties() BaseSignalDefinitionPropertiesImpl {
	return s
}

var _ SignalDefinitionProperties = RawSignalDefinitionPropertiesImpl{}

// RawSignalDefinitionPropertiesImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawSignalDefinitionPropertiesImpl struct {
	signalDefinitionProperties BaseSignalDefinitionPropertiesImpl
	Type                       string
	Values                     map[string]interface{}
}

func (s RawSignalDefinitionPropertiesImpl) SignalDefinitionProperties() BaseSignalDefinitionPropertiesImpl {
	return s.signalDefinitionProperties
}

func (s RawSignalDefinitionPropertiesImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalSignalDefinitionPropertiesImplementation(input []byte) (SignalDefinitionProperties, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling SignalDefinitionProperties into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["signalKind"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "LogAnalyticsQuery") {
		var out LogAnalyticsQuerySignalDefinitionProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into LogAnalyticsQuerySignalDefinitionProperties: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "PrometheusMetricsQuery") {
		var out PrometheusMetricsSignalDefinitionProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrometheusMetricsSignalDefinitionProperties: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "AzureResourceMetric") {
		var out ResourceMetricSignalDefinitionProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ResourceMetricSignalDefinitionProperties: %+v", err)
		}
		return out, nil
	}

	var parent BaseSignalDefinitionPropertiesImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseSignalDefinitionPropertiesImpl: %+v", err)
	}

	return RawSignalDefinitionPropertiesImpl{
		signalDefinitionProperties: parent,
		Type:                       value,
		Values:                     temp,
	}, nil

}
