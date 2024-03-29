package virtualmachinescalesets

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualMachineScaleSetUpdatePublicIPAddressConfigurationProperties struct {
	DeleteOption         *DeleteOptions                                                 `json:"deleteOption,omitempty"`
	DnsSettings          *VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettings `json:"dnsSettings,omitempty"`
	IdleTimeoutInMinutes *int64                                                         `json:"idleTimeoutInMinutes,omitempty"`
}
