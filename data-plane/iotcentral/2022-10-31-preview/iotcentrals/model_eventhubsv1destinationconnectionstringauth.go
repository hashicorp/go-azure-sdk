package iotcentrals

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EventHubsV1DestinationAuth = EventHubsV1DestinationConnectionStringAuth{}

type EventHubsV1DestinationConnectionStringAuth struct {
	ConnectionString string `json:"connectionString"`

	// Fields inherited from EventHubsV1DestinationAuth

	Type string `json:"type"`
}

func (s EventHubsV1DestinationConnectionStringAuth) EventHubsV1DestinationAuth() BaseEventHubsV1DestinationAuthImpl {
	return BaseEventHubsV1DestinationAuthImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = EventHubsV1DestinationConnectionStringAuth{}

func (s EventHubsV1DestinationConnectionStringAuth) MarshalJSON() ([]byte, error) {
	type wrapper EventHubsV1DestinationConnectionStringAuth
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EventHubsV1DestinationConnectionStringAuth: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EventHubsV1DestinationConnectionStringAuth: %+v", err)
	}

	decoded["type"] = "connectionString"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EventHubsV1DestinationConnectionStringAuth: %+v", err)
	}

	return encoded, nil
}
