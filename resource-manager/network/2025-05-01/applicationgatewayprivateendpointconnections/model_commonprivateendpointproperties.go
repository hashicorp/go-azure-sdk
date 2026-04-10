package applicationgatewayprivateendpointconnections

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommonPrivateEndpointProperties struct {
	ApplicationSecurityGroups           *[]CommonApplicationSecurityGroup        `json:"applicationSecurityGroups,omitempty"`
	CustomDnsConfigs                    *[]CommonCustomDnsConfigPropertiesFormat `json:"customDnsConfigs,omitempty"`
	CustomNetworkInterfaceName          *string                                  `json:"customNetworkInterfaceName,omitempty"`
	IPConfigurations                    *[]CommonPrivateEndpointIPConfiguration  `json:"ipConfigurations,omitempty"`
	IPVersionType                       *PrivateEndpointIPVersionType            `json:"ipVersionType,omitempty"`
	ManualPrivateLinkServiceConnections *[]CommonPrivateLinkServiceConnection    `json:"manualPrivateLinkServiceConnections,omitempty"`
	NetworkInterfaces                   *[]CommonNetworkInterface                `json:"networkInterfaces,omitempty"`
	PrivateLinkServiceConnections       *[]CommonPrivateLinkServiceConnection    `json:"privateLinkServiceConnections,omitempty"`
	ProvisioningState                   *ProvisioningState                       `json:"provisioningState,omitempty"`
	Subnet                              *CommonSubnet                            `json:"subnet,omitempty"`
}
