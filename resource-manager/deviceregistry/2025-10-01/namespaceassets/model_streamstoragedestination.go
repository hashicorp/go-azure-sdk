package namespaceassets

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ StreamDestination = StreamStorageDestination{}

type StreamStorageDestination struct {
	Configuration StorageDestinationConfiguration `json:"configuration"`

	// Fields inherited from StreamDestination

	Target StreamDestinationTarget `json:"target"`
}

func (s StreamStorageDestination) StreamDestination() BaseStreamDestinationImpl {
	return BaseStreamDestinationImpl{
		Target: s.Target,
	}
}

var _ json.Marshaler = StreamStorageDestination{}

func (s StreamStorageDestination) MarshalJSON() ([]byte, error) {
	type wrapper StreamStorageDestination
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling StreamStorageDestination: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling StreamStorageDestination: %+v", err)
	}

	decoded["target"] = "Storage"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling StreamStorageDestination: %+v", err)
	}

	return encoded, nil
}
