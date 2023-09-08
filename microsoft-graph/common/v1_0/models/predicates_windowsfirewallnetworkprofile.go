package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsFirewallNetworkProfileOperationPredicate struct {
	AuthorizedApplicationRulesFromGroupPolicyMerged *bool
	ConnectionSecurityRulesFromGroupPolicyMerged    *bool
	GlobalPortRulesFromGroupPolicyMerged            *bool
	InboundConnectionsBlocked                       *bool
	InboundNotificationsBlocked                     *bool
	IncomingTrafficBlocked                          *bool
	ODataType                                       *string
	OutboundConnectionsBlocked                      *bool
	PolicyRulesFromGroupPolicyMerged                *bool
	SecuredPacketExemptionAllowed                   *bool
	StealthModeBlocked                              *bool
	UnicastResponsesToMulticastBroadcastsBlocked    *bool
}

func (p WindowsFirewallNetworkProfileOperationPredicate) Matches(input WindowsFirewallNetworkProfile) bool {

	if p.AuthorizedApplicationRulesFromGroupPolicyMerged != nil && (input.AuthorizedApplicationRulesFromGroupPolicyMerged == nil || *p.AuthorizedApplicationRulesFromGroupPolicyMerged != *input.AuthorizedApplicationRulesFromGroupPolicyMerged) {
		return false
	}

	if p.ConnectionSecurityRulesFromGroupPolicyMerged != nil && (input.ConnectionSecurityRulesFromGroupPolicyMerged == nil || *p.ConnectionSecurityRulesFromGroupPolicyMerged != *input.ConnectionSecurityRulesFromGroupPolicyMerged) {
		return false
	}

	if p.GlobalPortRulesFromGroupPolicyMerged != nil && (input.GlobalPortRulesFromGroupPolicyMerged == nil || *p.GlobalPortRulesFromGroupPolicyMerged != *input.GlobalPortRulesFromGroupPolicyMerged) {
		return false
	}

	if p.InboundConnectionsBlocked != nil && (input.InboundConnectionsBlocked == nil || *p.InboundConnectionsBlocked != *input.InboundConnectionsBlocked) {
		return false
	}

	if p.InboundNotificationsBlocked != nil && (input.InboundNotificationsBlocked == nil || *p.InboundNotificationsBlocked != *input.InboundNotificationsBlocked) {
		return false
	}

	if p.IncomingTrafficBlocked != nil && (input.IncomingTrafficBlocked == nil || *p.IncomingTrafficBlocked != *input.IncomingTrafficBlocked) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OutboundConnectionsBlocked != nil && (input.OutboundConnectionsBlocked == nil || *p.OutboundConnectionsBlocked != *input.OutboundConnectionsBlocked) {
		return false
	}

	if p.PolicyRulesFromGroupPolicyMerged != nil && (input.PolicyRulesFromGroupPolicyMerged == nil || *p.PolicyRulesFromGroupPolicyMerged != *input.PolicyRulesFromGroupPolicyMerged) {
		return false
	}

	if p.SecuredPacketExemptionAllowed != nil && (input.SecuredPacketExemptionAllowed == nil || *p.SecuredPacketExemptionAllowed != *input.SecuredPacketExemptionAllowed) {
		return false
	}

	if p.StealthModeBlocked != nil && (input.StealthModeBlocked == nil || *p.StealthModeBlocked != *input.StealthModeBlocked) {
		return false
	}

	if p.UnicastResponsesToMulticastBroadcastsBlocked != nil && (input.UnicastResponsesToMulticastBroadcastsBlocked == nil || *p.UnicastResponsesToMulticastBroadcastsBlocked != *input.UnicastResponsesToMulticastBroadcastsBlocked) {
		return false
	}

	return true
}
