package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsFirewallNetworkProfile struct {
	AuthorizedApplicationRulesFromGroupPolicyMerged *bool                                         `json:"authorizedApplicationRulesFromGroupPolicyMerged,omitempty"`
	ConnectionSecurityRulesFromGroupPolicyMerged    *bool                                         `json:"connectionSecurityRulesFromGroupPolicyMerged,omitempty"`
	FirewallEnabled                                 *WindowsFirewallNetworkProfileFirewallEnabled `json:"firewallEnabled,omitempty"`
	GlobalPortRulesFromGroupPolicyMerged            *bool                                         `json:"globalPortRulesFromGroupPolicyMerged,omitempty"`
	InboundConnectionsBlocked                       *bool                                         `json:"inboundConnectionsBlocked,omitempty"`
	InboundNotificationsBlocked                     *bool                                         `json:"inboundNotificationsBlocked,omitempty"`
	IncomingTrafficBlocked                          *bool                                         `json:"incomingTrafficBlocked,omitempty"`
	ODataType                                       *string                                       `json:"@odata.type,omitempty"`
	OutboundConnectionsBlocked                      *bool                                         `json:"outboundConnectionsBlocked,omitempty"`
	PolicyRulesFromGroupPolicyMerged                *bool                                         `json:"policyRulesFromGroupPolicyMerged,omitempty"`
	SecuredPacketExemptionAllowed                   *bool                                         `json:"securedPacketExemptionAllowed,omitempty"`
	StealthModeBlocked                              *bool                                         `json:"stealthModeBlocked,omitempty"`
	UnicastResponsesToMulticastBroadcastsBlocked    *bool                                         `json:"unicastResponsesToMulticastBroadcastsBlocked,omitempty"`
}
