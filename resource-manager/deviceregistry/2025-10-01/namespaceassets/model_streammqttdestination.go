package namespaceassets

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ StreamDestination = StreamMqttDestination{}

type StreamMqttDestination struct {
	Configuration MqttDestinationConfiguration `json:"configuration"`

	// Fields inherited from StreamDestination

	Target StreamDestinationTarget `json:"target"`
}

func (s StreamMqttDestination) StreamDestination() BaseStreamDestinationImpl {
	return BaseStreamDestinationImpl{
		Target: s.Target,
	}
}

var _ json.Marshaler = StreamMqttDestination{}

func (s StreamMqttDestination) MarshalJSON() ([]byte, error) {
	type wrapper StreamMqttDestination
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling StreamMqttDestination: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling StreamMqttDestination: %+v", err)
	}

	decoded["target"] = "Mqtt"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling StreamMqttDestination: %+v", err)
	}

	return encoded, nil
}
