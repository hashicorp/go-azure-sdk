package rackskus

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BootstrapProtocol string

const (
	BootstrapProtocolPXE BootstrapProtocol = "PXE"
)

func PossibleValuesForBootstrapProtocol() []string {
	return []string{
		string(BootstrapProtocolPXE),
	}
}

func (s *BootstrapProtocol) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBootstrapProtocol(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBootstrapProtocol(input string) (*BootstrapProtocol, error) {
	vals := map[string]BootstrapProtocol{
		"pxe": BootstrapProtocolPXE,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BootstrapProtocol(input)
	return &out, nil
}

type DeviceConnectionType string

const (
	DeviceConnectionTypePCI DeviceConnectionType = "PCI"
)

func PossibleValuesForDeviceConnectionType() []string {
	return []string{
		string(DeviceConnectionTypePCI),
	}
}

func (s *DeviceConnectionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceConnectionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceConnectionType(input string) (*DeviceConnectionType, error) {
	vals := map[string]DeviceConnectionType{
		"pci": DeviceConnectionTypePCI,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceConnectionType(input)
	return &out, nil
}

type DiskType string

const (
	DiskTypeHDD DiskType = "HDD"
	DiskTypeSSD DiskType = "SSD"
)

func PossibleValuesForDiskType() []string {
	return []string{
		string(DiskTypeHDD),
		string(DiskTypeSSD),
	}
}

func (s *DiskType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDiskType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDiskType(input string) (*DiskType, error) {
	vals := map[string]DiskType{
		"hdd": DiskTypeHDD,
		"ssd": DiskTypeSSD,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DiskType(input)
	return &out, nil
}

type MachineSkuDiskConnectionType string

const (
	MachineSkuDiskConnectionTypePCIE MachineSkuDiskConnectionType = "PCIE"
	MachineSkuDiskConnectionTypeRAID MachineSkuDiskConnectionType = "RAID"
	MachineSkuDiskConnectionTypeSAS  MachineSkuDiskConnectionType = "SAS"
	MachineSkuDiskConnectionTypeSATA MachineSkuDiskConnectionType = "SATA"
)

func PossibleValuesForMachineSkuDiskConnectionType() []string {
	return []string{
		string(MachineSkuDiskConnectionTypePCIE),
		string(MachineSkuDiskConnectionTypeRAID),
		string(MachineSkuDiskConnectionTypeSAS),
		string(MachineSkuDiskConnectionTypeSATA),
	}
}

func (s *MachineSkuDiskConnectionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMachineSkuDiskConnectionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMachineSkuDiskConnectionType(input string) (*MachineSkuDiskConnectionType, error) {
	vals := map[string]MachineSkuDiskConnectionType{
		"pcie": MachineSkuDiskConnectionTypePCIE,
		"raid": MachineSkuDiskConnectionTypeRAID,
		"sas":  MachineSkuDiskConnectionTypeSAS,
		"sata": MachineSkuDiskConnectionTypeSATA,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MachineSkuDiskConnectionType(input)
	return &out, nil
}

type RackSkuProvisioningState string

const (
	RackSkuProvisioningStateCanceled  RackSkuProvisioningState = "Canceled"
	RackSkuProvisioningStateFailed    RackSkuProvisioningState = "Failed"
	RackSkuProvisioningStateSucceeded RackSkuProvisioningState = "Succeeded"
)

func PossibleValuesForRackSkuProvisioningState() []string {
	return []string{
		string(RackSkuProvisioningStateCanceled),
		string(RackSkuProvisioningStateFailed),
		string(RackSkuProvisioningStateSucceeded),
	}
}

func (s *RackSkuProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRackSkuProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRackSkuProvisioningState(input string) (*RackSkuProvisioningState, error) {
	vals := map[string]RackSkuProvisioningState{
		"canceled":  RackSkuProvisioningStateCanceled,
		"failed":    RackSkuProvisioningStateFailed,
		"succeeded": RackSkuProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RackSkuProvisioningState(input)
	return &out, nil
}

type RackSkuType string

const (
	RackSkuTypeAggregator RackSkuType = "Aggregator"
	RackSkuTypeCompute    RackSkuType = "Compute"
	RackSkuTypeSingle     RackSkuType = "Single"
)

func PossibleValuesForRackSkuType() []string {
	return []string{
		string(RackSkuTypeAggregator),
		string(RackSkuTypeCompute),
		string(RackSkuTypeSingle),
	}
}

func (s *RackSkuType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRackSkuType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRackSkuType(input string) (*RackSkuType, error) {
	vals := map[string]RackSkuType{
		"aggregator": RackSkuTypeAggregator,
		"compute":    RackSkuTypeCompute,
		"single":     RackSkuTypeSingle,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RackSkuType(input)
	return &out, nil
}
