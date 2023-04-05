package fhirservices

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FhirResourceVersionPolicy string

const (
	FhirResourceVersionPolicyNoNegativeversion       FhirResourceVersionPolicy = "no-version"
	FhirResourceVersionPolicyVersioned               FhirResourceVersionPolicy = "versioned"
	FhirResourceVersionPolicyVersionedNegativeupdate FhirResourceVersionPolicy = "versioned-update"
)

func PossibleValuesForFhirResourceVersionPolicy() []string {
	return []string{
		string(FhirResourceVersionPolicyNoNegativeversion),
		string(FhirResourceVersionPolicyVersioned),
		string(FhirResourceVersionPolicyVersionedNegativeupdate),
	}
}

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

type PrivateEndpointConnectionProvisioningState string

const (
	PrivateEndpointConnectionProvisioningStateCreating  PrivateEndpointConnectionProvisioningState = "Creating"
	PrivateEndpointConnectionProvisioningStateDeleting  PrivateEndpointConnectionProvisioningState = "Deleting"
	PrivateEndpointConnectionProvisioningStateFailed    PrivateEndpointConnectionProvisioningState = "Failed"
	PrivateEndpointConnectionProvisioningStateSucceeded PrivateEndpointConnectionProvisioningState = "Succeeded"
)

func PossibleValuesForPrivateEndpointConnectionProvisioningState() []string {
	return []string{
		string(PrivateEndpointConnectionProvisioningStateCreating),
		string(PrivateEndpointConnectionProvisioningStateDeleting),
		string(PrivateEndpointConnectionProvisioningStateFailed),
		string(PrivateEndpointConnectionProvisioningStateSucceeded),
	}
}

type PrivateEndpointServiceConnectionStatus string

const (
	PrivateEndpointServiceConnectionStatusApproved PrivateEndpointServiceConnectionStatus = "Approved"
	PrivateEndpointServiceConnectionStatusPending  PrivateEndpointServiceConnectionStatus = "Pending"
	PrivateEndpointServiceConnectionStatusRejected PrivateEndpointServiceConnectionStatus = "Rejected"
)

func PossibleValuesForPrivateEndpointServiceConnectionStatus() []string {
	return []string{
		string(PrivateEndpointServiceConnectionStatusApproved),
		string(PrivateEndpointServiceConnectionStatusPending),
		string(PrivateEndpointServiceConnectionStatusRejected),
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

type PublicNetworkAccess string

const (
	PublicNetworkAccessDisabled PublicNetworkAccess = "Disabled"
	PublicNetworkAccessEnabled  PublicNetworkAccess = "Enabled"
)

func PossibleValuesForPublicNetworkAccess() []string {
	return []string{
		string(PublicNetworkAccessDisabled),
		string(PublicNetworkAccessEnabled),
	}
}

type ServiceEventState string

const (
	ServiceEventStateDisabled ServiceEventState = "Disabled"
	ServiceEventStateEnabled  ServiceEventState = "Enabled"
	ServiceEventStateUpdating ServiceEventState = "Updating"
)

func PossibleValuesForServiceEventState() []string {
	return []string{
		string(ServiceEventStateDisabled),
		string(ServiceEventStateEnabled),
		string(ServiceEventStateUpdating),
	}
}
