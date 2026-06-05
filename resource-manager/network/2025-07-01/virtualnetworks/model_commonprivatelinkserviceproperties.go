package virtualnetworks

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommonPrivateLinkServiceProperties struct {
	AccessMode                           *AccessMode                                `json:"accessMode,omitempty"`
	Alias                                *string                                    `json:"alias,omitempty"`
	AutoApproval                         *CommonResourceSet                         `json:"autoApproval,omitempty"`
	DestinationIPAddress                 *string                                    `json:"destinationIPAddress,omitempty"`
	EnableProxyProtocol                  *bool                                      `json:"enableProxyProtocol,omitempty"`
	Fqdns                                *[]string                                  `json:"fqdns,omitempty"`
	IPConfigurations                     *[]CommonPrivateLinkServiceIPConfiguration `json:"ipConfigurations,omitempty"`
	LoadBalancerFrontendIPConfigurations *[]CommonFrontendIPConfiguration           `json:"loadBalancerFrontendIpConfigurations,omitempty"`
	NetworkInterfaces                    *[]CommonNetworkInterface                  `json:"networkInterfaces,omitempty"`
	PrivateEndpointConnections           *[]CommonPrivateEndpointConnection         `json:"privateEndpointConnections,omitempty"`
	ProvisioningState                    *ProvisioningState                         `json:"provisioningState,omitempty"`
	Visibility                           *CommonResourceSet                         `json:"visibility,omitempty"`
}
