package cluster

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClusterNodeType string

const (
	ClusterNodeTypeFirstParty ClusterNodeType = "FirstParty"
	ClusterNodeTypeThirdParty ClusterNodeType = "ThirdParty"
)

func PossibleValuesForClusterNodeType() []string {
	return []string{
		string(ClusterNodeTypeFirstParty),
		string(ClusterNodeTypeThirdParty),
	}
}

type DiagnosticLevel string

const (
	DiagnosticLevelBasic    DiagnosticLevel = "Basic"
	DiagnosticLevelEnhanced DiagnosticLevel = "Enhanced"
	DiagnosticLevelOff      DiagnosticLevel = "Off"
)

func PossibleValuesForDiagnosticLevel() []string {
	return []string{
		string(DiagnosticLevelBasic),
		string(DiagnosticLevelEnhanced),
		string(DiagnosticLevelOff),
	}
}

type ImdsAttestation string

const (
	ImdsAttestationDisabled ImdsAttestation = "Disabled"
	ImdsAttestationEnabled  ImdsAttestation = "Enabled"
)

func PossibleValuesForImdsAttestation() []string {
	return []string{
		string(ImdsAttestationDisabled),
		string(ImdsAttestationEnabled),
	}
}

type ProvisioningState string

const (
	ProvisioningStateAccepted           ProvisioningState = "Accepted"
	ProvisioningStateCanceled           ProvisioningState = "Canceled"
	ProvisioningStateConnected          ProvisioningState = "Connected"
	ProvisioningStateCreating           ProvisioningState = "Creating"
	ProvisioningStateDeleted            ProvisioningState = "Deleted"
	ProvisioningStateDeleting           ProvisioningState = "Deleting"
	ProvisioningStateDisableInProgress  ProvisioningState = "DisableInProgress"
	ProvisioningStateDisconnected       ProvisioningState = "Disconnected"
	ProvisioningStateFailed             ProvisioningState = "Failed"
	ProvisioningStateInProgress         ProvisioningState = "InProgress"
	ProvisioningStateMoving             ProvisioningState = "Moving"
	ProvisioningStateNotSpecified       ProvisioningState = "NotSpecified"
	ProvisioningStatePartiallyConnected ProvisioningState = "PartiallyConnected"
	ProvisioningStatePartiallySucceeded ProvisioningState = "PartiallySucceeded"
	ProvisioningStateProvisioning       ProvisioningState = "Provisioning"
	ProvisioningStateSucceeded          ProvisioningState = "Succeeded"
	ProvisioningStateUpdating           ProvisioningState = "Updating"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateAccepted),
		string(ProvisioningStateCanceled),
		string(ProvisioningStateConnected),
		string(ProvisioningStateCreating),
		string(ProvisioningStateDeleted),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateDisableInProgress),
		string(ProvisioningStateDisconnected),
		string(ProvisioningStateFailed),
		string(ProvisioningStateInProgress),
		string(ProvisioningStateMoving),
		string(ProvisioningStateNotSpecified),
		string(ProvisioningStatePartiallyConnected),
		string(ProvisioningStatePartiallySucceeded),
		string(ProvisioningStateProvisioning),
		string(ProvisioningStateSucceeded),
		string(ProvisioningStateUpdating),
	}
}

type SoftwareAssuranceIntent string

const (
	SoftwareAssuranceIntentDisable SoftwareAssuranceIntent = "Disable"
	SoftwareAssuranceIntentEnable  SoftwareAssuranceIntent = "Enable"
)

func PossibleValuesForSoftwareAssuranceIntent() []string {
	return []string{
		string(SoftwareAssuranceIntentDisable),
		string(SoftwareAssuranceIntentEnable),
	}
}

type SoftwareAssuranceStatus string

const (
	SoftwareAssuranceStatusDisabled SoftwareAssuranceStatus = "Disabled"
	SoftwareAssuranceStatusEnabled  SoftwareAssuranceStatus = "Enabled"
)

func PossibleValuesForSoftwareAssuranceStatus() []string {
	return []string{
		string(SoftwareAssuranceStatusDisabled),
		string(SoftwareAssuranceStatusEnabled),
	}
}

type Status string

const (
	StatusConnectedRecently    Status = "ConnectedRecently"
	StatusDisconnected         Status = "Disconnected"
	StatusError                Status = "Error"
	StatusNotConnectedRecently Status = "NotConnectedRecently"
	StatusNotSpecified         Status = "NotSpecified"
	StatusNotYetRegistered     Status = "NotYetRegistered"
)

func PossibleValuesForStatus() []string {
	return []string{
		string(StatusConnectedRecently),
		string(StatusDisconnected),
		string(StatusError),
		string(StatusNotConnectedRecently),
		string(StatusNotSpecified),
		string(StatusNotYetRegistered),
	}
}

type WindowsServerSubscription string

const (
	WindowsServerSubscriptionDisabled WindowsServerSubscription = "Disabled"
	WindowsServerSubscriptionEnabled  WindowsServerSubscription = "Enabled"
)

func PossibleValuesForWindowsServerSubscription() []string {
	return []string{
		string(WindowsServerSubscriptionDisabled),
		string(WindowsServerSubscriptionEnabled),
	}
}
