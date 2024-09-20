package inventoryitems

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InventoryItemProperties interface {
	InventoryItemProperties() BaseInventoryItemPropertiesImpl
}

var _ InventoryItemProperties = BaseInventoryItemPropertiesImpl{}

type BaseInventoryItemPropertiesImpl struct {
	InventoryType     InventoryType      `json:"inventoryType"`
	ManagedResourceId *string            `json:"managedResourceId,omitempty"`
	MoName            *string            `json:"moName,omitempty"`
	MoRefId           *string            `json:"moRefId,omitempty"`
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty"`
}

func (s BaseInventoryItemPropertiesImpl) InventoryItemProperties() BaseInventoryItemPropertiesImpl {
	return s
}

var _ InventoryItemProperties = RawInventoryItemPropertiesImpl{}

// RawInventoryItemPropertiesImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawInventoryItemPropertiesImpl struct {
	inventoryItemProperties BaseInventoryItemPropertiesImpl
	Type                    string
	Values                  map[string]interface{}
}

func (s RawInventoryItemPropertiesImpl) InventoryItemProperties() BaseInventoryItemPropertiesImpl {
	return s.inventoryItemProperties
}

func UnmarshalInventoryItemPropertiesImplementation(input []byte) (InventoryItemProperties, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling InventoryItemProperties into map[string]interface: %+v", err)
	}

	value, ok := temp["inventoryType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "Cluster") {
		var out ClusterInventoryItem
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ClusterInventoryItem: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Datastore") {
		var out DatastoreInventoryItem
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DatastoreInventoryItem: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Host") {
		var out HostInventoryItem
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into HostInventoryItem: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "ResourcePool") {
		var out ResourcePoolInventoryItem
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ResourcePoolInventoryItem: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "VirtualMachine") {
		var out VirtualMachineInventoryItem
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VirtualMachineInventoryItem: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "VirtualMachineTemplate") {
		var out VirtualMachineTemplateInventoryItem
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VirtualMachineTemplateInventoryItem: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "VirtualNetwork") {
		var out VirtualNetworkInventoryItem
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VirtualNetworkInventoryItem: %+v", err)
		}
		return out, nil
	}

	var parent BaseInventoryItemPropertiesImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseInventoryItemPropertiesImpl: %+v", err)
	}

	return RawInventoryItemPropertiesImpl{
		inventoryItemProperties: parent,
		Type:                    value,
		Values:                  temp,
	}, nil

}
