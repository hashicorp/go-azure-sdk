package virtualmachines

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultGateway string

const (
	DefaultGatewayFalse DefaultGateway = "False"
	DefaultGatewayTrue  DefaultGateway = "True"
)

func PossibleValuesForDefaultGateway() []string {
	return []string{
		string(DefaultGatewayFalse),
		string(DefaultGatewayTrue),
	}
}

func (s *DefaultGateway) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDefaultGateway(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDefaultGateway(input string) (*DefaultGateway, error) {
	vals := map[string]DefaultGateway{
		"false": DefaultGatewayFalse,
		"true":  DefaultGatewayTrue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultGateway(input)
	return &out, nil
}

type ExtendedLocationType string

const (
	ExtendedLocationTypeCustomLocation ExtendedLocationType = "CustomLocation"
	ExtendedLocationTypeEdgeZone       ExtendedLocationType = "EdgeZone"
)

func PossibleValuesForExtendedLocationType() []string {
	return []string{
		string(ExtendedLocationTypeCustomLocation),
		string(ExtendedLocationTypeEdgeZone),
	}
}

func (s *ExtendedLocationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseExtendedLocationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseExtendedLocationType(input string) (*ExtendedLocationType, error) {
	vals := map[string]ExtendedLocationType{
		"customlocation": ExtendedLocationTypeCustomLocation,
		"edgezone":       ExtendedLocationTypeEdgeZone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExtendedLocationType(input)
	return &out, nil
}

type OsDiskCreateOption string

const (
	OsDiskCreateOptionEphemeral  OsDiskCreateOption = "Ephemeral"
	OsDiskCreateOptionPersistent OsDiskCreateOption = "Persistent"
)

func PossibleValuesForOsDiskCreateOption() []string {
	return []string{
		string(OsDiskCreateOptionEphemeral),
		string(OsDiskCreateOptionPersistent),
	}
}

func (s *OsDiskCreateOption) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOsDiskCreateOption(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOsDiskCreateOption(input string) (*OsDiskCreateOption, error) {
	vals := map[string]OsDiskCreateOption{
		"ephemeral":  OsDiskCreateOptionEphemeral,
		"persistent": OsDiskCreateOptionPersistent,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OsDiskCreateOption(input)
	return &out, nil
}

type OsDiskDeleteOption string

const (
	OsDiskDeleteOptionDelete OsDiskDeleteOption = "Delete"
)

func PossibleValuesForOsDiskDeleteOption() []string {
	return []string{
		string(OsDiskDeleteOptionDelete),
	}
}

func (s *OsDiskDeleteOption) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOsDiskDeleteOption(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOsDiskDeleteOption(input string) (*OsDiskDeleteOption, error) {
	vals := map[string]OsDiskDeleteOption{
		"delete": OsDiskDeleteOptionDelete,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OsDiskDeleteOption(input)
	return &out, nil
}

type RelayType string

const (
	RelayTypePlatform RelayType = "Platform"
	RelayTypePublic   RelayType = "Public"
)

func PossibleValuesForRelayType() []string {
	return []string{
		string(RelayTypePlatform),
		string(RelayTypePublic),
	}
}

func (s *RelayType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRelayType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRelayType(input string) (*RelayType, error) {
	vals := map[string]RelayType{
		"platform": RelayTypePlatform,
		"public":   RelayTypePublic,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RelayType(input)
	return &out, nil
}

type SkipShutdown string

const (
	SkipShutdownFalse SkipShutdown = "False"
	SkipShutdownTrue  SkipShutdown = "True"
)

func PossibleValuesForSkipShutdown() []string {
	return []string{
		string(SkipShutdownFalse),
		string(SkipShutdownTrue),
	}
}

func (s *SkipShutdown) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSkipShutdown(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSkipShutdown(input string) (*SkipShutdown, error) {
	vals := map[string]SkipShutdown{
		"false": SkipShutdownFalse,
		"true":  SkipShutdownTrue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SkipShutdown(input)
	return &out, nil
}

type VirtualMachineBootMethod string

const (
	VirtualMachineBootMethodBIOS VirtualMachineBootMethod = "BIOS"
	VirtualMachineBootMethodUEFI VirtualMachineBootMethod = "UEFI"
)

func PossibleValuesForVirtualMachineBootMethod() []string {
	return []string{
		string(VirtualMachineBootMethodBIOS),
		string(VirtualMachineBootMethodUEFI),
	}
}

func (s *VirtualMachineBootMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVirtualMachineBootMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVirtualMachineBootMethod(input string) (*VirtualMachineBootMethod, error) {
	vals := map[string]VirtualMachineBootMethod{
		"bios": VirtualMachineBootMethodBIOS,
		"uefi": VirtualMachineBootMethodUEFI,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualMachineBootMethod(input)
	return &out, nil
}

type VirtualMachineDetailedStatus string

const (
	VirtualMachineDetailedStatusAvailable    VirtualMachineDetailedStatus = "Available"
	VirtualMachineDetailedStatusError        VirtualMachineDetailedStatus = "Error"
	VirtualMachineDetailedStatusProvisioning VirtualMachineDetailedStatus = "Provisioning"
	VirtualMachineDetailedStatusRunning      VirtualMachineDetailedStatus = "Running"
	VirtualMachineDetailedStatusScheduling   VirtualMachineDetailedStatus = "Scheduling"
	VirtualMachineDetailedStatusStopped      VirtualMachineDetailedStatus = "Stopped"
	VirtualMachineDetailedStatusTerminating  VirtualMachineDetailedStatus = "Terminating"
	VirtualMachineDetailedStatusUnknown      VirtualMachineDetailedStatus = "Unknown"
)

func PossibleValuesForVirtualMachineDetailedStatus() []string {
	return []string{
		string(VirtualMachineDetailedStatusAvailable),
		string(VirtualMachineDetailedStatusError),
		string(VirtualMachineDetailedStatusProvisioning),
		string(VirtualMachineDetailedStatusRunning),
		string(VirtualMachineDetailedStatusScheduling),
		string(VirtualMachineDetailedStatusStopped),
		string(VirtualMachineDetailedStatusTerminating),
		string(VirtualMachineDetailedStatusUnknown),
	}
}

func (s *VirtualMachineDetailedStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVirtualMachineDetailedStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVirtualMachineDetailedStatus(input string) (*VirtualMachineDetailedStatus, error) {
	vals := map[string]VirtualMachineDetailedStatus{
		"available":    VirtualMachineDetailedStatusAvailable,
		"error":        VirtualMachineDetailedStatusError,
		"provisioning": VirtualMachineDetailedStatusProvisioning,
		"running":      VirtualMachineDetailedStatusRunning,
		"scheduling":   VirtualMachineDetailedStatusScheduling,
		"stopped":      VirtualMachineDetailedStatusStopped,
		"terminating":  VirtualMachineDetailedStatusTerminating,
		"unknown":      VirtualMachineDetailedStatusUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualMachineDetailedStatus(input)
	return &out, nil
}

type VirtualMachineDeviceModelType string

const (
	VirtualMachineDeviceModelTypeTOne   VirtualMachineDeviceModelType = "T1"
	VirtualMachineDeviceModelTypeTThree VirtualMachineDeviceModelType = "T3"
	VirtualMachineDeviceModelTypeTTwo   VirtualMachineDeviceModelType = "T2"
)

func PossibleValuesForVirtualMachineDeviceModelType() []string {
	return []string{
		string(VirtualMachineDeviceModelTypeTOne),
		string(VirtualMachineDeviceModelTypeTThree),
		string(VirtualMachineDeviceModelTypeTTwo),
	}
}

func (s *VirtualMachineDeviceModelType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVirtualMachineDeviceModelType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVirtualMachineDeviceModelType(input string) (*VirtualMachineDeviceModelType, error) {
	vals := map[string]VirtualMachineDeviceModelType{
		"t1": VirtualMachineDeviceModelTypeTOne,
		"t3": VirtualMachineDeviceModelTypeTThree,
		"t2": VirtualMachineDeviceModelTypeTTwo,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualMachineDeviceModelType(input)
	return &out, nil
}

type VirtualMachineIPAllocationMethod string

const (
	VirtualMachineIPAllocationMethodDisabled VirtualMachineIPAllocationMethod = "Disabled"
	VirtualMachineIPAllocationMethodDynamic  VirtualMachineIPAllocationMethod = "Dynamic"
	VirtualMachineIPAllocationMethodStatic   VirtualMachineIPAllocationMethod = "Static"
)

func PossibleValuesForVirtualMachineIPAllocationMethod() []string {
	return []string{
		string(VirtualMachineIPAllocationMethodDisabled),
		string(VirtualMachineIPAllocationMethodDynamic),
		string(VirtualMachineIPAllocationMethodStatic),
	}
}

func (s *VirtualMachineIPAllocationMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVirtualMachineIPAllocationMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVirtualMachineIPAllocationMethod(input string) (*VirtualMachineIPAllocationMethod, error) {
	vals := map[string]VirtualMachineIPAllocationMethod{
		"disabled": VirtualMachineIPAllocationMethodDisabled,
		"dynamic":  VirtualMachineIPAllocationMethodDynamic,
		"static":   VirtualMachineIPAllocationMethodStatic,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualMachineIPAllocationMethod(input)
	return &out, nil
}

type VirtualMachineIsolateEmulatorThread string

const (
	VirtualMachineIsolateEmulatorThreadFalse VirtualMachineIsolateEmulatorThread = "False"
	VirtualMachineIsolateEmulatorThreadTrue  VirtualMachineIsolateEmulatorThread = "True"
)

func PossibleValuesForVirtualMachineIsolateEmulatorThread() []string {
	return []string{
		string(VirtualMachineIsolateEmulatorThreadFalse),
		string(VirtualMachineIsolateEmulatorThreadTrue),
	}
}

func (s *VirtualMachineIsolateEmulatorThread) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVirtualMachineIsolateEmulatorThread(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVirtualMachineIsolateEmulatorThread(input string) (*VirtualMachineIsolateEmulatorThread, error) {
	vals := map[string]VirtualMachineIsolateEmulatorThread{
		"false": VirtualMachineIsolateEmulatorThreadFalse,
		"true":  VirtualMachineIsolateEmulatorThreadTrue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualMachineIsolateEmulatorThread(input)
	return &out, nil
}

type VirtualMachinePlacementHintPodAffinityScope string

const (
	VirtualMachinePlacementHintPodAffinityScopeMachine VirtualMachinePlacementHintPodAffinityScope = "Machine"
	VirtualMachinePlacementHintPodAffinityScopeRack    VirtualMachinePlacementHintPodAffinityScope = "Rack"
)

func PossibleValuesForVirtualMachinePlacementHintPodAffinityScope() []string {
	return []string{
		string(VirtualMachinePlacementHintPodAffinityScopeMachine),
		string(VirtualMachinePlacementHintPodAffinityScopeRack),
	}
}

func (s *VirtualMachinePlacementHintPodAffinityScope) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVirtualMachinePlacementHintPodAffinityScope(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVirtualMachinePlacementHintPodAffinityScope(input string) (*VirtualMachinePlacementHintPodAffinityScope, error) {
	vals := map[string]VirtualMachinePlacementHintPodAffinityScope{
		"machine": VirtualMachinePlacementHintPodAffinityScopeMachine,
		"rack":    VirtualMachinePlacementHintPodAffinityScopeRack,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualMachinePlacementHintPodAffinityScope(input)
	return &out, nil
}

type VirtualMachinePlacementHintType string

const (
	VirtualMachinePlacementHintTypeAffinity     VirtualMachinePlacementHintType = "Affinity"
	VirtualMachinePlacementHintTypeAntiAffinity VirtualMachinePlacementHintType = "AntiAffinity"
)

func PossibleValuesForVirtualMachinePlacementHintType() []string {
	return []string{
		string(VirtualMachinePlacementHintTypeAffinity),
		string(VirtualMachinePlacementHintTypeAntiAffinity),
	}
}

func (s *VirtualMachinePlacementHintType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVirtualMachinePlacementHintType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVirtualMachinePlacementHintType(input string) (*VirtualMachinePlacementHintType, error) {
	vals := map[string]VirtualMachinePlacementHintType{
		"affinity":     VirtualMachinePlacementHintTypeAffinity,
		"antiaffinity": VirtualMachinePlacementHintTypeAntiAffinity,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualMachinePlacementHintType(input)
	return &out, nil
}

type VirtualMachinePowerState string

const (
	VirtualMachinePowerStateOff     VirtualMachinePowerState = "Off"
	VirtualMachinePowerStateOn      VirtualMachinePowerState = "On"
	VirtualMachinePowerStateUnknown VirtualMachinePowerState = "Unknown"
)

func PossibleValuesForVirtualMachinePowerState() []string {
	return []string{
		string(VirtualMachinePowerStateOff),
		string(VirtualMachinePowerStateOn),
		string(VirtualMachinePowerStateUnknown),
	}
}

func (s *VirtualMachinePowerState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVirtualMachinePowerState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVirtualMachinePowerState(input string) (*VirtualMachinePowerState, error) {
	vals := map[string]VirtualMachinePowerState{
		"off":     VirtualMachinePowerStateOff,
		"on":      VirtualMachinePowerStateOn,
		"unknown": VirtualMachinePowerStateUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualMachinePowerState(input)
	return &out, nil
}

type VirtualMachineProvisioningState string

const (
	VirtualMachineProvisioningStateAccepted     VirtualMachineProvisioningState = "Accepted"
	VirtualMachineProvisioningStateCanceled     VirtualMachineProvisioningState = "Canceled"
	VirtualMachineProvisioningStateFailed       VirtualMachineProvisioningState = "Failed"
	VirtualMachineProvisioningStateProvisioning VirtualMachineProvisioningState = "Provisioning"
	VirtualMachineProvisioningStateSucceeded    VirtualMachineProvisioningState = "Succeeded"
)

func PossibleValuesForVirtualMachineProvisioningState() []string {
	return []string{
		string(VirtualMachineProvisioningStateAccepted),
		string(VirtualMachineProvisioningStateCanceled),
		string(VirtualMachineProvisioningStateFailed),
		string(VirtualMachineProvisioningStateProvisioning),
		string(VirtualMachineProvisioningStateSucceeded),
	}
}

func (s *VirtualMachineProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVirtualMachineProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVirtualMachineProvisioningState(input string) (*VirtualMachineProvisioningState, error) {
	vals := map[string]VirtualMachineProvisioningState{
		"accepted":     VirtualMachineProvisioningStateAccepted,
		"canceled":     VirtualMachineProvisioningStateCanceled,
		"failed":       VirtualMachineProvisioningStateFailed,
		"provisioning": VirtualMachineProvisioningStateProvisioning,
		"succeeded":    VirtualMachineProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualMachineProvisioningState(input)
	return &out, nil
}

type VirtualMachineSchedulingExecution string

const (
	VirtualMachineSchedulingExecutionHard VirtualMachineSchedulingExecution = "Hard"
	VirtualMachineSchedulingExecutionSoft VirtualMachineSchedulingExecution = "Soft"
)

func PossibleValuesForVirtualMachineSchedulingExecution() []string {
	return []string{
		string(VirtualMachineSchedulingExecutionHard),
		string(VirtualMachineSchedulingExecutionSoft),
	}
}

func (s *VirtualMachineSchedulingExecution) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVirtualMachineSchedulingExecution(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVirtualMachineSchedulingExecution(input string) (*VirtualMachineSchedulingExecution, error) {
	vals := map[string]VirtualMachineSchedulingExecution{
		"hard": VirtualMachineSchedulingExecutionHard,
		"soft": VirtualMachineSchedulingExecutionSoft,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualMachineSchedulingExecution(input)
	return &out, nil
}

type VirtualMachineVirtioInterfaceType string

const (
	VirtualMachineVirtioInterfaceTypeModern       VirtualMachineVirtioInterfaceType = "Modern"
	VirtualMachineVirtioInterfaceTypeTransitional VirtualMachineVirtioInterfaceType = "Transitional"
)

func PossibleValuesForVirtualMachineVirtioInterfaceType() []string {
	return []string{
		string(VirtualMachineVirtioInterfaceTypeModern),
		string(VirtualMachineVirtioInterfaceTypeTransitional),
	}
}

func (s *VirtualMachineVirtioInterfaceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVirtualMachineVirtioInterfaceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVirtualMachineVirtioInterfaceType(input string) (*VirtualMachineVirtioInterfaceType, error) {
	vals := map[string]VirtualMachineVirtioInterfaceType{
		"modern":       VirtualMachineVirtioInterfaceTypeModern,
		"transitional": VirtualMachineVirtioInterfaceTypeTransitional,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualMachineVirtioInterfaceType(input)
	return &out, nil
}
