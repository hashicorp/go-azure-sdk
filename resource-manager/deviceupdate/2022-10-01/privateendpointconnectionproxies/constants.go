package privateendpointconnectionproxies

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateEndpointConnectionProxyProvisioningState string

const (
	PrivateEndpointConnectionProxyProvisioningStateCreating  PrivateEndpointConnectionProxyProvisioningState = "Creating"
	PrivateEndpointConnectionProxyProvisioningStateDeleting  PrivateEndpointConnectionProxyProvisioningState = "Deleting"
	PrivateEndpointConnectionProxyProvisioningStateFailed    PrivateEndpointConnectionProxyProvisioningState = "Failed"
	PrivateEndpointConnectionProxyProvisioningStateSucceeded PrivateEndpointConnectionProxyProvisioningState = "Succeeded"
)

func PossibleValuesForPrivateEndpointConnectionProxyProvisioningState() []string {
	return []string{
		string(PrivateEndpointConnectionProxyProvisioningStateCreating),
		string(PrivateEndpointConnectionProxyProvisioningStateDeleting),
		string(PrivateEndpointConnectionProxyProvisioningStateFailed),
		string(PrivateEndpointConnectionProxyProvisioningStateSucceeded),
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
