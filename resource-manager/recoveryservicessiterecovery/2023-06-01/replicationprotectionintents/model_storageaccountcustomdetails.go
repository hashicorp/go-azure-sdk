package replicationprotectionintents

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StorageAccountCustomDetails interface {
	StorageAccountCustomDetails() BaseStorageAccountCustomDetailsImpl
}

var _ StorageAccountCustomDetails = BaseStorageAccountCustomDetailsImpl{}

type BaseStorageAccountCustomDetailsImpl struct {
	ResourceType string `json:"resourceType"`
}

func (s BaseStorageAccountCustomDetailsImpl) StorageAccountCustomDetails() BaseStorageAccountCustomDetailsImpl {
	return s
}

var _ StorageAccountCustomDetails = RawStorageAccountCustomDetailsImpl{}

// RawStorageAccountCustomDetailsImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawStorageAccountCustomDetailsImpl struct {
	storageAccountCustomDetails BaseStorageAccountCustomDetailsImpl
	Type                        string
	Values                      map[string]interface{}
}

func (s RawStorageAccountCustomDetailsImpl) StorageAccountCustomDetails() BaseStorageAccountCustomDetailsImpl {
	return s.storageAccountCustomDetails
}

func UnmarshalStorageAccountCustomDetailsImplementation(input []byte) (StorageAccountCustomDetails, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling StorageAccountCustomDetails into map[string]interface: %+v", err)
	}

	value, ok := temp["resourceType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "Existing") {
		var out ExistingStorageAccount
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExistingStorageAccount: %+v", err)
		}
		return out, nil
	}

	var parent BaseStorageAccountCustomDetailsImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseStorageAccountCustomDetailsImpl: %+v", err)
	}

	return RawStorageAccountCustomDetailsImpl{
		storageAccountCustomDetails: parent,
		Type:                        value,
		Values:                      temp,
	}, nil

}
