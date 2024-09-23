package protectablecontainers

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProtectableContainer interface {
	ProtectableContainer() BaseProtectableContainerImpl
}

var _ ProtectableContainer = BaseProtectableContainerImpl{}

type BaseProtectableContainerImpl struct {
	BackupManagementType     *BackupManagementType    `json:"backupManagementType,omitempty"`
	ContainerId              *string                  `json:"containerId,omitempty"`
	FriendlyName             *string                  `json:"friendlyName,omitempty"`
	HealthStatus             *string                  `json:"healthStatus,omitempty"`
	ProtectableContainerType ProtectableContainerType `json:"protectableContainerType"`
}

func (s BaseProtectableContainerImpl) ProtectableContainer() BaseProtectableContainerImpl {
	return s
}

var _ ProtectableContainer = RawProtectableContainerImpl{}

// RawProtectableContainerImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawProtectableContainerImpl struct {
	protectableContainer BaseProtectableContainerImpl
	Type                 string
	Values               map[string]interface{}
}

func (s RawProtectableContainerImpl) ProtectableContainer() BaseProtectableContainerImpl {
	return s.protectableContainer
}

func UnmarshalProtectableContainerImplementation(input []byte) (ProtectableContainer, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ProtectableContainer into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["protectableContainerType"]; ok {
		value = fmt.Sprintf("%v", v)
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

	var parent BaseProtectableContainerImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseProtectableContainerImpl: %+v", err)
	}

	return RawProtectableContainerImpl{
		protectableContainer: parent,
		Type:                 value,
		Values:               temp,
	}, nil

}
