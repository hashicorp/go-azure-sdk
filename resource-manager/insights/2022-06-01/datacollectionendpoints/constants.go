package datacollectionendpoints

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KnownDataCollectionEndpointProvisioningState string

const (
	KnownDataCollectionEndpointProvisioningStateCanceled  KnownDataCollectionEndpointProvisioningState = "Canceled"
	KnownDataCollectionEndpointProvisioningStateCreating  KnownDataCollectionEndpointProvisioningState = "Creating"
	KnownDataCollectionEndpointProvisioningStateDeleting  KnownDataCollectionEndpointProvisioningState = "Deleting"
	KnownDataCollectionEndpointProvisioningStateFailed    KnownDataCollectionEndpointProvisioningState = "Failed"
	KnownDataCollectionEndpointProvisioningStateSucceeded KnownDataCollectionEndpointProvisioningState = "Succeeded"
	KnownDataCollectionEndpointProvisioningStateUpdating  KnownDataCollectionEndpointProvisioningState = "Updating"
)

func PossibleValuesForKnownDataCollectionEndpointProvisioningState() []string {
	return []string{
		string(KnownDataCollectionEndpointProvisioningStateCanceled),
		string(KnownDataCollectionEndpointProvisioningStateCreating),
		string(KnownDataCollectionEndpointProvisioningStateDeleting),
		string(KnownDataCollectionEndpointProvisioningStateFailed),
		string(KnownDataCollectionEndpointProvisioningStateSucceeded),
		string(KnownDataCollectionEndpointProvisioningStateUpdating),
	}
}

type KnownDataCollectionEndpointResourceKind string

const (
	KnownDataCollectionEndpointResourceKindLinux   KnownDataCollectionEndpointResourceKind = "Linux"
	KnownDataCollectionEndpointResourceKindWindows KnownDataCollectionEndpointResourceKind = "Windows"
)

func PossibleValuesForKnownDataCollectionEndpointResourceKind() []string {
	return []string{
		string(KnownDataCollectionEndpointResourceKindLinux),
		string(KnownDataCollectionEndpointResourceKindWindows),
	}
}

type KnownLocationSpecProvisioningStatus string

const (
	KnownLocationSpecProvisioningStatusCanceled  KnownLocationSpecProvisioningStatus = "Canceled"
	KnownLocationSpecProvisioningStatusCreating  KnownLocationSpecProvisioningStatus = "Creating"
	KnownLocationSpecProvisioningStatusDeleting  KnownLocationSpecProvisioningStatus = "Deleting"
	KnownLocationSpecProvisioningStatusFailed    KnownLocationSpecProvisioningStatus = "Failed"
	KnownLocationSpecProvisioningStatusSucceeded KnownLocationSpecProvisioningStatus = "Succeeded"
	KnownLocationSpecProvisioningStatusUpdating  KnownLocationSpecProvisioningStatus = "Updating"
)

func PossibleValuesForKnownLocationSpecProvisioningStatus() []string {
	return []string{
		string(KnownLocationSpecProvisioningStatusCanceled),
		string(KnownLocationSpecProvisioningStatusCreating),
		string(KnownLocationSpecProvisioningStatusDeleting),
		string(KnownLocationSpecProvisioningStatusFailed),
		string(KnownLocationSpecProvisioningStatusSucceeded),
		string(KnownLocationSpecProvisioningStatusUpdating),
	}
}

type KnownPublicNetworkAccessOptions string

const (
	KnownPublicNetworkAccessOptionsDisabled           KnownPublicNetworkAccessOptions = "Disabled"
	KnownPublicNetworkAccessOptionsEnabled            KnownPublicNetworkAccessOptions = "Enabled"
	KnownPublicNetworkAccessOptionsSecuredByPerimeter KnownPublicNetworkAccessOptions = "SecuredByPerimeter"
)

func PossibleValuesForKnownPublicNetworkAccessOptions() []string {
	return []string{
		string(KnownPublicNetworkAccessOptionsDisabled),
		string(KnownPublicNetworkAccessOptionsEnabled),
		string(KnownPublicNetworkAccessOptionsSecuredByPerimeter),
	}
}
