package digitaltwinsinstance

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

type ProvisioningState string

const (
	ProvisioningStateCanceled     ProvisioningState = "Canceled"
	ProvisioningStateDeleted      ProvisioningState = "Deleted"
	ProvisioningStateDeleting     ProvisioningState = "Deleting"
	ProvisioningStateFailed       ProvisioningState = "Failed"
	ProvisioningStateMoving       ProvisioningState = "Moving"
	ProvisioningStateProvisioning ProvisioningState = "Provisioning"
	ProvisioningStateRestoring    ProvisioningState = "Restoring"
	ProvisioningStateSucceeded    ProvisioningState = "Succeeded"
	ProvisioningStateSuspending   ProvisioningState = "Suspending"
	ProvisioningStateUpdating     ProvisioningState = "Updating"
	ProvisioningStateWarning      ProvisioningState = "Warning"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCanceled),
		string(ProvisioningStateDeleted),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateMoving),
		string(ProvisioningStateProvisioning),
		string(ProvisioningStateRestoring),
		string(ProvisioningStateSucceeded),
		string(ProvisioningStateSuspending),
		string(ProvisioningStateUpdating),
		string(ProvisioningStateWarning),
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
