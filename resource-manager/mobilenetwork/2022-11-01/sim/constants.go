package sim

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

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

type SimState string

const (
	SimStateDisabled SimState = "Disabled"
	SimStateEnabled  SimState = "Enabled"
	SimStateInvalid  SimState = "Invalid"
)

func PossibleValuesForSimState() []string {
	return []string{
		string(SimStateDisabled),
		string(SimStateEnabled),
		string(SimStateInvalid),
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
