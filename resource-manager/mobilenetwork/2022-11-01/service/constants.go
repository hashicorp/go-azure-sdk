package service

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PreemptionCapability string

const (
	PreemptionCapabilityMayPreempt PreemptionCapability = "MayPreempt"
	PreemptionCapabilityNotPreempt PreemptionCapability = "NotPreempt"
)

func PossibleValuesForPreemptionCapability() []string {
	return []string{
		string(PreemptionCapabilityMayPreempt),
		string(PreemptionCapabilityNotPreempt),
	}
}

type PreemptionVulnerability string

const (
	PreemptionVulnerabilityNotPreemptable PreemptionVulnerability = "NotPreemptable"
	PreemptionVulnerabilityPreemptable    PreemptionVulnerability = "Preemptable"
)

func PossibleValuesForPreemptionVulnerability() []string {
	return []string{
		string(PreemptionVulnerabilityNotPreemptable),
		string(PreemptionVulnerabilityPreemptable),
	}
}

type ProvisioningState string

const (
	ProvisioningStateAccepted  ProvisioningState = "Accepted"
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateDeleted   ProvisioningState = "Deleted"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUnknown   ProvisioningState = "Unknown"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateAccepted),
		string(ProvisioningStateCanceled),
		string(ProvisioningStateDeleted),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateSucceeded),
		string(ProvisioningStateUnknown),
	}
}

type SdfDirection string

const (
	SdfDirectionBidirectional SdfDirection = "Bidirectional"
	SdfDirectionDownlink      SdfDirection = "Downlink"
	SdfDirectionUplink        SdfDirection = "Uplink"
)

func PossibleValuesForSdfDirection() []string {
	return []string{
		string(SdfDirectionBidirectional),
		string(SdfDirectionDownlink),
		string(SdfDirectionUplink),
	}
}

type TrafficControlPermission string

const (
	TrafficControlPermissionBlocked TrafficControlPermission = "Blocked"
	TrafficControlPermissionEnabled TrafficControlPermission = "Enabled"
)

func PossibleValuesForTrafficControlPermission() []string {
	return []string{
		string(TrafficControlPermissionBlocked),
		string(TrafficControlPermissionEnabled),
	}
}
