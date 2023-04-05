package simpolicies

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PduSessionType string

const (
	PduSessionTypeIPvFour PduSessionType = "IPv4"
	PduSessionTypeIPvSix  PduSessionType = "IPv6"
)

func PossibleValuesForPduSessionType() []string {
	return []string{
		string(PduSessionTypeIPvFour),
		string(PduSessionTypeIPvSix),
	}
}

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

type SiteProvisioningState string

const (
	SiteProvisioningStateAdding        SiteProvisioningState = "Adding"
	SiteProvisioningStateDeleting      SiteProvisioningState = "Deleting"
	SiteProvisioningStateFailed        SiteProvisioningState = "Failed"
	SiteProvisioningStateNotApplicable SiteProvisioningState = "NotApplicable"
	SiteProvisioningStateProvisioned   SiteProvisioningState = "Provisioned"
	SiteProvisioningStateUpdating      SiteProvisioningState = "Updating"
)

func PossibleValuesForSiteProvisioningState() []string {
	return []string{
		string(SiteProvisioningStateAdding),
		string(SiteProvisioningStateDeleting),
		string(SiteProvisioningStateFailed),
		string(SiteProvisioningStateNotApplicable),
		string(SiteProvisioningStateProvisioned),
		string(SiteProvisioningStateUpdating),
	}
}
