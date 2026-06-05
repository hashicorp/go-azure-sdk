package networkinterfaces

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommonLoadBalancingRulePropertiesFormat struct {
	BackendAddressPool       *CommonSubResource   `json:"backendAddressPool,omitempty"`
	BackendAddressPools      *[]CommonSubResource `json:"backendAddressPools,omitempty"`
	BackendPort              *int64               `json:"backendPort,omitempty"`
	DisableOutboundSnat      *bool                `json:"disableOutboundSnat,omitempty"`
	EnableConnectionTracking *bool                `json:"enableConnectionTracking,omitempty"`
	EnableFloatingIP         *bool                `json:"enableFloatingIP,omitempty"`
	EnableTcpReset           *bool                `json:"enableTcpReset,omitempty"`
	FrontendIPConfiguration  *CommonSubResource   `json:"frontendIPConfiguration,omitempty"`
	FrontendPort             int64                `json:"frontendPort"`
	IdleTimeoutInMinutes     *int64               `json:"idleTimeoutInMinutes,omitempty"`
	LoadDistribution         *LoadDistribution    `json:"loadDistribution,omitempty"`
	Probe                    *CommonSubResource   `json:"probe,omitempty"`
	Protocol                 TransportProtocol    `json:"protocol"`
	ProvisioningState        *ProvisioningState   `json:"provisioningState,omitempty"`
}
