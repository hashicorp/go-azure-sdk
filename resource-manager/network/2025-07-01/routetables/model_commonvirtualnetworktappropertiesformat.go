package routetables

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommonVirtualNetworkTapPropertiesFormat struct {
	DestinationLoadBalancerFrontEndIPConfiguration *CommonFrontendIPConfiguration            `json:"destinationLoadBalancerFrontEndIPConfiguration,omitempty"`
	DestinationNetworkInterfaceIPConfiguration     *CommonNetworkInterfaceIPConfiguration    `json:"destinationNetworkInterfaceIPConfiguration,omitempty"`
	DestinationPort                                *int64                                    `json:"destinationPort,omitempty"`
	NetworkInterfaceTapConfigurations              *[]CommonNetworkInterfaceTapConfiguration `json:"networkInterfaceTapConfigurations,omitempty"`
	ProvisioningState                              *ProvisioningState                        `json:"provisioningState,omitempty"`
	ResourceGuid                                   *string                                   `json:"resourceGuid,omitempty"`
}
