package mhsmlistprivateendpointconnections

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActionsRequired string

const (
	ActionsRequiredNone ActionsRequired = "None"
)

func PossibleValuesForActionsRequired() []string {
	return []string{
		string(ActionsRequiredNone),
	}
}

type ManagedHsmSkuFamily string

const (
	ManagedHsmSkuFamilyB ManagedHsmSkuFamily = "B"
)

func PossibleValuesForManagedHsmSkuFamily() []string {
	return []string{
		string(ManagedHsmSkuFamilyB),
	}
}

type ManagedHsmSkuName string

const (
	ManagedHsmSkuNameCustomBThreeTwo ManagedHsmSkuName = "Custom_B32"
	ManagedHsmSkuNameStandardBOne    ManagedHsmSkuName = "Standard_B1"
)

func PossibleValuesForManagedHsmSkuName() []string {
	return []string{
		string(ManagedHsmSkuNameCustomBThreeTwo),
		string(ManagedHsmSkuNameStandardBOne),
	}
}

type PrivateEndpointConnectionProvisioningState string

const (
	PrivateEndpointConnectionProvisioningStateCreating     PrivateEndpointConnectionProvisioningState = "Creating"
	PrivateEndpointConnectionProvisioningStateDeleting     PrivateEndpointConnectionProvisioningState = "Deleting"
	PrivateEndpointConnectionProvisioningStateDisconnected PrivateEndpointConnectionProvisioningState = "Disconnected"
	PrivateEndpointConnectionProvisioningStateFailed       PrivateEndpointConnectionProvisioningState = "Failed"
	PrivateEndpointConnectionProvisioningStateSucceeded    PrivateEndpointConnectionProvisioningState = "Succeeded"
	PrivateEndpointConnectionProvisioningStateUpdating     PrivateEndpointConnectionProvisioningState = "Updating"
)

func PossibleValuesForPrivateEndpointConnectionProvisioningState() []string {
	return []string{
		string(PrivateEndpointConnectionProvisioningStateCreating),
		string(PrivateEndpointConnectionProvisioningStateDeleting),
		string(PrivateEndpointConnectionProvisioningStateDisconnected),
		string(PrivateEndpointConnectionProvisioningStateFailed),
		string(PrivateEndpointConnectionProvisioningStateSucceeded),
		string(PrivateEndpointConnectionProvisioningStateUpdating),
	}
}

type PrivateEndpointServiceConnectionStatus string

const (
	PrivateEndpointServiceConnectionStatusApproved     PrivateEndpointServiceConnectionStatus = "Approved"
	PrivateEndpointServiceConnectionStatusDisconnected PrivateEndpointServiceConnectionStatus = "Disconnected"
	PrivateEndpointServiceConnectionStatusPending      PrivateEndpointServiceConnectionStatus = "Pending"
	PrivateEndpointServiceConnectionStatusRejected     PrivateEndpointServiceConnectionStatus = "Rejected"
)

func PossibleValuesForPrivateEndpointServiceConnectionStatus() []string {
	return []string{
		string(PrivateEndpointServiceConnectionStatusApproved),
		string(PrivateEndpointServiceConnectionStatusDisconnected),
		string(PrivateEndpointServiceConnectionStatusPending),
		string(PrivateEndpointServiceConnectionStatusRejected),
	}
}
