package batchaccount

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccountKeyType string

const (
	AccountKeyTypePrimary   AccountKeyType = "Primary"
	AccountKeyTypeSecondary AccountKeyType = "Secondary"
)

func PossibleValuesForAccountKeyType() []string {
	return []string{
		string(AccountKeyTypePrimary),
		string(AccountKeyTypeSecondary),
	}
}

type AuthenticationMode string

const (
	AuthenticationModeAAD                     AuthenticationMode = "AAD"
	AuthenticationModeSharedKey               AuthenticationMode = "SharedKey"
	AuthenticationModeTaskAuthenticationToken AuthenticationMode = "TaskAuthenticationToken"
)

func PossibleValuesForAuthenticationMode() []string {
	return []string{
		string(AuthenticationModeAAD),
		string(AuthenticationModeSharedKey),
		string(AuthenticationModeTaskAuthenticationToken),
	}
}

type AutoStorageAuthenticationMode string

const (
	AutoStorageAuthenticationModeBatchAccountManagedIdentity AutoStorageAuthenticationMode = "BatchAccountManagedIdentity"
	AutoStorageAuthenticationModeStorageKeys                 AutoStorageAuthenticationMode = "StorageKeys"
)

func PossibleValuesForAutoStorageAuthenticationMode() []string {
	return []string{
		string(AutoStorageAuthenticationModeBatchAccountManagedIdentity),
		string(AutoStorageAuthenticationModeStorageKeys),
	}
}

type EndpointAccessDefaultAction string

const (
	EndpointAccessDefaultActionAllow EndpointAccessDefaultAction = "Allow"
	EndpointAccessDefaultActionDeny  EndpointAccessDefaultAction = "Deny"
)

func PossibleValuesForEndpointAccessDefaultAction() []string {
	return []string{
		string(EndpointAccessDefaultActionAllow),
		string(EndpointAccessDefaultActionDeny),
	}
}

type IPRuleAction string

const (
	IPRuleActionAllow IPRuleAction = "Allow"
)

func PossibleValuesForIPRuleAction() []string {
	return []string{
		string(IPRuleActionAllow),
	}
}

type KeySource string

const (
	KeySourceMicrosoftPointBatch    KeySource = "Microsoft.Batch"
	KeySourceMicrosoftPointKeyVault KeySource = "Microsoft.KeyVault"
)

func PossibleValuesForKeySource() []string {
	return []string{
		string(KeySourceMicrosoftPointBatch),
		string(KeySourceMicrosoftPointKeyVault),
	}
}

type PoolAllocationMode string

const (
	PoolAllocationModeBatchService     PoolAllocationMode = "BatchService"
	PoolAllocationModeUserSubscription PoolAllocationMode = "UserSubscription"
)

func PossibleValuesForPoolAllocationMode() []string {
	return []string{
		string(PoolAllocationModeBatchService),
		string(PoolAllocationModeUserSubscription),
	}
}

type PrivateEndpointConnectionProvisioningState string

const (
	PrivateEndpointConnectionProvisioningStateCancelled PrivateEndpointConnectionProvisioningState = "Cancelled"
	PrivateEndpointConnectionProvisioningStateCreating  PrivateEndpointConnectionProvisioningState = "Creating"
	PrivateEndpointConnectionProvisioningStateDeleting  PrivateEndpointConnectionProvisioningState = "Deleting"
	PrivateEndpointConnectionProvisioningStateFailed    PrivateEndpointConnectionProvisioningState = "Failed"
	PrivateEndpointConnectionProvisioningStateSucceeded PrivateEndpointConnectionProvisioningState = "Succeeded"
	PrivateEndpointConnectionProvisioningStateUpdating  PrivateEndpointConnectionProvisioningState = "Updating"
)

func PossibleValuesForPrivateEndpointConnectionProvisioningState() []string {
	return []string{
		string(PrivateEndpointConnectionProvisioningStateCancelled),
		string(PrivateEndpointConnectionProvisioningStateCreating),
		string(PrivateEndpointConnectionProvisioningStateDeleting),
		string(PrivateEndpointConnectionProvisioningStateFailed),
		string(PrivateEndpointConnectionProvisioningStateSucceeded),
		string(PrivateEndpointConnectionProvisioningStateUpdating),
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
	ProvisioningStateCancelled ProvisioningState = "Cancelled"
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateInvalid   ProvisioningState = "Invalid"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCancelled),
		string(ProvisioningStateCreating),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateInvalid),
		string(ProvisioningStateSucceeded),
	}
}

type PublicNetworkAccessType string

const (
	PublicNetworkAccessTypeDisabled PublicNetworkAccessType = "Disabled"
	PublicNetworkAccessTypeEnabled  PublicNetworkAccessType = "Enabled"
)

func PossibleValuesForPublicNetworkAccessType() []string {
	return []string{
		string(PublicNetworkAccessTypeDisabled),
		string(PublicNetworkAccessTypeEnabled),
	}
}
