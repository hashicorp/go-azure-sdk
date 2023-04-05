package fhirservices

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FhirServiceKind string

const (
	FhirServiceKindFhirNegativeRFour    FhirServiceKind = "fhir-R4"
	FhirServiceKindFhirNegativeStuThree FhirServiceKind = "fhir-Stu3"
)

func PossibleValuesForFhirServiceKind() []string {
	return []string{
		string(FhirServiceKindFhirNegativeRFour),
		string(FhirServiceKindFhirNegativeStuThree),
	}
}

type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityTypeNone           ManagedServiceIdentityType = "None"
	ManagedServiceIdentityTypeSystemAssigned ManagedServiceIdentityType = "SystemAssigned"
)

func PossibleValuesForManagedServiceIdentityType() []string {
	return []string{
		string(ManagedServiceIdentityTypeNone),
		string(ManagedServiceIdentityTypeSystemAssigned),
	}
}

type ProvisioningState string

const (
	ProvisioningStateAccepted          ProvisioningState = "Accepted"
	ProvisioningStateCanceled          ProvisioningState = "Canceled"
	ProvisioningStateCreating          ProvisioningState = "Creating"
	ProvisioningStateDeleting          ProvisioningState = "Deleting"
	ProvisioningStateDeprovisioned     ProvisioningState = "Deprovisioned"
	ProvisioningStateFailed            ProvisioningState = "Failed"
	ProvisioningStateMoving            ProvisioningState = "Moving"
	ProvisioningStateSucceeded         ProvisioningState = "Succeeded"
	ProvisioningStateSuspended         ProvisioningState = "Suspended"
	ProvisioningStateSystemMaintenance ProvisioningState = "SystemMaintenance"
	ProvisioningStateUpdating          ProvisioningState = "Updating"
	ProvisioningStateVerifying         ProvisioningState = "Verifying"
	ProvisioningStateWarned            ProvisioningState = "Warned"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateAccepted),
		string(ProvisioningStateCanceled),
		string(ProvisioningStateCreating),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateDeprovisioned),
		string(ProvisioningStateFailed),
		string(ProvisioningStateMoving),
		string(ProvisioningStateSucceeded),
		string(ProvisioningStateSuspended),
		string(ProvisioningStateSystemMaintenance),
		string(ProvisioningStateUpdating),
		string(ProvisioningStateVerifying),
		string(ProvisioningStateWarned),
	}
}
