package apps

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppSku string

const (
	AppSkuSTOne  AppSku = "ST1"
	AppSkuSTTwo  AppSku = "ST2"
	AppSkuSTZero AppSku = "ST0"
)

func PossibleValuesForAppSku() []string {
	return []string{
		string(AppSkuSTOne),
		string(AppSkuSTTwo),
		string(AppSkuSTZero),
	}
}

type AppState string

const (
	AppStateCreated   AppState = "created"
	AppStateSuspended AppState = "suspended"
)

func PossibleValuesForAppState() []string {
	return []string{
		string(AppStateCreated),
		string(AppStateSuspended),
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

type NetworkAction string

const (
	NetworkActionAllow NetworkAction = "Allow"
	NetworkActionDeny  NetworkAction = "Deny"
)

func PossibleValuesForNetworkAction() []string {
	return []string{
		string(NetworkActionAllow),
		string(NetworkActionDeny),
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
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUpdating  ProvisioningState = "Updating"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCanceled),
		string(ProvisioningStateCreating),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateSucceeded),
		string(ProvisioningStateUpdating),
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
