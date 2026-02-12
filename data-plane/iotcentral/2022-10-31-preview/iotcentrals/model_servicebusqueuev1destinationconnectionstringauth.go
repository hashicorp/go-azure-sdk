package iotcentrals

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ServiceBusQueueV1DestinationAuth = ServiceBusQueueV1DestinationConnectionStringAuth{}

type ServiceBusQueueV1DestinationConnectionStringAuth struct {
	ConnectionString string `json:"connectionString"`

	// Fields inherited from ServiceBusQueueV1DestinationAuth

	Type string `json:"type"`
}

func (s ServiceBusQueueV1DestinationConnectionStringAuth) ServiceBusQueueV1DestinationAuth() BaseServiceBusQueueV1DestinationAuthImpl {
	return BaseServiceBusQueueV1DestinationAuthImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = ServiceBusQueueV1DestinationConnectionStringAuth{}

func (s ServiceBusQueueV1DestinationConnectionStringAuth) MarshalJSON() ([]byte, error) {
	type wrapper ServiceBusQueueV1DestinationConnectionStringAuth
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ServiceBusQueueV1DestinationConnectionStringAuth: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ServiceBusQueueV1DestinationConnectionStringAuth: %+v", err)
	}

	decoded["type"] = "connectionString"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ServiceBusQueueV1DestinationConnectionStringAuth: %+v", err)
	}

	return encoded, nil
}
