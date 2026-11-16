package entities

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignalInstanceProperties interface {
	SignalInstanceProperties() BaseSignalInstancePropertiesImpl
}

var _ SignalInstanceProperties = BaseSignalInstancePropertiesImpl{}

type BaseSignalInstancePropertiesImpl struct {
	Name                 string        `json:"name"`
	SignalDefinitionName *string       `json:"signalDefinitionName,omitempty"`
	SignalKind           SignalKind    `json:"signalKind"`
	Status               *SignalStatus `json:"status,omitempty"`
}

func (s BaseSignalInstancePropertiesImpl) SignalInstanceProperties() BaseSignalInstancePropertiesImpl {
	return s
}

var _ SignalInstanceProperties = RawSignalInstancePropertiesImpl{}

// RawSignalInstancePropertiesImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawSignalInstancePropertiesImpl struct {
	signalInstanceProperties BaseSignalInstancePropertiesImpl
	Type                     string
	Values                   map[string]interface{}
}

func (s RawSignalInstancePropertiesImpl) SignalInstanceProperties() BaseSignalInstancePropertiesImpl {
	return s.signalInstanceProperties
}

func (s RawSignalInstancePropertiesImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalSignalInstancePropertiesImplementation(input []byte) (SignalInstanceProperties, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling SignalInstanceProperties into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["signalKind"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "AzureResourceMetric") {
		var out AzureResourceSignal
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureResourceSignal: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "External") {
		var out ExternalSignal
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExternalSignal: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "LogAnalyticsQuery") {
		var out LogAnalyticsSignal
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into LogAnalyticsSignal: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "PrometheusMetricsQuery") {
		var out PrometheusMetricsSignal
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrometheusMetricsSignal: %+v", err)
		}
		return out, nil
	}

	var parent BaseSignalInstancePropertiesImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseSignalInstancePropertiesImpl: %+v", err)
	}

	return RawSignalInstancePropertiesImpl{
		signalInstanceProperties: parent,
		Type:                     value,
		Values:                   temp,
	}, nil

}
