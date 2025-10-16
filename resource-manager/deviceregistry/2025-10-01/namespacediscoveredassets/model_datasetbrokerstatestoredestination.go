package namespacediscoveredassets

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DatasetDestination = DatasetBrokerStateStoreDestination{}

type DatasetBrokerStateStoreDestination struct {
	Configuration BrokerStateStoreDestinationConfiguration `json:"configuration"`

	// Fields inherited from DatasetDestination

	Target DatasetDestinationTarget `json:"target"`
}

func (s DatasetBrokerStateStoreDestination) DatasetDestination() BaseDatasetDestinationImpl {
	return BaseDatasetDestinationImpl{
		Target: s.Target,
	}
}

var _ json.Marshaler = DatasetBrokerStateStoreDestination{}

func (s DatasetBrokerStateStoreDestination) MarshalJSON() ([]byte, error) {
	type wrapper DatasetBrokerStateStoreDestination
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DatasetBrokerStateStoreDestination: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DatasetBrokerStateStoreDestination: %+v", err)
	}

	decoded["target"] = "BrokerStateStore"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DatasetBrokerStateStoreDestination: %+v", err)
	}

	return encoded, nil
}
