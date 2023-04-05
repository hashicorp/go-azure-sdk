package namespacesnetworksecurityperimeterconfigurations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkSecurityPerimeterConfigurationProvisioningState string

const (
	NetworkSecurityPerimeterConfigurationProvisioningStateAccepted            NetworkSecurityPerimeterConfigurationProvisioningState = "Accepted"
	NetworkSecurityPerimeterConfigurationProvisioningStateCanceled            NetworkSecurityPerimeterConfigurationProvisioningState = "Canceled"
	NetworkSecurityPerimeterConfigurationProvisioningStateCreating            NetworkSecurityPerimeterConfigurationProvisioningState = "Creating"
	NetworkSecurityPerimeterConfigurationProvisioningStateDeleted             NetworkSecurityPerimeterConfigurationProvisioningState = "Deleted"
	NetworkSecurityPerimeterConfigurationProvisioningStateDeleting            NetworkSecurityPerimeterConfigurationProvisioningState = "Deleting"
	NetworkSecurityPerimeterConfigurationProvisioningStateFailed              NetworkSecurityPerimeterConfigurationProvisioningState = "Failed"
	NetworkSecurityPerimeterConfigurationProvisioningStateInvalidResponse     NetworkSecurityPerimeterConfigurationProvisioningState = "InvalidResponse"
	NetworkSecurityPerimeterConfigurationProvisioningStateSucceeded           NetworkSecurityPerimeterConfigurationProvisioningState = "Succeeded"
	NetworkSecurityPerimeterConfigurationProvisioningStateSucceededWithIssues NetworkSecurityPerimeterConfigurationProvisioningState = "SucceededWithIssues"
	NetworkSecurityPerimeterConfigurationProvisioningStateUnknown             NetworkSecurityPerimeterConfigurationProvisioningState = "Unknown"
	NetworkSecurityPerimeterConfigurationProvisioningStateUpdating            NetworkSecurityPerimeterConfigurationProvisioningState = "Updating"
)

func PossibleValuesForNetworkSecurityPerimeterConfigurationProvisioningState() []string {
	return []string{
		string(NetworkSecurityPerimeterConfigurationProvisioningStateAccepted),
		string(NetworkSecurityPerimeterConfigurationProvisioningStateCanceled),
		string(NetworkSecurityPerimeterConfigurationProvisioningStateCreating),
		string(NetworkSecurityPerimeterConfigurationProvisioningStateDeleted),
		string(NetworkSecurityPerimeterConfigurationProvisioningStateDeleting),
		string(NetworkSecurityPerimeterConfigurationProvisioningStateFailed),
		string(NetworkSecurityPerimeterConfigurationProvisioningStateInvalidResponse),
		string(NetworkSecurityPerimeterConfigurationProvisioningStateSucceeded),
		string(NetworkSecurityPerimeterConfigurationProvisioningStateSucceededWithIssues),
		string(NetworkSecurityPerimeterConfigurationProvisioningStateUnknown),
		string(NetworkSecurityPerimeterConfigurationProvisioningStateUpdating),
	}
}

type NspAccessRuleDirection string

const (
	NspAccessRuleDirectionInbound  NspAccessRuleDirection = "Inbound"
	NspAccessRuleDirectionOutbound NspAccessRuleDirection = "Outbound"
)

func PossibleValuesForNspAccessRuleDirection() []string {
	return []string{
		string(NspAccessRuleDirectionInbound),
		string(NspAccessRuleDirectionOutbound),
	}
}

type ResourceAssociationAccessMode string

const (
	ResourceAssociationAccessModeAuditMode         ResourceAssociationAccessMode = "AuditMode"
	ResourceAssociationAccessModeEnforcedMode      ResourceAssociationAccessMode = "EnforcedMode"
	ResourceAssociationAccessModeLearningMode      ResourceAssociationAccessMode = "LearningMode"
	ResourceAssociationAccessModeNoAssociationMode ResourceAssociationAccessMode = "NoAssociationMode"
	ResourceAssociationAccessModeUnspecifiedMode   ResourceAssociationAccessMode = "UnspecifiedMode"
)

func PossibleValuesForResourceAssociationAccessMode() []string {
	return []string{
		string(ResourceAssociationAccessModeAuditMode),
		string(ResourceAssociationAccessModeEnforcedMode),
		string(ResourceAssociationAccessModeLearningMode),
		string(ResourceAssociationAccessModeNoAssociationMode),
		string(ResourceAssociationAccessModeUnspecifiedMode),
	}
}
