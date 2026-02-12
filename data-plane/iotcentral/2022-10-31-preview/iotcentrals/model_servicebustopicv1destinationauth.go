package iotcentrals

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceBusTopicV1DestinationAuth interface {
	ServiceBusTopicV1DestinationAuth() BaseServiceBusTopicV1DestinationAuthImpl
}

var _ ServiceBusTopicV1DestinationAuth = BaseServiceBusTopicV1DestinationAuthImpl{}

type BaseServiceBusTopicV1DestinationAuthImpl struct {
	Type string `json:"type"`
}

func (s BaseServiceBusTopicV1DestinationAuthImpl) ServiceBusTopicV1DestinationAuth() BaseServiceBusTopicV1DestinationAuthImpl {
	return s
}

var _ ServiceBusTopicV1DestinationAuth = RawServiceBusTopicV1DestinationAuthImpl{}

// RawServiceBusTopicV1DestinationAuthImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawServiceBusTopicV1DestinationAuthImpl struct {
	serviceBusTopicV1DestinationAuth BaseServiceBusTopicV1DestinationAuthImpl
	Type                             string
	Values                           map[string]interface{}
}

func (s RawServiceBusTopicV1DestinationAuthImpl) ServiceBusTopicV1DestinationAuth() BaseServiceBusTopicV1DestinationAuthImpl {
	return s.serviceBusTopicV1DestinationAuth
}

func UnmarshalServiceBusTopicV1DestinationAuthImplementation(input []byte) (ServiceBusTopicV1DestinationAuth, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ServiceBusTopicV1DestinationAuth into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "connectionString") {
		var out ServiceBusTopicV1DestinationConnectionStringAuth
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ServiceBusTopicV1DestinationConnectionStringAuth: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "systemAssignedManagedIdentity") {
		var out ServiceBusTopicV1DestinationSystemAssignedManagedIdentityAuth
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ServiceBusTopicV1DestinationSystemAssignedManagedIdentityAuth: %+v", err)
		}
		return out, nil
	}

	var parent BaseServiceBusTopicV1DestinationAuthImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseServiceBusTopicV1DestinationAuthImpl: %+v", err)
	}

	return RawServiceBusTopicV1DestinationAuthImpl{
		serviceBusTopicV1DestinationAuth: parent,
		Type:                             value,
		Values:                           temp,
	}, nil

}
