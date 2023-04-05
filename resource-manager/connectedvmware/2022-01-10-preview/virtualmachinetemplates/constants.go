package virtualmachinetemplates

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiskMode string

const (
	DiskModeIndependentNonpersistent DiskMode = "independent_nonpersistent"
	DiskModeIndependentPersistent    DiskMode = "independent_persistent"
	DiskModePersistent               DiskMode = "persistent"
)

func PossibleValuesForDiskMode() []string {
	return []string{
		string(DiskModeIndependentNonpersistent),
		string(DiskModeIndependentPersistent),
		string(DiskModePersistent),
	}
}

func (s *DiskMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForDiskMode() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = DiskMode(decoded)
	return nil
}

type DiskType string

const (
	DiskTypeFlat        DiskType = "flat"
	DiskTypePmem        DiskType = "pmem"
	DiskTypeRawphysical DiskType = "rawphysical"
	DiskTypeRawvirtual  DiskType = "rawvirtual"
	DiskTypeSesparse    DiskType = "sesparse"
	DiskTypeSparse      DiskType = "sparse"
	DiskTypeUnknown     DiskType = "unknown"
)

func PossibleValuesForDiskType() []string {
	return []string{
		string(DiskTypeFlat),
		string(DiskTypePmem),
		string(DiskTypeRawphysical),
		string(DiskTypeRawvirtual),
		string(DiskTypeSesparse),
		string(DiskTypeSparse),
		string(DiskTypeUnknown),
	}
}

func (s *DiskType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForDiskType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = DiskType(decoded)
	return nil
}

type FirmwareType string

const (
	FirmwareTypeBios FirmwareType = "bios"
	FirmwareTypeEfi  FirmwareType = "efi"
)

func PossibleValuesForFirmwareType() []string {
	return []string{
		string(FirmwareTypeBios),
		string(FirmwareTypeEfi),
	}
}

func (s *FirmwareType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForFirmwareType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = FirmwareType(decoded)
	return nil
}

type IPAddressAllocationMethod string

const (
	IPAddressAllocationMethodDynamic   IPAddressAllocationMethod = "dynamic"
	IPAddressAllocationMethodLinklayer IPAddressAllocationMethod = "linklayer"
	IPAddressAllocationMethodOther     IPAddressAllocationMethod = "other"
	IPAddressAllocationMethodRandom    IPAddressAllocationMethod = "random"
	IPAddressAllocationMethodStatic    IPAddressAllocationMethod = "static"
	IPAddressAllocationMethodUnset     IPAddressAllocationMethod = "unset"
)

func PossibleValuesForIPAddressAllocationMethod() []string {
	return []string{
		string(IPAddressAllocationMethodDynamic),
		string(IPAddressAllocationMethodLinklayer),
		string(IPAddressAllocationMethodOther),
		string(IPAddressAllocationMethodRandom),
		string(IPAddressAllocationMethodStatic),
		string(IPAddressAllocationMethodUnset),
	}
}

func (s *IPAddressAllocationMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForIPAddressAllocationMethod() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = IPAddressAllocationMethod(decoded)
	return nil
}

type NICType string

const (
	NICTypeEOneThousand  NICType = "e1000"
	NICTypeEOneThousande NICType = "e1000e"
	NICTypePcnetThreeTwo NICType = "pcnet32"
	NICTypeVMxnet        NICType = "vmxnet"
	NICTypeVMxnetThree   NICType = "vmxnet3"
	NICTypeVMxnetTwo     NICType = "vmxnet2"
)

func PossibleValuesForNICType() []string {
	return []string{
		string(NICTypeEOneThousand),
		string(NICTypeEOneThousande),
		string(NICTypePcnetThreeTwo),
		string(NICTypeVMxnet),
		string(NICTypeVMxnetThree),
		string(NICTypeVMxnetTwo),
	}
}

func (s *NICType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForNICType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = NICType(decoded)
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

type PowerOnBootOption string

const (
	PowerOnBootOptionDisabled PowerOnBootOption = "disabled"
	PowerOnBootOptionEnabled  PowerOnBootOption = "enabled"
)

func PossibleValuesForPowerOnBootOption() []string {
	return []string{
		string(PowerOnBootOptionDisabled),
		string(PowerOnBootOptionEnabled),
	}
}

func (s *PowerOnBootOption) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForPowerOnBootOption() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = PowerOnBootOption(decoded)
	return nil
}
