package inventoryitems

import "strings"

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

func parseInventoryType(input string) (*InventoryType, error) {
	vals := map[string]InventoryType{
		"cluster":                InventoryTypeCluster,
		"datastore":              InventoryTypeDatastore,
		"host":                   InventoryTypeHost,
		"resourcepool":           InventoryTypeResourcePool,
		"virtualmachine":         InventoryTypeVirtualMachine,
		"virtualmachinetemplate": InventoryTypeVirtualMachineTemplate,
		"virtualnetwork":         InventoryTypeVirtualNetwork,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InventoryType(input)
	return &out, nil
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

func parseOsType(input string) (*OsType, error) {
	vals := map[string]OsType{
		"linux":   OsTypeLinux,
		"other":   OsTypeOther,
		"windows": OsTypeWindows,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OsType(input)
	return &out, nil
}
