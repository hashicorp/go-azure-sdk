package deviceupdates

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationType string

const (
	AuthenticationTypeKeyBased AuthenticationType = "KeyBased"
)

func PossibleValuesForAuthenticationType() []string {
	return []string{
		string(AuthenticationTypeKeyBased),
	}
}

type CheckNameAvailabilityReason string

const (
	CheckNameAvailabilityReasonAlreadyExists CheckNameAvailabilityReason = "AlreadyExists"
	CheckNameAvailabilityReasonInvalid       CheckNameAvailabilityReason = "Invalid"
)

func PossibleValuesForCheckNameAvailabilityReason() []string {
	return []string{
		string(CheckNameAvailabilityReasonAlreadyExists),
		string(CheckNameAvailabilityReasonInvalid),
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

type ProvisioningState string

const (
	ProvisioningStateAccepted  ProvisioningState = "Accepted"
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleted   ProvisioningState = "Deleted"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateAccepted),
		string(ProvisioningStateCanceled),
		string(ProvisioningStateCreating),
		string(ProvisioningStateDeleted),
		string(ProvisioningStateFailed),
		string(ProvisioningStateSucceeded),
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

type Role string

const (
	RoleFailover Role = "Failover"
	RolePrimary  Role = "Primary"
)

func PossibleValuesForRole() []string {
	return []string{
		string(RoleFailover),
		string(RolePrimary),
	}
}

type SKU string

const (
	SKUFree     SKU = "Free"
	SKUStandard SKU = "Standard"
)

func PossibleValuesForSKU() []string {
	return []string{
		string(SKUFree),
		string(SKUStandard),
	}
}
