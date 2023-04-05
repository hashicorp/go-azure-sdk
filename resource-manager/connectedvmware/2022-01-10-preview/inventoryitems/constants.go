package inventoryitems

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InventoryType string

const (
	InventoryTypeCluster                InventoryType = "Cluster"
	InventoryTypeDatastore              InventoryType = "Datastore"
	InventoryTypeHost                   InventoryType = "Host"
	InventoryTypeResourcePool           InventoryType = "ResourcePool"
	InventoryTypeVirtualMachine         InventoryType = "VirtualMachine"
	InventoryTypeVirtualMachineTemplate InventoryType = "VirtualMachineTemplate"
	InventoryTypeVirtualNetwork         InventoryType = "VirtualNetwork"
)

func PossibleValuesForInventoryType() []string {
	return []string{
		string(InventoryTypeCluster),
		string(InventoryTypeDatastore),
		string(InventoryTypeHost),
		string(InventoryTypeResourcePool),
		string(InventoryTypeVirtualMachine),
		string(InventoryTypeVirtualMachineTemplate),
		string(InventoryTypeVirtualNetwork),
	}
}

func (s *InventoryType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForInventoryType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = InventoryType(decoded)
	return nil
}

type OsType string

const (
	OsTypeLinux   OsType = "Linux"
	OsTypeOther   OsType = "Other"
	OsTypeWindows OsType = "Windows"
)

func PossibleValuesForOsType() []string {
	return []string{
		string(OsTypeLinux),
		string(OsTypeOther),
		string(OsTypeWindows),
	}
}

func (s *OsType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForOsType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = OsType(decoded)
	return nil
}
