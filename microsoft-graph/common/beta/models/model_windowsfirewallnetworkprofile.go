package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsFirewallNetworkProfile struct {
	AuthorizedApplicationRulesFromGroupPolicyMerged    *bool                                         `json:"authorizedApplicationRulesFromGroupPolicyMerged,omitempty"`
	AuthorizedApplicationRulesFromGroupPolicyNotMerged *bool                                         `json:"authorizedApplicationRulesFromGroupPolicyNotMerged,omitempty"`
	ConnectionSecurityRulesFromGroupPolicyMerged       *bool                                         `json:"connectionSecurityRulesFromGroupPolicyMerged,omitempty"`
	ConnectionSecurityRulesFromGroupPolicyNotMerged    *bool                                         `json:"connectionSecurityRulesFromGroupPolicyNotMerged,omitempty"`
	FirewallEnabled                                    *WindowsFirewallNetworkProfileFirewallEnabled `json:"firewallEnabled,omitempty"`
	GlobalPortRulesFromGroupPolicyMerged               *bool                                         `json:"globalPortRulesFromGroupPolicyMerged,omitempty"`
	GlobalPortRulesFromGroupPolicyNotMerged            *bool                                         `json:"globalPortRulesFromGroupPolicyNotMerged,omitempty"`
	InboundConnectionsBlocked                          *bool                                         `json:"inboundConnectionsBlocked,omitempty"`
	InboundConnectionsRequired                         *bool                                         `json:"inboundConnectionsRequired,omitempty"`
	InboundNotificationsBlocked                        *bool                                         `json:"inboundNotificationsBlocked,omitempty"`
	InboundNotificationsRequired                       *bool                                         `json:"inboundNotificationsRequired,omitempty"`
	IncomingTrafficBlocked                             *bool                                         `json:"incomingTrafficBlocked,omitempty"`
	IncomingTrafficRequired                            *bool                                         `json:"incomingTrafficRequired,omitempty"`
	ODataType                                          *string                                       `json:"@odata.type,omitempty"`
	OutboundConnectionsBlocked                         *bool                                         `json:"outboundConnectionsBlocked,omitempty"`
	OutboundConnectionsRequired                        *bool                                         `json:"outboundConnectionsRequired,omitempty"`
	PolicyRulesFromGroupPolicyMerged                   *bool                                         `json:"policyRulesFromGroupPolicyMerged,omitempty"`
	PolicyRulesFromGroupPolicyNotMerged                *bool                                         `json:"policyRulesFromGroupPolicyNotMerged,omitempty"`
	SecuredPacketExemptionAllowed                      *bool                                         `json:"securedPacketExemptionAllowed,omitempty"`
	SecuredPacketExemptionBlocked                      *bool                                         `json:"securedPacketExemptionBlocked,omitempty"`
	StealthModeBlocked                                 *bool                                         `json:"stealthModeBlocked,omitempty"`
	StealthModeRequired                                *bool                                         `json:"stealthModeRequired,omitempty"`
	UnicastResponsesToMulticastBroadcastsBlocked       *bool                                         `json:"unicastResponsesToMulticastBroadcastsBlocked,omitempty"`
	UnicastResponsesToMulticastBroadcastsRequired      *bool                                         `json:"unicastResponsesToMulticastBroadcastsRequired,omitempty"`
}
