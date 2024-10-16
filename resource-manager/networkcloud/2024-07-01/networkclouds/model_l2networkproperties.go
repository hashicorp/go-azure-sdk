package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type L2NetworkProperties struct {
	AssociatedResourceIds          *[]string                   `json:"associatedResourceIds,omitempty"`
	ClusterId                      *string                     `json:"clusterId,omitempty"`
	DetailedStatus                 *L2NetworkDetailedStatus    `json:"detailedStatus,omitempty"`
	DetailedStatusMessage          *string                     `json:"detailedStatusMessage,omitempty"`
	HybridAksClustersAssociatedIds *[]string                   `json:"hybridAksClustersAssociatedIds,omitempty"`
	HybridAksPluginType            *HybridAksPluginType        `json:"hybridAksPluginType,omitempty"`
	InterfaceName                  *string                     `json:"interfaceName,omitempty"`
	L2IsolationDomainId            string                      `json:"l2IsolationDomainId"`
	ProvisioningState              *L2NetworkProvisioningState `json:"provisioningState,omitempty"`
	VirtualMachinesAssociatedIds   *[]string                   `json:"virtualMachinesAssociatedIds,omitempty"`
}
