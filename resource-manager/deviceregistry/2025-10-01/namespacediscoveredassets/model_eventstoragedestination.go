package namespacediscoveredassets

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EventDestination = EventStorageDestination{}

type EventStorageDestination struct {
	Configuration StorageDestinationConfiguration `json:"configuration"`

	// Fields inherited from EventDestination

	Target EventDestinationTarget `json:"target"`
}

func (s EventStorageDestination) EventDestination() BaseEventDestinationImpl {
	return BaseEventDestinationImpl{
		Target: s.Target,
	}
}

var _ json.Marshaler = EventStorageDestination{}

func (s EventStorageDestination) MarshalJSON() ([]byte, error) {
	type wrapper EventStorageDestination
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EventStorageDestination: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EventStorageDestination: %+v", err)
	}

	decoded["target"] = "Storage"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EventStorageDestination: %+v", err)
	}

	return encoded, nil
}
