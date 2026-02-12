package iotcentrals

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceBusQueueV1DestinationAuth interface {
	ServiceBusQueueV1DestinationAuth() BaseServiceBusQueueV1DestinationAuthImpl
}

var _ ServiceBusQueueV1DestinationAuth = BaseServiceBusQueueV1DestinationAuthImpl{}

type BaseServiceBusQueueV1DestinationAuthImpl struct {
	Type string `json:"type"`
}

func (s BaseServiceBusQueueV1DestinationAuthImpl) ServiceBusQueueV1DestinationAuth() BaseServiceBusQueueV1DestinationAuthImpl {
	return s
}

var _ ServiceBusQueueV1DestinationAuth = RawServiceBusQueueV1DestinationAuthImpl{}

// RawServiceBusQueueV1DestinationAuthImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawServiceBusQueueV1DestinationAuthImpl struct {
	serviceBusQueueV1DestinationAuth BaseServiceBusQueueV1DestinationAuthImpl
	Type                             string
	Values                           map[string]interface{}
}

func (s RawServiceBusQueueV1DestinationAuthImpl) ServiceBusQueueV1DestinationAuth() BaseServiceBusQueueV1DestinationAuthImpl {
	return s.serviceBusQueueV1DestinationAuth
}

func UnmarshalServiceBusQueueV1DestinationAuthImplementation(input []byte) (ServiceBusQueueV1DestinationAuth, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ServiceBusQueueV1DestinationAuth into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "connectionString") {
		var out ServiceBusQueueV1DestinationConnectionStringAuth
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ServiceBusQueueV1DestinationConnectionStringAuth: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "systemAssignedManagedIdentity") {
		var out ServiceBusQueueV1DestinationSystemAssignedManagedIdentityAuth
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ServiceBusQueueV1DestinationSystemAssignedManagedIdentityAuth: %+v", err)
		}
		return out, nil
	}

	var parent BaseServiceBusQueueV1DestinationAuthImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseServiceBusQueueV1DestinationAuthImpl: %+v", err)
	}

	return RawServiceBusQueueV1DestinationAuthImpl{
		serviceBusQueueV1DestinationAuth: parent,
		Type:                             value,
		Values:                           temp,
	}, nil

}
