package simpolicy

import "strings"

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

func parsePduSessionType(input string) (*PduSessionType, error) {
	vals := map[string]PduSessionType{
		"ipv4": PduSessionTypeIPvFour,
		"ipv6": PduSessionTypeIPvSix,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PduSessionType(input)
	return &out, nil
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

func parsePreemptionCapability(input string) (*PreemptionCapability, error) {
	vals := map[string]PreemptionCapability{
		"maypreempt": PreemptionCapabilityMayPreempt,
		"notpreempt": PreemptionCapabilityNotPreempt,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PreemptionCapability(input)
	return &out, nil
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

func parsePreemptionVulnerability(input string) (*PreemptionVulnerability, error) {
	vals := map[string]PreemptionVulnerability{
		"notpreemptable": PreemptionVulnerabilityNotPreemptable,
		"preemptable":    PreemptionVulnerabilityPreemptable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PreemptionVulnerability(input)
	return &out, nil
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

func parseProvisioningState(input string) (*ProvisioningState, error) {
	vals := map[string]ProvisioningState{
		"accepted":  ProvisioningStateAccepted,
		"canceled":  ProvisioningStateCanceled,
		"deleted":   ProvisioningStateDeleted,
		"deleting":  ProvisioningStateDeleting,
		"failed":    ProvisioningStateFailed,
		"succeeded": ProvisioningStateSucceeded,
		"unknown":   ProvisioningStateUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningState(input)
	return &out, nil
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

func parseSiteProvisioningState(input string) (*SiteProvisioningState, error) {
	vals := map[string]SiteProvisioningState{
		"adding":        SiteProvisioningStateAdding,
		"deleting":      SiteProvisioningStateDeleting,
		"failed":        SiteProvisioningStateFailed,
		"notapplicable": SiteProvisioningStateNotApplicable,
		"provisioned":   SiteProvisioningStateProvisioned,
		"updating":      SiteProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SiteProvisioningState(input)
	return &out, nil
}
