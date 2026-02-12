package iotcentrals

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventHubsV1DestinationAuth interface {
	EventHubsV1DestinationAuth() BaseEventHubsV1DestinationAuthImpl
}

var _ EventHubsV1DestinationAuth = BaseEventHubsV1DestinationAuthImpl{}

type BaseEventHubsV1DestinationAuthImpl struct {
	Type string `json:"type"`
}

func (s BaseEventHubsV1DestinationAuthImpl) EventHubsV1DestinationAuth() BaseEventHubsV1DestinationAuthImpl {
	return s
}

var _ EventHubsV1DestinationAuth = RawEventHubsV1DestinationAuthImpl{}

// RawEventHubsV1DestinationAuthImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawEventHubsV1DestinationAuthImpl struct {
	eventHubsV1DestinationAuth BaseEventHubsV1DestinationAuthImpl
	Type                       string
	Values                     map[string]interface{}
}

func (s RawEventHubsV1DestinationAuthImpl) EventHubsV1DestinationAuth() BaseEventHubsV1DestinationAuthImpl {
	return s.eventHubsV1DestinationAuth
}

func UnmarshalEventHubsV1DestinationAuthImplementation(input []byte) (EventHubsV1DestinationAuth, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling EventHubsV1DestinationAuth into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "connectionString") {
		var out EventHubsV1DestinationConnectionStringAuth
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EventHubsV1DestinationConnectionStringAuth: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "systemAssignedManagedIdentity") {
		var out EventHubsV1DestinationSystemAssignedManagedIdentityAuth
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EventHubsV1DestinationSystemAssignedManagedIdentityAuth: %+v", err)
		}
		return out, nil
	}

	var parent BaseEventHubsV1DestinationAuthImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseEventHubsV1DestinationAuthImpl: %+v", err)
	}

	return RawEventHubsV1DestinationAuthImpl{
		eventHubsV1DestinationAuth: parent,
		Type:                       value,
		Values:                     temp,
	}, nil

}
