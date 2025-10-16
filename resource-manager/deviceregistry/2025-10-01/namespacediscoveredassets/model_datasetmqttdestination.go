package namespacediscoveredassets

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DatasetDestination = DatasetMqttDestination{}

type DatasetMqttDestination struct {
	Configuration MqttDestinationConfiguration `json:"configuration"`

	// Fields inherited from DatasetDestination

	Target DatasetDestinationTarget `json:"target"`
}

func (s DatasetMqttDestination) DatasetDestination() BaseDatasetDestinationImpl {
	return BaseDatasetDestinationImpl{
		Target: s.Target,
	}
}

var _ json.Marshaler = DatasetMqttDestination{}

func (s DatasetMqttDestination) MarshalJSON() ([]byte, error) {
	type wrapper DatasetMqttDestination
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DatasetMqttDestination: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DatasetMqttDestination: %+v", err)
	}

	decoded["target"] = "Mqtt"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DatasetMqttDestination: %+v", err)
	}

	return encoded, nil
}
