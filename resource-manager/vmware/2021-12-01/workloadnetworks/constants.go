package workloadnetworks

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DhcpTypeEnum string

const (
	DhcpTypeEnumRELAY  DhcpTypeEnum = "RELAY"
	DhcpTypeEnumSERVER DhcpTypeEnum = "SERVER"
)

func PossibleValuesForDhcpTypeEnum() []string {
	return []string{
		string(DhcpTypeEnumRELAY),
		string(DhcpTypeEnumSERVER),
	}
}

func parseDhcpTypeEnum(input string) (*DhcpTypeEnum, error) {
	vals := map[string]DhcpTypeEnum{
		"relay":  DhcpTypeEnumRELAY,
		"server": DhcpTypeEnumSERVER,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DhcpTypeEnum(input)
	return &out, nil
}

type DnsServiceLogLevelEnum string

const (
	DnsServiceLogLevelEnumDEBUG   DnsServiceLogLevelEnum = "DEBUG"
	DnsServiceLogLevelEnumERROR   DnsServiceLogLevelEnum = "ERROR"
	DnsServiceLogLevelEnumFATAL   DnsServiceLogLevelEnum = "FATAL"
	DnsServiceLogLevelEnumINFO    DnsServiceLogLevelEnum = "INFO"
	DnsServiceLogLevelEnumWARNING DnsServiceLogLevelEnum = "WARNING"
)

func PossibleValuesForDnsServiceLogLevelEnum() []string {
	return []string{
		string(DnsServiceLogLevelEnumDEBUG),
		string(DnsServiceLogLevelEnumERROR),
		string(DnsServiceLogLevelEnumFATAL),
		string(DnsServiceLogLevelEnumINFO),
		string(DnsServiceLogLevelEnumWARNING),
	}
}

func parseDnsServiceLogLevelEnum(input string) (*DnsServiceLogLevelEnum, error) {
	vals := map[string]DnsServiceLogLevelEnum{
		"debug":   DnsServiceLogLevelEnumDEBUG,
		"error":   DnsServiceLogLevelEnumERROR,
		"fatal":   DnsServiceLogLevelEnumFATAL,
		"info":    DnsServiceLogLevelEnumINFO,
		"warning": DnsServiceLogLevelEnumWARNING,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DnsServiceLogLevelEnum(input)
	return &out, nil
}

type DnsServiceStatusEnum string

const (
	DnsServiceStatusEnumFAILURE DnsServiceStatusEnum = "FAILURE"
	DnsServiceStatusEnumSUCCESS DnsServiceStatusEnum = "SUCCESS"
)

func PossibleValuesForDnsServiceStatusEnum() []string {
	return []string{
		string(DnsServiceStatusEnumFAILURE),
		string(DnsServiceStatusEnumSUCCESS),
	}
}

func parseDnsServiceStatusEnum(input string) (*DnsServiceStatusEnum, error) {
	vals := map[string]DnsServiceStatusEnum{
		"failure": DnsServiceStatusEnumFAILURE,
		"success": DnsServiceStatusEnumSUCCESS,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DnsServiceStatusEnum(input)
	return &out, nil
}

type PortMirroringDirectionEnum string

const (
	PortMirroringDirectionEnumBIDIRECTIONAL PortMirroringDirectionEnum = "BIDIRECTIONAL"
	PortMirroringDirectionEnumEGRESS        PortMirroringDirectionEnum = "EGRESS"
	PortMirroringDirectionEnumINGRESS       PortMirroringDirectionEnum = "INGRESS"
)

func PossibleValuesForPortMirroringDirectionEnum() []string {
	return []string{
		string(PortMirroringDirectionEnumBIDIRECTIONAL),
		string(PortMirroringDirectionEnumEGRESS),
		string(PortMirroringDirectionEnumINGRESS),
	}
}

func parsePortMirroringDirectionEnum(input string) (*PortMirroringDirectionEnum, error) {
	vals := map[string]PortMirroringDirectionEnum{
		"bidirectional": PortMirroringDirectionEnumBIDIRECTIONAL,
		"egress":        PortMirroringDirectionEnumEGRESS,
		"ingress":       PortMirroringDirectionEnumINGRESS,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PortMirroringDirectionEnum(input)
	return &out, nil
}

type PortMirroringStatusEnum string

const (
	PortMirroringStatusEnumFAILURE PortMirroringStatusEnum = "FAILURE"
	PortMirroringStatusEnumSUCCESS PortMirroringStatusEnum = "SUCCESS"
)

func PossibleValuesForPortMirroringStatusEnum() []string {
	return []string{
		string(PortMirroringStatusEnumFAILURE),
		string(PortMirroringStatusEnumSUCCESS),
	}
}

func parsePortMirroringStatusEnum(input string) (*PortMirroringStatusEnum, error) {
	vals := map[string]PortMirroringStatusEnum{
		"failure": PortMirroringStatusEnumFAILURE,
		"success": PortMirroringStatusEnumSUCCESS,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PortMirroringStatusEnum(input)
	return &out, nil
}

type SegmentStatusEnum string

const (
	SegmentStatusEnumFAILURE SegmentStatusEnum = "FAILURE"
	SegmentStatusEnumSUCCESS SegmentStatusEnum = "SUCCESS"
)

func PossibleValuesForSegmentStatusEnum() []string {
	return []string{
		string(SegmentStatusEnumFAILURE),
		string(SegmentStatusEnumSUCCESS),
	}
}

func parseSegmentStatusEnum(input string) (*SegmentStatusEnum, error) {
	vals := map[string]SegmentStatusEnum{
		"failure": SegmentStatusEnumFAILURE,
		"success": SegmentStatusEnumSUCCESS,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SegmentStatusEnum(input)
	return &out, nil
}

type VMGroupStatusEnum string

const (
	VMGroupStatusEnumFAILURE VMGroupStatusEnum = "FAILURE"
	VMGroupStatusEnumSUCCESS VMGroupStatusEnum = "SUCCESS"
)

func PossibleValuesForVMGroupStatusEnum() []string {
	return []string{
		string(VMGroupStatusEnumFAILURE),
		string(VMGroupStatusEnumSUCCESS),
	}
}

func parseVMGroupStatusEnum(input string) (*VMGroupStatusEnum, error) {
	vals := map[string]VMGroupStatusEnum{
		"failure": VMGroupStatusEnumFAILURE,
		"success": VMGroupStatusEnumSUCCESS,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VMGroupStatusEnum(input)
	return &out, nil
}

type VMTypeEnum string

const (
	VMTypeEnumEDGE    VMTypeEnum = "EDGE"
	VMTypeEnumREGULAR VMTypeEnum = "REGULAR"
	VMTypeEnumSERVICE VMTypeEnum = "SERVICE"
)

func PossibleValuesForVMTypeEnum() []string {
	return []string{
		string(VMTypeEnumEDGE),
		string(VMTypeEnumREGULAR),
		string(VMTypeEnumSERVICE),
	}
}

func parseVMTypeEnum(input string) (*VMTypeEnum, error) {
	vals := map[string]VMTypeEnum{
		"edge":    VMTypeEnumEDGE,
		"regular": VMTypeEnumREGULAR,
		"service": VMTypeEnumSERVICE,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VMTypeEnum(input)
	return &out, nil
}

type WorkloadNetworkDhcpProvisioningState string

const (
	WorkloadNetworkDhcpProvisioningStateBuilding  WorkloadNetworkDhcpProvisioningState = "Building"
	WorkloadNetworkDhcpProvisioningStateDeleting  WorkloadNetworkDhcpProvisioningState = "Deleting"
	WorkloadNetworkDhcpProvisioningStateFailed    WorkloadNetworkDhcpProvisioningState = "Failed"
	WorkloadNetworkDhcpProvisioningStateSucceeded WorkloadNetworkDhcpProvisioningState = "Succeeded"
	WorkloadNetworkDhcpProvisioningStateUpdating  WorkloadNetworkDhcpProvisioningState = "Updating"
)

func PossibleValuesForWorkloadNetworkDhcpProvisioningState() []string {
	return []string{
		string(WorkloadNetworkDhcpProvisioningStateBuilding),
		string(WorkloadNetworkDhcpProvisioningStateDeleting),
		string(WorkloadNetworkDhcpProvisioningStateFailed),
		string(WorkloadNetworkDhcpProvisioningStateSucceeded),
		string(WorkloadNetworkDhcpProvisioningStateUpdating),
	}
}

func parseWorkloadNetworkDhcpProvisioningState(input string) (*WorkloadNetworkDhcpProvisioningState, error) {
	vals := map[string]WorkloadNetworkDhcpProvisioningState{
		"building":  WorkloadNetworkDhcpProvisioningStateBuilding,
		"deleting":  WorkloadNetworkDhcpProvisioningStateDeleting,
		"failed":    WorkloadNetworkDhcpProvisioningStateFailed,
		"succeeded": WorkloadNetworkDhcpProvisioningStateSucceeded,
		"updating":  WorkloadNetworkDhcpProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WorkloadNetworkDhcpProvisioningState(input)
	return &out, nil
}

type WorkloadNetworkDnsServiceProvisioningState string

const (
	WorkloadNetworkDnsServiceProvisioningStateBuilding  WorkloadNetworkDnsServiceProvisioningState = "Building"
	WorkloadNetworkDnsServiceProvisioningStateDeleting  WorkloadNetworkDnsServiceProvisioningState = "Deleting"
	WorkloadNetworkDnsServiceProvisioningStateFailed    WorkloadNetworkDnsServiceProvisioningState = "Failed"
	WorkloadNetworkDnsServiceProvisioningStateSucceeded WorkloadNetworkDnsServiceProvisioningState = "Succeeded"
	WorkloadNetworkDnsServiceProvisioningStateUpdating  WorkloadNetworkDnsServiceProvisioningState = "Updating"
)

func PossibleValuesForWorkloadNetworkDnsServiceProvisioningState() []string {
	return []string{
		string(WorkloadNetworkDnsServiceProvisioningStateBuilding),
		string(WorkloadNetworkDnsServiceProvisioningStateDeleting),
		string(WorkloadNetworkDnsServiceProvisioningStateFailed),
		string(WorkloadNetworkDnsServiceProvisioningStateSucceeded),
		string(WorkloadNetworkDnsServiceProvisioningStateUpdating),
	}
}

func parseWorkloadNetworkDnsServiceProvisioningState(input string) (*WorkloadNetworkDnsServiceProvisioningState, error) {
	vals := map[string]WorkloadNetworkDnsServiceProvisioningState{
		"building":  WorkloadNetworkDnsServiceProvisioningStateBuilding,
		"deleting":  WorkloadNetworkDnsServiceProvisioningStateDeleting,
		"failed":    WorkloadNetworkDnsServiceProvisioningStateFailed,
		"succeeded": WorkloadNetworkDnsServiceProvisioningStateSucceeded,
		"updating":  WorkloadNetworkDnsServiceProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WorkloadNetworkDnsServiceProvisioningState(input)
	return &out, nil
}

type WorkloadNetworkDnsZoneProvisioningState string

const (
	WorkloadNetworkDnsZoneProvisioningStateBuilding  WorkloadNetworkDnsZoneProvisioningState = "Building"
	WorkloadNetworkDnsZoneProvisioningStateDeleting  WorkloadNetworkDnsZoneProvisioningState = "Deleting"
	WorkloadNetworkDnsZoneProvisioningStateFailed    WorkloadNetworkDnsZoneProvisioningState = "Failed"
	WorkloadNetworkDnsZoneProvisioningStateSucceeded WorkloadNetworkDnsZoneProvisioningState = "Succeeded"
	WorkloadNetworkDnsZoneProvisioningStateUpdating  WorkloadNetworkDnsZoneProvisioningState = "Updating"
)

func PossibleValuesForWorkloadNetworkDnsZoneProvisioningState() []string {
	return []string{
		string(WorkloadNetworkDnsZoneProvisioningStateBuilding),
		string(WorkloadNetworkDnsZoneProvisioningStateDeleting),
		string(WorkloadNetworkDnsZoneProvisioningStateFailed),
		string(WorkloadNetworkDnsZoneProvisioningStateSucceeded),
		string(WorkloadNetworkDnsZoneProvisioningStateUpdating),
	}
}

func parseWorkloadNetworkDnsZoneProvisioningState(input string) (*WorkloadNetworkDnsZoneProvisioningState, error) {
	vals := map[string]WorkloadNetworkDnsZoneProvisioningState{
		"building":  WorkloadNetworkDnsZoneProvisioningStateBuilding,
		"deleting":  WorkloadNetworkDnsZoneProvisioningStateDeleting,
		"failed":    WorkloadNetworkDnsZoneProvisioningStateFailed,
		"succeeded": WorkloadNetworkDnsZoneProvisioningStateSucceeded,
		"updating":  WorkloadNetworkDnsZoneProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WorkloadNetworkDnsZoneProvisioningState(input)
	return &out, nil
}

type WorkloadNetworkPortMirroringProvisioningState string

const (
	WorkloadNetworkPortMirroringProvisioningStateBuilding  WorkloadNetworkPortMirroringProvisioningState = "Building"
	WorkloadNetworkPortMirroringProvisioningStateDeleting  WorkloadNetworkPortMirroringProvisioningState = "Deleting"
	WorkloadNetworkPortMirroringProvisioningStateFailed    WorkloadNetworkPortMirroringProvisioningState = "Failed"
	WorkloadNetworkPortMirroringProvisioningStateSucceeded WorkloadNetworkPortMirroringProvisioningState = "Succeeded"
	WorkloadNetworkPortMirroringProvisioningStateUpdating  WorkloadNetworkPortMirroringProvisioningState = "Updating"
)

func PossibleValuesForWorkloadNetworkPortMirroringProvisioningState() []string {
	return []string{
		string(WorkloadNetworkPortMirroringProvisioningStateBuilding),
		string(WorkloadNetworkPortMirroringProvisioningStateDeleting),
		string(WorkloadNetworkPortMirroringProvisioningStateFailed),
		string(WorkloadNetworkPortMirroringProvisioningStateSucceeded),
		string(WorkloadNetworkPortMirroringProvisioningStateUpdating),
	}
}

func parseWorkloadNetworkPortMirroringProvisioningState(input string) (*WorkloadNetworkPortMirroringProvisioningState, error) {
	vals := map[string]WorkloadNetworkPortMirroringProvisioningState{
		"building":  WorkloadNetworkPortMirroringProvisioningStateBuilding,
		"deleting":  WorkloadNetworkPortMirroringProvisioningStateDeleting,
		"failed":    WorkloadNetworkPortMirroringProvisioningStateFailed,
		"succeeded": WorkloadNetworkPortMirroringProvisioningStateSucceeded,
		"updating":  WorkloadNetworkPortMirroringProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WorkloadNetworkPortMirroringProvisioningState(input)
	return &out, nil
}

type WorkloadNetworkPublicIPProvisioningState string

const (
	WorkloadNetworkPublicIPProvisioningStateBuilding  WorkloadNetworkPublicIPProvisioningState = "Building"
	WorkloadNetworkPublicIPProvisioningStateDeleting  WorkloadNetworkPublicIPProvisioningState = "Deleting"
	WorkloadNetworkPublicIPProvisioningStateFailed    WorkloadNetworkPublicIPProvisioningState = "Failed"
	WorkloadNetworkPublicIPProvisioningStateSucceeded WorkloadNetworkPublicIPProvisioningState = "Succeeded"
	WorkloadNetworkPublicIPProvisioningStateUpdating  WorkloadNetworkPublicIPProvisioningState = "Updating"
)

func PossibleValuesForWorkloadNetworkPublicIPProvisioningState() []string {
	return []string{
		string(WorkloadNetworkPublicIPProvisioningStateBuilding),
		string(WorkloadNetworkPublicIPProvisioningStateDeleting),
		string(WorkloadNetworkPublicIPProvisioningStateFailed),
		string(WorkloadNetworkPublicIPProvisioningStateSucceeded),
		string(WorkloadNetworkPublicIPProvisioningStateUpdating),
	}
}

func parseWorkloadNetworkPublicIPProvisioningState(input string) (*WorkloadNetworkPublicIPProvisioningState, error) {
	vals := map[string]WorkloadNetworkPublicIPProvisioningState{
		"building":  WorkloadNetworkPublicIPProvisioningStateBuilding,
		"deleting":  WorkloadNetworkPublicIPProvisioningStateDeleting,
		"failed":    WorkloadNetworkPublicIPProvisioningStateFailed,
		"succeeded": WorkloadNetworkPublicIPProvisioningStateSucceeded,
		"updating":  WorkloadNetworkPublicIPProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WorkloadNetworkPublicIPProvisioningState(input)
	return &out, nil
}

type WorkloadNetworkSegmentProvisioningState string

const (
	WorkloadNetworkSegmentProvisioningStateBuilding  WorkloadNetworkSegmentProvisioningState = "Building"
	WorkloadNetworkSegmentProvisioningStateDeleting  WorkloadNetworkSegmentProvisioningState = "Deleting"
	WorkloadNetworkSegmentProvisioningStateFailed    WorkloadNetworkSegmentProvisioningState = "Failed"
	WorkloadNetworkSegmentProvisioningStateSucceeded WorkloadNetworkSegmentProvisioningState = "Succeeded"
	WorkloadNetworkSegmentProvisioningStateUpdating  WorkloadNetworkSegmentProvisioningState = "Updating"
)

func PossibleValuesForWorkloadNetworkSegmentProvisioningState() []string {
	return []string{
		string(WorkloadNetworkSegmentProvisioningStateBuilding),
		string(WorkloadNetworkSegmentProvisioningStateDeleting),
		string(WorkloadNetworkSegmentProvisioningStateFailed),
		string(WorkloadNetworkSegmentProvisioningStateSucceeded),
		string(WorkloadNetworkSegmentProvisioningStateUpdating),
	}
}

func parseWorkloadNetworkSegmentProvisioningState(input string) (*WorkloadNetworkSegmentProvisioningState, error) {
	vals := map[string]WorkloadNetworkSegmentProvisioningState{
		"building":  WorkloadNetworkSegmentProvisioningStateBuilding,
		"deleting":  WorkloadNetworkSegmentProvisioningStateDeleting,
		"failed":    WorkloadNetworkSegmentProvisioningStateFailed,
		"succeeded": WorkloadNetworkSegmentProvisioningStateSucceeded,
		"updating":  WorkloadNetworkSegmentProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WorkloadNetworkSegmentProvisioningState(input)
	return &out, nil
}

type WorkloadNetworkVMGroupProvisioningState string

const (
	WorkloadNetworkVMGroupProvisioningStateBuilding  WorkloadNetworkVMGroupProvisioningState = "Building"
	WorkloadNetworkVMGroupProvisioningStateDeleting  WorkloadNetworkVMGroupProvisioningState = "Deleting"
	WorkloadNetworkVMGroupProvisioningStateFailed    WorkloadNetworkVMGroupProvisioningState = "Failed"
	WorkloadNetworkVMGroupProvisioningStateSucceeded WorkloadNetworkVMGroupProvisioningState = "Succeeded"
	WorkloadNetworkVMGroupProvisioningStateUpdating  WorkloadNetworkVMGroupProvisioningState = "Updating"
)

func PossibleValuesForWorkloadNetworkVMGroupProvisioningState() []string {
	return []string{
		string(WorkloadNetworkVMGroupProvisioningStateBuilding),
		string(WorkloadNetworkVMGroupProvisioningStateDeleting),
		string(WorkloadNetworkVMGroupProvisioningStateFailed),
		string(WorkloadNetworkVMGroupProvisioningStateSucceeded),
		string(WorkloadNetworkVMGroupProvisioningStateUpdating),
	}
}

func parseWorkloadNetworkVMGroupProvisioningState(input string) (*WorkloadNetworkVMGroupProvisioningState, error) {
	vals := map[string]WorkloadNetworkVMGroupProvisioningState{
		"building":  WorkloadNetworkVMGroupProvisioningStateBuilding,
		"deleting":  WorkloadNetworkVMGroupProvisioningStateDeleting,
		"failed":    WorkloadNetworkVMGroupProvisioningStateFailed,
		"succeeded": WorkloadNetworkVMGroupProvisioningStateSucceeded,
		"updating":  WorkloadNetworkVMGroupProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WorkloadNetworkVMGroupProvisioningState(input)
	return &out, nil
}
