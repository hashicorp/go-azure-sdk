package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type L3NetworkProperties struct {
	AssociatedResourceIds          *[]string                   `json:"associatedResourceIds,omitempty"`
	ClusterId                      *string                     `json:"clusterId,omitempty"`
	DetailedStatus                 *L3NetworkDetailedStatus    `json:"detailedStatus,omitempty"`
	DetailedStatusMessage          *string                     `json:"detailedStatusMessage,omitempty"`
	HybridAksClustersAssociatedIds *[]string                   `json:"hybridAksClustersAssociatedIds,omitempty"`
	HybridAksIPamEnabled           *HybridAksIPamEnabled       `json:"hybridAksIpamEnabled,omitempty"`
	HybridAksPluginType            *HybridAksPluginType        `json:"hybridAksPluginType,omitempty"`
	IPAllocationType               *IPAllocationType           `json:"ipAllocationType,omitempty"`
	IPv4ConnectedPrefix            *string                     `json:"ipv4ConnectedPrefix,omitempty"`
	IPv6ConnectedPrefix            *string                     `json:"ipv6ConnectedPrefix,omitempty"`
	InterfaceName                  *string                     `json:"interfaceName,omitempty"`
	L3IsolationDomainId            string                      `json:"l3IsolationDomainId"`
	ProvisioningState              *L3NetworkProvisioningState `json:"provisioningState,omitempty"`
	VirtualMachinesAssociatedIds   *[]string                   `json:"virtualMachinesAssociatedIds,omitempty"`
	Vlan                           int64                       `json:"vlan"`
}
