package namespacediscoveredassets

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DatasetDestination interface {
	DatasetDestination() BaseDatasetDestinationImpl
}

var _ DatasetDestination = BaseDatasetDestinationImpl{}

type BaseDatasetDestinationImpl struct {
	Target DatasetDestinationTarget `json:"target"`
}

func (s BaseDatasetDestinationImpl) DatasetDestination() BaseDatasetDestinationImpl {
	return s
}

var _ DatasetDestination = RawDatasetDestinationImpl{}

// RawDatasetDestinationImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDatasetDestinationImpl struct {
	datasetDestination BaseDatasetDestinationImpl
	Type               string
	Values             map[string]interface{}
}

func (s RawDatasetDestinationImpl) DatasetDestination() BaseDatasetDestinationImpl {
	return s.datasetDestination
}

func UnmarshalDatasetDestinationImplementation(input []byte) (DatasetDestination, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DatasetDestination into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["target"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "BrokerStateStore") {
		var out DatasetBrokerStateStoreDestination
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DatasetBrokerStateStoreDestination: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Mqtt") {
		var out DatasetMqttDestination
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DatasetMqttDestination: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Storage") {
		var out DatasetStorageDestination
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DatasetStorageDestination: %+v", err)
		}
		return out, nil
	}

	var parent BaseDatasetDestinationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDatasetDestinationImpl: %+v", err)
	}

	return RawDatasetDestinationImpl{
		datasetDestination: parent,
		Type:               value,
		Values:             temp,
	}, nil

}
