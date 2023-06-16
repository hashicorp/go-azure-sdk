package protectablecontainers

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProtectableContainer interface {
}

func unmarshalProtectableContainerImplementation(input []byte) (ProtectableContainer, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ProtectableContainer into map[string]interface: %+v", err)
	}

	value, ok := temp["protectableContainerType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "StorageContainer") {
		var out AzureStorageProtectableContainer
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureStorageProtectableContainer: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "VMAppContainer") {
		var out AzureVMAppContainerProtectableContainer
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureVMAppContainerProtectableContainer: %+v", err)
		}
		return out, nil
	}

	type RawProtectableContainerImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawProtectableContainerImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
