package iotcentrals

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ServiceBusQueueV1DestinationAuth = ServiceBusQueueV1DestinationSystemAssignedManagedIdentityAuth{}

type ServiceBusQueueV1DestinationSystemAssignedManagedIdentityAuth struct {
	HostName  string `json:"hostName"`
	QueueName string `json:"queueName"`

	// Fields inherited from ServiceBusQueueV1DestinationAuth

	Type string `json:"type"`
}

func (s ServiceBusQueueV1DestinationSystemAssignedManagedIdentityAuth) ServiceBusQueueV1DestinationAuth() BaseServiceBusQueueV1DestinationAuthImpl {
	return BaseServiceBusQueueV1DestinationAuthImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = ServiceBusQueueV1DestinationSystemAssignedManagedIdentityAuth{}

func (s ServiceBusQueueV1DestinationSystemAssignedManagedIdentityAuth) MarshalJSON() ([]byte, error) {
	type wrapper ServiceBusQueueV1DestinationSystemAssignedManagedIdentityAuth
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ServiceBusQueueV1DestinationSystemAssignedManagedIdentityAuth: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ServiceBusQueueV1DestinationSystemAssignedManagedIdentityAuth: %+v", err)
	}

	decoded["type"] = "systemAssignedManagedIdentity"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ServiceBusQueueV1DestinationSystemAssignedManagedIdentityAuth: %+v", err)
	}

	return encoded, nil
}
