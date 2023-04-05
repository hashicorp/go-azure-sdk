package virtualmachines

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

type OsTypeUM string

const (
	OsTypeUMLinux   OsTypeUM = "Linux"
	OsTypeUMWindows OsTypeUM = "Windows"
)

func PossibleValuesForOsTypeUM() []string {
	return []string{
		string(OsTypeUMLinux),
		string(OsTypeUMWindows),
	}
}

func (s *OsTypeUM) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForOsTypeUM() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = OsTypeUM(decoded)
	return nil
}

type PatchOperationStartedBy string

const (
	PatchOperationStartedByPlatform PatchOperationStartedBy = "Platform"
	PatchOperationStartedByUser     PatchOperationStartedBy = "User"
)

func PossibleValuesForPatchOperationStartedBy() []string {
	return []string{
		string(PatchOperationStartedByPlatform),
		string(PatchOperationStartedByUser),
	}
}

func (s *PatchOperationStartedBy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForPatchOperationStartedBy() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = PatchOperationStartedBy(decoded)
	return nil
}

type PatchOperationStatus string

const (
	PatchOperationStatusCompletedWithWarnings PatchOperationStatus = "CompletedWithWarnings"
	PatchOperationStatusFailed                PatchOperationStatus = "Failed"
	PatchOperationStatusInProgress            PatchOperationStatus = "InProgress"
	PatchOperationStatusSucceeded             PatchOperationStatus = "Succeeded"
	PatchOperationStatusUnknown               PatchOperationStatus = "Unknown"
)

func PossibleValuesForPatchOperationStatus() []string {
	return []string{
		string(PatchOperationStatusCompletedWithWarnings),
		string(PatchOperationStatusFailed),
		string(PatchOperationStatusInProgress),
		string(PatchOperationStatusSucceeded),
		string(PatchOperationStatusUnknown),
	}
}

func (s *PatchOperationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForPatchOperationStatus() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = PatchOperationStatus(decoded)
	return nil
}

type PatchServiceUsed string

const (
	PatchServiceUsedAPT     PatchServiceUsed = "APT"
	PatchServiceUsedUnknown PatchServiceUsed = "Unknown"
	PatchServiceUsedWU      PatchServiceUsed = "WU"
	PatchServiceUsedWUWSUS  PatchServiceUsed = "WU_WSUS"
	PatchServiceUsedYUM     PatchServiceUsed = "YUM"
	PatchServiceUsedZypper  PatchServiceUsed = "Zypper"
)

func PossibleValuesForPatchServiceUsed() []string {
	return []string{
		string(PatchServiceUsedAPT),
		string(PatchServiceUsedUnknown),
		string(PatchServiceUsedWU),
		string(PatchServiceUsedWUWSUS),
		string(PatchServiceUsedYUM),
		string(PatchServiceUsedZypper),
	}
}

func (s *PatchServiceUsed) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForPatchServiceUsed() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = PatchServiceUsed(decoded)
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

type SCSIControllerType string

const (
	SCSIControllerTypeBuslogic    SCSIControllerType = "buslogic"
	SCSIControllerTypeLsilogic    SCSIControllerType = "lsilogic"
	SCSIControllerTypeLsilogicsas SCSIControllerType = "lsilogicsas"
	SCSIControllerTypePvscsi      SCSIControllerType = "pvscsi"
)

func PossibleValuesForSCSIControllerType() []string {
	return []string{
		string(SCSIControllerTypeBuslogic),
		string(SCSIControllerTypeLsilogic),
		string(SCSIControllerTypeLsilogicsas),
		string(SCSIControllerTypePvscsi),
	}
}

func (s *SCSIControllerType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForSCSIControllerType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = SCSIControllerType(decoded)
	return nil
}

type StatusTypes string

const (
	StatusTypesConnected    StatusTypes = "Connected"
	StatusTypesDisconnected StatusTypes = "Disconnected"
	StatusTypesError        StatusTypes = "Error"
)

func PossibleValuesForStatusTypes() []string {
	return []string{
		string(StatusTypesConnected),
		string(StatusTypesDisconnected),
		string(StatusTypesError),
	}
}

func (s *StatusTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForStatusTypes() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = StatusTypes(decoded)
	return nil
}

type VMGuestPatchClassificationLinux string

const (
	VMGuestPatchClassificationLinuxCritical VMGuestPatchClassificationLinux = "Critical"
	VMGuestPatchClassificationLinuxOther    VMGuestPatchClassificationLinux = "Other"
	VMGuestPatchClassificationLinuxSecurity VMGuestPatchClassificationLinux = "Security"
)

func PossibleValuesForVMGuestPatchClassificationLinux() []string {
	return []string{
		string(VMGuestPatchClassificationLinuxCritical),
		string(VMGuestPatchClassificationLinuxOther),
		string(VMGuestPatchClassificationLinuxSecurity),
	}
}

