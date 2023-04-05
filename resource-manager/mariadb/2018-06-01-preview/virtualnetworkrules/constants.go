package virtualnetworkrules

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualNetworkRuleState string

const (
	VirtualNetworkRuleStateDeleting     VirtualNetworkRuleState = "Deleting"
	VirtualNetworkRuleStateInProgress   VirtualNetworkRuleState = "InProgress"
	VirtualNetworkRuleStateInitializing VirtualNetworkRuleState = "Initializing"
	VirtualNetworkRuleStateReady        VirtualNetworkRuleState = "Ready"
	VirtualNetworkRuleStateUnknown      VirtualNetworkRuleState = "Unknown"
)

func PossibleValuesForVirtualNetworkRuleState() []string {
	return []string{
		string(VirtualNetworkRuleStateDeleting),
		string(VirtualNetworkRuleStateInProgress),
		string(VirtualNetworkRuleStateInitializing),
		string(VirtualNetworkRuleStateReady),
		string(VirtualNetworkRuleStateUnknown),
	}
}
