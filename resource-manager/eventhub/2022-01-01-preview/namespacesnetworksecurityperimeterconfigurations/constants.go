package namespacesnetworksecurityperimeterconfigurations

import "strings"

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

func parseNetworkSecurityPerimeterConfigurationProvisioningState(input string) (*NetworkSecurityPerimeterConfigurationProvisioningState, error) {
	vals := map[string]NetworkSecurityPerimeterConfigurationProvisioningState{
		"accepted":            NetworkSecurityPerimeterConfigurationProvisioningStateAccepted,
		"canceled":            NetworkSecurityPerimeterConfigurationProvisioningStateCanceled,
		"creating":            NetworkSecurityPerimeterConfigurationProvisioningStateCreating,
		"deleted":             NetworkSecurityPerimeterConfigurationProvisioningStateDeleted,
		"deleting":            NetworkSecurityPerimeterConfigurationProvisioningStateDeleting,
		"failed":              NetworkSecurityPerimeterConfigurationProvisioningStateFailed,
		"invalidresponse":     NetworkSecurityPerimeterConfigurationProvisioningStateInvalidResponse,
		"succeeded":           NetworkSecurityPerimeterConfigurationProvisioningStateSucceeded,
		"succeededwithissues": NetworkSecurityPerimeterConfigurationProvisioningStateSucceededWithIssues,
		"unknown":             NetworkSecurityPerimeterConfigurationProvisioningStateUnknown,
		"updating":            NetworkSecurityPerimeterConfigurationProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkSecurityPerimeterConfigurationProvisioningState(input)
	return &out, nil
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

func parseNspAccessRuleDirection(input string) (*NspAccessRuleDirection, error) {
	vals := map[string]NspAccessRuleDirection{
		"inbound":  NspAccessRuleDirectionInbound,
		"outbound": NspAccessRuleDirectionOutbound,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NspAccessRuleDirection(input)
	return &out, nil
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

func parseResourceAssociationAccessMode(input string) (*ResourceAssociationAccessMode, error) {
	vals := map[string]ResourceAssociationAccessMode{
		"auditmode":         ResourceAssociationAccessModeAuditMode,
		"enforcedmode":      ResourceAssociationAccessModeEnforcedMode,
		"learningmode":      ResourceAssociationAccessModeLearningMode,
		"noassociationmode": ResourceAssociationAccessModeNoAssociationMode,
		"unspecifiedmode":   ResourceAssociationAccessModeUnspecifiedMode,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ResourceAssociationAccessMode(input)
	return &out, nil
}
