package networkinterfaces

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommonOutboundRulePropertiesFormat struct {
	AllocatedOutboundPorts   *int64                           `json:"allocatedOutboundPorts,omitempty"`
	BackendAddressPool       CommonSubResource                `json:"backendAddressPool"`
	EnableTcpReset           *bool                            `json:"enableTcpReset,omitempty"`
	FrontendIPConfigurations []CommonSubResource              `json:"frontendIPConfigurations"`
	IdleTimeoutInMinutes     *int64                           `json:"idleTimeoutInMinutes,omitempty"`
	Protocol                 LoadBalancerOutboundRuleProtocol `json:"protocol"`
	ProvisioningState        *ProvisioningState               `json:"provisioningState,omitempty"`
}
