package connectedclusters

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationMethod string

const (
	AuthenticationMethodAAD   AuthenticationMethod = "AAD"
	AuthenticationMethodToken AuthenticationMethod = "Token"
)

func PossibleValuesForAuthenticationMethod() []string {
	return []string{
		string(AuthenticationMethodAAD),
		string(AuthenticationMethodToken),
	}
}

type ConnectivityStatus string

const (
	ConnectivityStatusConnected  ConnectivityStatus = "Connected"
	ConnectivityStatusConnecting ConnectivityStatus = "Connecting"
	ConnectivityStatusExpired    ConnectivityStatus = "Expired"
	ConnectivityStatusOffline    ConnectivityStatus = "Offline"
)

func PossibleValuesForConnectivityStatus() []string {
	return []string{
		string(ConnectivityStatusConnected),
		string(ConnectivityStatusConnecting),
		string(ConnectivityStatusExpired),
		string(ConnectivityStatusOffline),
	}
}

type ProvisioningState string

const (
	ProvisioningStateAccepted     ProvisioningState = "Accepted"
	ProvisioningStateCanceled     ProvisioningState = "Canceled"
	ProvisioningStateDeleting     ProvisioningState = "Deleting"
	ProvisioningStateFailed       ProvisioningState = "Failed"
	ProvisioningStateProvisioning ProvisioningState = "Provisioning"
	ProvisioningStateSucceeded    ProvisioningState = "Succeeded"
	ProvisioningStateUpdating     ProvisioningState = "Updating"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateAccepted),
		string(ProvisioningStateCanceled),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateProvisioning),
		string(ProvisioningStateSucceeded),
		string(ProvisioningStateUpdating),
	}
}
