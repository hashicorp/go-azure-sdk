package namespaceassets

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EventDestination = EventMqttDestination{}

type EventMqttDestination struct {
	Configuration MqttDestinationConfiguration `json:"configuration"`

	// Fields inherited from EventDestination

	Target EventDestinationTarget `json:"target"`
}

func (s EventMqttDestination) EventDestination() BaseEventDestinationImpl {
	return BaseEventDestinationImpl{
		Target: s.Target,
	}
}

var _ json.Marshaler = EventMqttDestination{}

func (s EventMqttDestination) MarshalJSON() ([]byte, error) {
	type wrapper EventMqttDestination
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EventMqttDestination: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EventMqttDestination: %+v", err)
	}

	decoded["target"] = "Mqtt"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EventMqttDestination: %+v", err)
	}

	return encoded, nil
}
