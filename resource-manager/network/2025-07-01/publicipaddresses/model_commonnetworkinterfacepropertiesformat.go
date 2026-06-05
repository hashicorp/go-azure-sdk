package publicipaddresses

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommonNetworkInterfacePropertiesFormat struct {
	AuxiliaryMode                      *NetworkInterfaceAuxiliaryMode            `json:"auxiliaryMode,omitempty"`
	AuxiliarySku                       *NetworkInterfaceAuxiliarySku             `json:"auxiliarySku,omitempty"`
	DefaultOutboundConnectivityEnabled *bool                                     `json:"defaultOutboundConnectivityEnabled,omitempty"`
	DisableTcpStateTracking            *bool                                     `json:"disableTcpStateTracking,omitempty"`
	DnsSettings                        *CommonNetworkInterfaceDnsSettings        `json:"dnsSettings,omitempty"`
	DscpConfiguration                  *CommonSubResource                        `json:"dscpConfiguration,omitempty"`
	EnableAcceleratedNetworking        *bool                                     `json:"enableAcceleratedNetworking,omitempty"`
	EnableIPForwarding                 *bool                                     `json:"enableIPForwarding,omitempty"`
	HostedWorkloads                    *[]string                                 `json:"hostedWorkloads,omitempty"`
	IPConfigurations                   *[]CommonNetworkInterfaceIPConfiguration  `json:"ipConfigurations,omitempty"`
	MacAddress                         *string                                   `json:"macAddress,omitempty"`
	MigrationPhase                     *NetworkInterfaceMigrationPhase           `json:"migrationPhase,omitempty"`
	NetworkSecurityGroup               *CommonNetworkSecurityGroup               `json:"networkSecurityGroup,omitempty"`
	NicType                            *NetworkInterfaceNicType                  `json:"nicType,omitempty"`
	Primary                            *bool                                     `json:"primary,omitempty"`
	PrivateEndpoint                    *CommonPrivateEndpoint                    `json:"privateEndpoint,omitempty"`
	PrivateLinkService                 *CommonPrivateLinkService                 `json:"privateLinkService,omitempty"`
	ProvisioningState                  *ProvisioningState                        `json:"provisioningState,omitempty"`
	ResourceGuid                       *string                                   `json:"resourceGuid,omitempty"`
	TapConfigurations                  *[]CommonNetworkInterfaceTapConfiguration `json:"tapConfigurations,omitempty"`
	VirtualMachine                     *CommonSubResource                        `json:"virtualMachine,omitempty"`
	VnetEncryptionSupported            *bool                                     `json:"vnetEncryptionSupported,omitempty"`
	WorkloadType                       *string                                   `json:"workloadType,omitempty"`
}
