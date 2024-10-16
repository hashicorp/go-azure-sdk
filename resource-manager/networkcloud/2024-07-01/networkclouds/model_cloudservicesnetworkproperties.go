package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudServicesNetworkProperties struct {
	AdditionalEgressEndpoints      *[]EgressEndpoint                                 `json:"additionalEgressEndpoints,omitempty"`
	AssociatedResourceIds          *[]string                                         `json:"associatedResourceIds,omitempty"`
	ClusterId                      *string                                           `json:"clusterId,omitempty"`
	DetailedStatus                 *CloudServicesNetworkDetailedStatus               `json:"detailedStatus,omitempty"`
	DetailedStatusMessage          *string                                           `json:"detailedStatusMessage,omitempty"`
	EnableDefaultEgressEndpoints   *CloudServicesNetworkEnableDefaultEgressEndpoints `json:"enableDefaultEgressEndpoints,omitempty"`
	EnabledEgressEndpoints         *[]EgressEndpoint                                 `json:"enabledEgressEndpoints,omitempty"`
	HybridAksClustersAssociatedIds *[]string                                         `json:"hybridAksClustersAssociatedIds,omitempty"`
	InterfaceName                  *string                                           `json:"interfaceName,omitempty"`
	ProvisioningState              *CloudServicesNetworkProvisioningState            `json:"provisioningState,omitempty"`
	VirtualMachinesAssociatedIds   *[]string                                         `json:"virtualMachinesAssociatedIds,omitempty"`
}
