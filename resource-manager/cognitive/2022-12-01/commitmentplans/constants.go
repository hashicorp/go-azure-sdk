package commitmentplans

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommitmentPlanProvisioningState string

const (
	CommitmentPlanProvisioningStateAccepted  CommitmentPlanProvisioningState = "Accepted"
	CommitmentPlanProvisioningStateCanceled  CommitmentPlanProvisioningState = "Canceled"
	CommitmentPlanProvisioningStateCreating  CommitmentPlanProvisioningState = "Creating"
	CommitmentPlanProvisioningStateDeleting  CommitmentPlanProvisioningState = "Deleting"
	CommitmentPlanProvisioningStateFailed    CommitmentPlanProvisioningState = "Failed"
	CommitmentPlanProvisioningStateMoving    CommitmentPlanProvisioningState = "Moving"
	CommitmentPlanProvisioningStateSucceeded CommitmentPlanProvisioningState = "Succeeded"
)

func PossibleValuesForCommitmentPlanProvisioningState() []string {
	return []string{
		string(CommitmentPlanProvisioningStateAccepted),
		string(CommitmentPlanProvisioningStateCanceled),
		string(CommitmentPlanProvisioningStateCreating),
		string(CommitmentPlanProvisioningStateDeleting),
		string(CommitmentPlanProvisioningStateFailed),
		string(CommitmentPlanProvisioningStateMoving),
		string(CommitmentPlanProvisioningStateSucceeded),
	}
}

type HostingModel string

const (
	HostingModelConnectedContainer    HostingModel = "ConnectedContainer"
	HostingModelDisconnectedContainer HostingModel = "DisconnectedContainer"
	HostingModelWeb                   HostingModel = "Web"
)

func PossibleValuesForHostingModel() []string {
	return []string{
		string(HostingModelConnectedContainer),
		string(HostingModelDisconnectedContainer),
		string(HostingModelWeb),
	}
}

type SkuTier string

const (
	SkuTierBasic      SkuTier = "Basic"
	SkuTierEnterprise SkuTier = "Enterprise"
	SkuTierFree       SkuTier = "Free"
	SkuTierPremium    SkuTier = "Premium"
	SkuTierStandard   SkuTier = "Standard"
)

func PossibleValuesForSkuTier() []string {
	return []string{
		string(SkuTierBasic),
		string(SkuTierEnterprise),
		string(SkuTierFree),
		string(SkuTierPremium),
		string(SkuTierStandard),
	}
}
