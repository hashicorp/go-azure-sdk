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

// RawStorageAccountCustomDetailsImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawStorageAccountCustomDetailsImpl struct {
	storageAccountCustomDetails BaseStorageAccountCustomDetailsImpl
	Type                        string
	Values                      map[string]interface{}
}

func (s RawStorageAccountCustomDetailsImpl) StorageAccountCustomDetails() BaseStorageAccountCustomDetailsImpl {
	return s.storageAccountCustomDetails
}

func (s RawStorageAccountCustomDetailsImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalStorageAccountCustomDetailsImplementation(input []byte) (StorageAccountCustomDetails, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling StorageAccountCustomDetails into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["resourceType"]; ok {
		value = fmt.Sprintf("%v", v)
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
