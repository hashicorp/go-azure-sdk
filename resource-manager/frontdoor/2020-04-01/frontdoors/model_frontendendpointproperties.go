package frontdoors

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FrontendEndpointProperties struct {
	CustomHttpsConfiguration         *CustomHttpsConfiguration                                         `json:"customHttpsConfiguration,omitempty"`
	CustomHttpsProvisioningState     *CustomHttpsProvisioningState                                     `json:"customHttpsProvisioningState,omitempty"`
	CustomHttpsProvisioningSubstate  *CustomHttpsProvisioningSubstate                                  `json:"customHttpsProvisioningSubstate,omitempty"`
	HostName                         *string                                                           `json:"hostName,omitempty"`
	ResourceState                    *FrontDoorResourceState                                           `json:"resourceState,omitempty"`
	SessionAffinityEnabledState      *SessionAffinityEnabledState                                      `json:"sessionAffinityEnabledState,omitempty"`
	SessionAffinityTtlSeconds        *int64                                                            `json:"sessionAffinityTtlSeconds,omitempty"`
	WebApplicationFirewallPolicyLink *FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLink `json:"webApplicationFirewallPolicyLink,omitempty"`
}
