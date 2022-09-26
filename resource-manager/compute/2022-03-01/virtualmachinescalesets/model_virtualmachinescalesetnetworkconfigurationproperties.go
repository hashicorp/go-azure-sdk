package virtualmachinescalesets

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualMachineScaleSetNetworkConfigurationProperties struct {
	DeleteOption                *DeleteOptions                                         `json:"deleteOption,omitempty"`
	DnsSettings                 *VirtualMachineScaleSetNetworkConfigurationDnsSettings `json:"dnsSettings,omitempty"`
	EnableAcceleratedNetworking *bool                                                  `json:"enableAcceleratedNetworking,omitempty"`
	EnableFpga                  *bool                                                  `json:"enableFpga,omitempty"`
	EnableIPForwarding          *bool                                                  `json:"enableIPForwarding,omitempty"`
	IPConfigurations            []VirtualMachineScaleSetIPConfiguration                `json:"ipConfigurations"`
	NetworkSecurityGroup        *SubResource                                           `json:"networkSecurityGroup,omitempty"`
	Primary                     *bool                                                  `json:"primary,omitempty"`
}
