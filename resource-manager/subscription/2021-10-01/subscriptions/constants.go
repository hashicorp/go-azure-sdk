package subscriptions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AcceptOwnership string

const (
	AcceptOwnershipCompleted AcceptOwnership = "Completed"
	AcceptOwnershipExpired   AcceptOwnership = "Expired"
	AcceptOwnershipPending   AcceptOwnership = "Pending"
)

func PossibleValuesForAcceptOwnership() []string {
	return []string{
		string(AcceptOwnershipCompleted),
		string(AcceptOwnershipExpired),
		string(AcceptOwnershipPending),
	}
}

type Provisioning string

const (
	ProvisioningAccepted  Provisioning = "Accepted"
	ProvisioningPending   Provisioning = "Pending"
	ProvisioningSucceeded Provisioning = "Succeeded"
)

func PossibleValuesForProvisioning() []string {
	return []string{
		string(ProvisioningAccepted),
		string(ProvisioningPending),
		string(ProvisioningSucceeded),
	}
}

type ProvisioningState string

const (
	ProvisioningStateAccepted  ProvisioningState = "Accepted"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateAccepted),
		string(ProvisioningStateFailed),
		string(ProvisioningStateSucceeded),
	}
}

type Workload string

const (
	WorkloadDevTest    Workload = "DevTest"
	WorkloadProduction Workload = "Production"
)

func PossibleValuesForWorkload() []string {
	return []string{
		string(WorkloadDevTest),
		string(WorkloadProduction),
	}
}