func (s *VMGuestPatchClassificationLinux) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForVMGuestPatchClassificationLinux() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = VMGuestPatchClassificationLinux(decoded)
	return nil
}

type VMGuestPatchClassificationWindows string

const (
	VMGuestPatchClassificationWindowsCritical     VMGuestPatchClassificationWindows = "Critical"
	VMGuestPatchClassificationWindowsDefinition   VMGuestPatchClassificationWindows = "Definition"
	VMGuestPatchClassificationWindowsFeaturePack  VMGuestPatchClassificationWindows = "FeaturePack"
	VMGuestPatchClassificationWindowsSecurity     VMGuestPatchClassificationWindows = "Security"
	VMGuestPatchClassificationWindowsServicePack  VMGuestPatchClassificationWindows = "ServicePack"
	VMGuestPatchClassificationWindowsTools        VMGuestPatchClassificationWindows = "Tools"
	VMGuestPatchClassificationWindowsUpdateRollUp VMGuestPatchClassificationWindows = "UpdateRollUp"
	VMGuestPatchClassificationWindowsUpdates      VMGuestPatchClassificationWindows = "Updates"
)

func PossibleValuesForVMGuestPatchClassificationWindows() []string {
	return []string{
		string(VMGuestPatchClassificationWindowsCritical),
		string(VMGuestPatchClassificationWindowsDefinition),
		string(VMGuestPatchClassificationWindowsFeaturePack),
		string(VMGuestPatchClassificationWindowsSecurity),
		string(VMGuestPatchClassificationWindowsServicePack),
		string(VMGuestPatchClassificationWindowsTools),
		string(VMGuestPatchClassificationWindowsUpdateRollUp),
		string(VMGuestPatchClassificationWindowsUpdates),
	}
}

func (s *VMGuestPatchClassificationWindows) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForVMGuestPatchClassificationWindows() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = VMGuestPatchClassificationWindows(decoded)
	return nil
}

type VMGuestPatchRebootSetting string

const (
	VMGuestPatchRebootSettingAlways     VMGuestPatchRebootSetting = "Always"
	VMGuestPatchRebootSettingIfRequired VMGuestPatchRebootSetting = "IfRequired"
	VMGuestPatchRebootSettingNever      VMGuestPatchRebootSetting = "Never"
)

func PossibleValuesForVMGuestPatchRebootSetting() []string {
	return []string{
		string(VMGuestPatchRebootSettingAlways),
		string(VMGuestPatchRebootSettingIfRequired),
		string(VMGuestPatchRebootSettingNever),
	}
}

func (s *VMGuestPatchRebootSetting) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForVMGuestPatchRebootSetting() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = VMGuestPatchRebootSetting(decoded)
	return nil
}

type VMGuestPatchRebootStatus string

const (
	VMGuestPatchRebootStatusCompleted VMGuestPatchRebootStatus = "Completed"
	VMGuestPatchRebootStatusFailed    VMGuestPatchRebootStatus = "Failed"
	VMGuestPatchRebootStatusNotNeeded VMGuestPatchRebootStatus = "NotNeeded"
	VMGuestPatchRebootStatusRequired  VMGuestPatchRebootStatus = "Required"
	VMGuestPatchRebootStatusStarted   VMGuestPatchRebootStatus = "Started"
	VMGuestPatchRebootStatusUnknown   VMGuestPatchRebootStatus = "Unknown"
)

func PossibleValuesForVMGuestPatchRebootStatus() []string {
	return []string{
		string(VMGuestPatchRebootStatusCompleted),
		string(VMGuestPatchRebootStatusFailed),
		string(VMGuestPatchRebootStatusNotNeeded),
		string(VMGuestPatchRebootStatusRequired),
		string(VMGuestPatchRebootStatusStarted),
		string(VMGuestPatchRebootStatusUnknown),
	}
}

func (s *VMGuestPatchRebootStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForVMGuestPatchRebootStatus() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = VMGuestPatchRebootStatus(decoded)
	return nil
}

type VirtualSCSISharing string

const (
	VirtualSCSISharingNoSharing       VirtualSCSISharing = "noSharing"
	VirtualSCSISharingPhysicalSharing VirtualSCSISharing = "physicalSharing"
	VirtualSCSISharingVirtualSharing  VirtualSCSISharing = "virtualSharing"
)

func PossibleValuesForVirtualSCSISharing() []string {
	return []string{
		string(VirtualSCSISharingNoSharing),
		string(VirtualSCSISharingPhysicalSharing),
		string(VirtualSCSISharingVirtualSharing),
	}
}

func (s *VirtualSCSISharing) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForVirtualSCSISharing() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = VirtualSCSISharing(decoded)
	return nil
}
