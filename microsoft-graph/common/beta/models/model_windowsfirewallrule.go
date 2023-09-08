package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsFirewallRule struct {
	Action                  *WindowsFirewallRuleAction           `json:"action,omitempty"`
	Description             *string                              `json:"description,omitempty"`
	DisplayName             *string                              `json:"displayName,omitempty"`
	EdgeTraversal           *WindowsFirewallRuleEdgeTraversal    `json:"edgeTraversal,omitempty"`
	FilePath                *string                              `json:"filePath,omitempty"`
	InterfaceTypes          *WindowsFirewallRuleInterfaceTypes   `json:"interfaceTypes,omitempty"`
	LocalAddressRanges      *[]string                            `json:"localAddressRanges,omitempty"`
	LocalPortRanges         *[]string                            `json:"localPortRanges,omitempty"`
	LocalUserAuthorizations *string                              `json:"localUserAuthorizations,omitempty"`
	ODataType               *string                              `json:"@odata.type,omitempty"`
	PackageFamilyName       *string                              `json:"packageFamilyName,omitempty"`
	ProfileTypes            *WindowsFirewallRuleProfileTypes     `json:"profileTypes,omitempty"`
	Protocol                *int64                               `json:"protocol,omitempty"`
	RemoteAddressRanges     *[]string                            `json:"remoteAddressRanges,omitempty"`
	RemotePortRanges        *[]string                            `json:"remotePortRanges,omitempty"`
	ServiceName             *string                              `json:"serviceName,omitempty"`
	TrafficDirection        *WindowsFirewallRuleTrafficDirection `json:"trafficDirection,omitempty"`
}
