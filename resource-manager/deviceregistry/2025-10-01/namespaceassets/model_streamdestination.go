package namespaceassets

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StreamDestination interface {
	StreamDestination() BaseStreamDestinationImpl
}

var _ StreamDestination = BaseStreamDestinationImpl{}

type BaseStreamDestinationImpl struct {
	Target StreamDestinationTarget `json:"target"`
}

func (s BaseStreamDestinationImpl) StreamDestination() BaseStreamDestinationImpl {
	return s
}

var _ StreamDestination = RawStreamDestinationImpl{}

// RawStreamDestinationImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawStreamDestinationImpl struct {
	streamDestination BaseStreamDestinationImpl
	Type              string
	Values            map[string]interface{}
}

func (s RawStreamDestinationImpl) StreamDestination() BaseStreamDestinationImpl {
	return s.streamDestination
}

func (s RawStreamDestinationImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalStreamDestinationImplementation(input []byte) (StreamDestination, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling StreamDestination into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["target"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "Mqtt") {
		var out StreamMqttDestination
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into StreamMqttDestination: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Storage") {
		var out StreamStorageDestination
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into StreamStorageDestination: %+v", err)
		}
		return out, nil
	}

	var parent BaseStreamDestinationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseStreamDestinationImpl: %+v", err)
	}

	return RawStreamDestinationImpl{
		streamDestination: parent,
		Type:              value,
		Values:            temp,
	}, nil

}
