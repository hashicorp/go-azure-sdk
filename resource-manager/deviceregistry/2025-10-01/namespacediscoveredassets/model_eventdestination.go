package namespacediscoveredassets

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventDestination interface {
	EventDestination() BaseEventDestinationImpl
}

var _ EventDestination = BaseEventDestinationImpl{}

type BaseEventDestinationImpl struct {
	Target EventDestinationTarget `json:"target"`
}

func (s BaseEventDestinationImpl) EventDestination() BaseEventDestinationImpl {
	return s
}

var _ EventDestination = RawEventDestinationImpl{}

// RawEventDestinationImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawEventDestinationImpl struct {
	eventDestination BaseEventDestinationImpl
	Type             string
	Values           map[string]interface{}
}

func (s RawEventDestinationImpl) EventDestination() BaseEventDestinationImpl {
	return s.eventDestination
}

func (s RawEventDestinationImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalEventDestinationImplementation(input []byte) (EventDestination, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling EventDestination into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["target"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "Mqtt") {
		var out EventMqttDestination
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EventMqttDestination: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Storage") {
		var out EventStorageDestination
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EventStorageDestination: %+v", err)
		}
		return out, nil
	}

	var parent BaseEventDestinationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseEventDestinationImpl: %+v", err)
	}

	return RawEventDestinationImpl{
		eventDestination: parent,
		Type:             value,
		Values:           temp,
	}, nil

}
