package workloadnetworks

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

type WorkloadNetworkDhcpProvisioningState string

const (
	WorkloadNetworkDhcpProvisioningStateBuilding  WorkloadNetworkDhcpProvisioningState = "Building"
	WorkloadNetworkDhcpProvisioningStateCanceled  WorkloadNetworkDhcpProvisioningState = "Canceled"
	WorkloadNetworkDhcpProvisioningStateDeleting  WorkloadNetworkDhcpProvisioningState = "Deleting"
	WorkloadNetworkDhcpProvisioningStateFailed    WorkloadNetworkDhcpProvisioningState = "Failed"
	WorkloadNetworkDhcpProvisioningStateSucceeded WorkloadNetworkDhcpProvisioningState = "Succeeded"
	WorkloadNetworkDhcpProvisioningStateUpdating  WorkloadNetworkDhcpProvisioningState = "Updating"
)

func PossibleValuesForWorkloadNetworkDhcpProvisioningState() []string {
	return []string{
		string(WorkloadNetworkDhcpProvisioningStateBuilding),
		string(WorkloadNetworkDhcpProvisioningStateCanceled),
		string(WorkloadNetworkDhcpProvisioningStateDeleting),
		string(WorkloadNetworkDhcpProvisioningStateFailed),
		string(WorkloadNetworkDhcpProvisioningStateSucceeded),
		string(WorkloadNetworkDhcpProvisioningStateUpdating),
	}
}

type WorkloadNetworkDnsServiceProvisioningState string

const (
	WorkloadNetworkDnsServiceProvisioningStateBuilding  WorkloadNetworkDnsServiceProvisioningState = "Building"
	WorkloadNetworkDnsServiceProvisioningStateCanceled  WorkloadNetworkDnsServiceProvisioningState = "Canceled"
	WorkloadNetworkDnsServiceProvisioningStateDeleting  WorkloadNetworkDnsServiceProvisioningState = "Deleting"
	WorkloadNetworkDnsServiceProvisioningStateFailed    WorkloadNetworkDnsServiceProvisioningState = "Failed"
	WorkloadNetworkDnsServiceProvisioningStateSucceeded WorkloadNetworkDnsServiceProvisioningState = "Succeeded"
	WorkloadNetworkDnsServiceProvisioningStateUpdating  WorkloadNetworkDnsServiceProvisioningState = "Updating"
)

func PossibleValuesForWorkloadNetworkDnsServiceProvisioningState() []string {
	return []string{
		string(WorkloadNetworkDnsServiceProvisioningStateBuilding),
		string(WorkloadNetworkDnsServiceProvisioningStateCanceled),
		string(WorkloadNetworkDnsServiceProvisioningStateDeleting),
		string(WorkloadNetworkDnsServiceProvisioningStateFailed),
		string(WorkloadNetworkDnsServiceProvisioningStateSucceeded),
		string(WorkloadNetworkDnsServiceProvisioningStateUpdating),
	}
}

type WorkloadNetworkDnsZoneProvisioningState string

const (
	WorkloadNetworkDnsZoneProvisioningStateBuilding  WorkloadNetworkDnsZoneProvisioningState = "Building"
	WorkloadNetworkDnsZoneProvisioningStateCanceled  WorkloadNetworkDnsZoneProvisioningState = "Canceled"
	WorkloadNetworkDnsZoneProvisioningStateDeleting  WorkloadNetworkDnsZoneProvisioningState = "Deleting"
	WorkloadNetworkDnsZoneProvisioningStateFailed    WorkloadNetworkDnsZoneProvisioningState = "Failed"
	WorkloadNetworkDnsZoneProvisioningStateSucceeded WorkloadNetworkDnsZoneProvisioningState = "Succeeded"
	WorkloadNetworkDnsZoneProvisioningStateUpdating  WorkloadNetworkDnsZoneProvisioningState = "Updating"
)

func PossibleValuesForWorkloadNetworkDnsZoneProvisioningState() []string {
	return []string{
		string(WorkloadNetworkDnsZoneProvisioningStateBuilding),
		string(WorkloadNetworkDnsZoneProvisioningStateCanceled),
		string(WorkloadNetworkDnsZoneProvisioningStateDeleting),
		string(WorkloadNetworkDnsZoneProvisioningStateFailed),
		string(WorkloadNetworkDnsZoneProvisioningStateSucceeded),
		string(WorkloadNetworkDnsZoneProvisioningStateUpdating),
	}
}

type WorkloadNetworkPortMirroringProvisioningState string

const (
	WorkloadNetworkPortMirroringProvisioningStateBuilding  WorkloadNetworkPortMirroringProvisioningState = "Building"
	WorkloadNetworkPortMirroringProvisioningStateCanceled  WorkloadNetworkPortMirroringProvisioningState = "Canceled"
	WorkloadNetworkPortMirroringProvisioningStateDeleting  WorkloadNetworkPortMirroringProvisioningState = "Deleting"
	WorkloadNetworkPortMirroringProvisioningStateFailed    WorkloadNetworkPortMirroringProvisioningState = "Failed"
	WorkloadNetworkPortMirroringProvisioningStateSucceeded WorkloadNetworkPortMirroringProvisioningState = "Succeeded"
	WorkloadNetworkPortMirroringProvisioningStateUpdating  WorkloadNetworkPortMirroringProvisioningState = "Updating"
)

