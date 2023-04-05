package managedhsms

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

type CreateMode string

const (
	CreateModeDefault CreateMode = "default"
	CreateModeRecover CreateMode = "recover"
)

func PossibleValuesForCreateMode() []string {
	return []string{
		string(CreateModeDefault),
		string(CreateModeRecover),
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

type NetworkRuleAction string

const (
	NetworkRuleActionAllow NetworkRuleAction = "Allow"
	NetworkRuleActionDeny  NetworkRuleAction = "Deny"
)

func PossibleValuesForNetworkRuleAction() []string {
	return []string{
		string(NetworkRuleActionAllow),
		string(NetworkRuleActionDeny),
	}
}

type NetworkRuleBypassOptions string

const (
	NetworkRuleBypassOptionsAzureServices NetworkRuleBypassOptions = "AzureServices"
	NetworkRuleBypassOptionsNone          NetworkRuleBypassOptions = "None"
)

func PossibleValuesForNetworkRuleBypassOptions() []string {
	return []string{
		string(NetworkRuleBypassOptionsAzureServices),
		string(NetworkRuleBypassOptionsNone),
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

type ProvisioningState string

const (
	ProvisioningStateActivated             ProvisioningState = "Activated"
	ProvisioningStateDeleting              ProvisioningState = "Deleting"
	ProvisioningStateFailed                ProvisioningState = "Failed"
	ProvisioningStateProvisioning          ProvisioningState = "Provisioning"
	ProvisioningStateRestoring             ProvisioningState = "Restoring"
	ProvisioningStateSecurityDomainRestore ProvisioningState = "SecurityDomainRestore"
	ProvisioningStateSucceeded             ProvisioningState = "Succeeded"
	ProvisioningStateUpdating              ProvisioningState = "Updating"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateActivated),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateProvisioning),
		string(ProvisioningStateRestoring),
		string(ProvisioningStateSecurityDomainRestore),
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
