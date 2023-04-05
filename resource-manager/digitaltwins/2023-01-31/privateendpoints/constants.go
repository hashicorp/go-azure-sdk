package privateendpoints

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectionPropertiesProvisioningState string

const (
	ConnectionPropertiesProvisioningStateApproved     ConnectionPropertiesProvisioningState = "Approved"
	ConnectionPropertiesProvisioningStateDisconnected ConnectionPropertiesProvisioningState = "Disconnected"
	ConnectionPropertiesProvisioningStatePending      ConnectionPropertiesProvisioningState = "Pending"
	ConnectionPropertiesProvisioningStateRejected     ConnectionPropertiesProvisioningState = "Rejected"
)

func PossibleValuesForConnectionPropertiesProvisioningState() []string {
	return []string{
		string(ConnectionPropertiesProvisioningStateApproved),
		string(ConnectionPropertiesProvisioningStateDisconnected),
		string(ConnectionPropertiesProvisioningStatePending),
		string(ConnectionPropertiesProvisioningStateRejected),
	}
}

type PrivateLinkServiceConnectionStatus string

const (
	PrivateLinkServiceConnectionStatusApproved     PrivateLinkServiceConnectionStatus = "Approved"
	PrivateLinkServiceConnectionStatusDisconnected PrivateLinkServiceConnectionStatus = "Disconnected"
	PrivateLinkServiceConnectionStatusPending      PrivateLinkServiceConnectionStatus = "Pending"
	PrivateLinkServiceConnectionStatusRejected     PrivateLinkServiceConnectionStatus = "Rejected"
)

func PossibleValuesForPrivateLinkServiceConnectionStatus() []string {
	return []string{
		string(PrivateLinkServiceConnectionStatusApproved),
		string(PrivateLinkServiceConnectionStatusDisconnected),
		string(PrivateLinkServiceConnectionStatusPending),
		string(PrivateLinkServiceConnectionStatusRejected),
	}
}