func PossibleValuesForWorkloadNetworkPortMirroringProvisioningState() []string {
	return []string{
		string(WorkloadNetworkPortMirroringProvisioningStateBuilding),
		string(WorkloadNetworkPortMirroringProvisioningStateCanceled),
		string(WorkloadNetworkPortMirroringProvisioningStateDeleting),
		string(WorkloadNetworkPortMirroringProvisioningStateFailed),
		string(WorkloadNetworkPortMirroringProvisioningStateSucceeded),
		string(WorkloadNetworkPortMirroringProvisioningStateUpdating),
	}
}

type WorkloadNetworkPublicIPProvisioningState string

const (
	WorkloadNetworkPublicIPProvisioningStateBuilding  WorkloadNetworkPublicIPProvisioningState = "Building"
	WorkloadNetworkPublicIPProvisioningStateCanceled  WorkloadNetworkPublicIPProvisioningState = "Canceled"
	WorkloadNetworkPublicIPProvisioningStateDeleting  WorkloadNetworkPublicIPProvisioningState = "Deleting"
	WorkloadNetworkPublicIPProvisioningStateFailed    WorkloadNetworkPublicIPProvisioningState = "Failed"
	WorkloadNetworkPublicIPProvisioningStateSucceeded WorkloadNetworkPublicIPProvisioningState = "Succeeded"
	WorkloadNetworkPublicIPProvisioningStateUpdating  WorkloadNetworkPublicIPProvisioningState = "Updating"
)

func PossibleValuesForWorkloadNetworkPublicIPProvisioningState() []string {
	return []string{
		string(WorkloadNetworkPublicIPProvisioningStateBuilding),
		string(WorkloadNetworkPublicIPProvisioningStateCanceled),
		string(WorkloadNetworkPublicIPProvisioningStateDeleting),
		string(WorkloadNetworkPublicIPProvisioningStateFailed),
		string(WorkloadNetworkPublicIPProvisioningStateSucceeded),
		string(WorkloadNetworkPublicIPProvisioningStateUpdating),
	}
}

type WorkloadNetworkSegmentProvisioningState string

const (
	WorkloadNetworkSegmentProvisioningStateBuilding  WorkloadNetworkSegmentProvisioningState = "Building"
	WorkloadNetworkSegmentProvisioningStateCanceled  WorkloadNetworkSegmentProvisioningState = "Canceled"
	WorkloadNetworkSegmentProvisioningStateDeleting  WorkloadNetworkSegmentProvisioningState = "Deleting"
	WorkloadNetworkSegmentProvisioningStateFailed    WorkloadNetworkSegmentProvisioningState = "Failed"
	WorkloadNetworkSegmentProvisioningStateSucceeded WorkloadNetworkSegmentProvisioningState = "Succeeded"
	WorkloadNetworkSegmentProvisioningStateUpdating  WorkloadNetworkSegmentProvisioningState = "Updating"
)

func PossibleValuesForWorkloadNetworkSegmentProvisioningState() []string {
	return []string{
		string(WorkloadNetworkSegmentProvisioningStateBuilding),
		string(WorkloadNetworkSegmentProvisioningStateCanceled),
		string(WorkloadNetworkSegmentProvisioningStateDeleting),
		string(WorkloadNetworkSegmentProvisioningStateFailed),
		string(WorkloadNetworkSegmentProvisioningStateSucceeded),
		string(WorkloadNetworkSegmentProvisioningStateUpdating),
	}
}

type WorkloadNetworkVMGroupProvisioningState string

const (
	WorkloadNetworkVMGroupProvisioningStateBuilding  WorkloadNetworkVMGroupProvisioningState = "Building"
	WorkloadNetworkVMGroupProvisioningStateCanceled  WorkloadNetworkVMGroupProvisioningState = "Canceled"
	WorkloadNetworkVMGroupProvisioningStateDeleting  WorkloadNetworkVMGroupProvisioningState = "Deleting"
	WorkloadNetworkVMGroupProvisioningStateFailed    WorkloadNetworkVMGroupProvisioningState = "Failed"
	WorkloadNetworkVMGroupProvisioningStateSucceeded WorkloadNetworkVMGroupProvisioningState = "Succeeded"
	WorkloadNetworkVMGroupProvisioningStateUpdating  WorkloadNetworkVMGroupProvisioningState = "Updating"
)

func PossibleValuesForWorkloadNetworkVMGroupProvisioningState() []string {
	return []string{
		string(WorkloadNetworkVMGroupProvisioningStateBuilding),
		string(WorkloadNetworkVMGroupProvisioningStateCanceled),
		string(WorkloadNetworkVMGroupProvisioningStateDeleting),
		string(WorkloadNetworkVMGroupProvisioningStateFailed),
		string(WorkloadNetworkVMGroupProvisioningStateSucceeded),
		string(WorkloadNetworkVMGroupProvisioningStateUpdating),
	}
}
