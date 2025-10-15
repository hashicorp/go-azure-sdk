package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TrunkedNetworkProperties struct {
	AssociatedResourceIds          *[]string                        `json:"associatedResourceIds,omitempty"`
	ClusterId                      *string                          `json:"clusterId,omitempty"`
	DetailedStatus                 *TrunkedNetworkDetailedStatus    `json:"detailedStatus,omitempty"`
	DetailedStatusMessage          *string                          `json:"detailedStatusMessage,omitempty"`
	HybridAksClustersAssociatedIds *[]string                        `json:"hybridAksClustersAssociatedIds,omitempty"`
	HybridAksPluginType            *HybridAksPluginType             `json:"hybridAksPluginType,omitempty"`
	InterfaceName                  *string                          `json:"interfaceName,omitempty"`
	IsolationDomainIds             []string                         `json:"isolationDomainIds"`
	ProvisioningState              *TrunkedNetworkProvisioningState `json:"provisioningState,omitempty"`
	VirtualMachinesAssociatedIds   *[]string                        `json:"virtualMachinesAssociatedIds,omitempty"`
	Vlans                          []int64                          `json:"vlans"`
}
