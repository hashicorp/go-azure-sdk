package networksecuritygroups

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommonInboundNatRulePropertiesFormat struct {
	BackendAddressPool      *CommonSubResource                     `json:"backendAddressPool,omitempty"`
	BackendIPConfiguration  *CommonNetworkInterfaceIPConfiguration `json:"backendIPConfiguration,omitempty"`
	BackendPort             *int64                                 `json:"backendPort,omitempty"`
	EnableFloatingIP        *bool                                  `json:"enableFloatingIP,omitempty"`
	EnableTcpReset          *bool                                  `json:"enableTcpReset,omitempty"`
	FrontendIPConfiguration *CommonSubResource                     `json:"frontendIPConfiguration,omitempty"`
	FrontendPort            *int64                                 `json:"frontendPort,omitempty"`
	FrontendPortRangeEnd    *int64                                 `json:"frontendPortRangeEnd,omitempty"`
	FrontendPortRangeStart  *int64                                 `json:"frontendPortRangeStart,omitempty"`
	IdleTimeoutInMinutes    *int64                                 `json:"idleTimeoutInMinutes,omitempty"`
	Protocol                *TransportProtocol                     `json:"protocol,omitempty"`
	ProvisioningState       *ProvisioningState                     `json:"provisioningState,omitempty"`
}
