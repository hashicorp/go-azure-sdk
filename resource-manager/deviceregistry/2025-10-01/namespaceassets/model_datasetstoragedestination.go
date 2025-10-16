package namespaceassets

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DatasetDestination = DatasetStorageDestination{}

type DatasetStorageDestination struct {
	Configuration StorageDestinationConfiguration `json:"configuration"`

	// Fields inherited from DatasetDestination

	Target DatasetDestinationTarget `json:"target"`
}

func (s DatasetStorageDestination) DatasetDestination() BaseDatasetDestinationImpl {
	return BaseDatasetDestinationImpl{
		Target: s.Target,
	}
}

var _ json.Marshaler = DatasetStorageDestination{}

func (s DatasetStorageDestination) MarshalJSON() ([]byte, error) {
	type wrapper DatasetStorageDestination
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DatasetStorageDestination: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DatasetStorageDestination: %+v", err)
	}

	decoded["target"] = "Storage"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DatasetStorageDestination: %+v", err)
	}

	return encoded, nil
}
