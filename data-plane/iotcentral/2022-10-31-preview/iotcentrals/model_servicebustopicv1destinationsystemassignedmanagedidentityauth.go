package iotcentrals

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ServiceBusTopicV1DestinationAuth = ServiceBusTopicV1DestinationSystemAssignedManagedIdentityAuth{}

type ServiceBusTopicV1DestinationSystemAssignedManagedIdentityAuth struct {
	HostName  string `json:"hostName"`
	TopicName string `json:"topicName"`

	// Fields inherited from ServiceBusTopicV1DestinationAuth

	Type string `json:"type"`
}

func (s ServiceBusTopicV1DestinationSystemAssignedManagedIdentityAuth) ServiceBusTopicV1DestinationAuth() BaseServiceBusTopicV1DestinationAuthImpl {
	return BaseServiceBusTopicV1DestinationAuthImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = ServiceBusTopicV1DestinationSystemAssignedManagedIdentityAuth{}

func (s ServiceBusTopicV1DestinationSystemAssignedManagedIdentityAuth) MarshalJSON() ([]byte, error) {
	type wrapper ServiceBusTopicV1DestinationSystemAssignedManagedIdentityAuth
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ServiceBusTopicV1DestinationSystemAssignedManagedIdentityAuth: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ServiceBusTopicV1DestinationSystemAssignedManagedIdentityAuth: %+v", err)
	}

	decoded["type"] = "systemAssignedManagedIdentity"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ServiceBusTopicV1DestinationSystemAssignedManagedIdentityAuth: %+v", err)
	}

	return encoded, nil
}
