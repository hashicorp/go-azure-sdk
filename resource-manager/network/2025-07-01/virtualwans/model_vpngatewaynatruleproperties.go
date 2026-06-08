package virtualwans

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VpnGatewayNatRuleProperties struct {
	EgressVpnSiteLinkConnections  *[]CommonSubResource `json:"egressVpnSiteLinkConnections,omitempty"`
	ExternalMappings              *[]VpnNatRuleMapping `json:"externalMappings,omitempty"`
	IPConfigurationId             *string              `json:"ipConfigurationId,omitempty"`
	IngressVpnSiteLinkConnections *[]CommonSubResource `json:"ingressVpnSiteLinkConnections,omitempty"`
	InternalMappings              *[]VpnNatRuleMapping `json:"internalMappings,omitempty"`
	Mode                          *VpnNatRuleMode      `json:"mode,omitempty"`
	ProvisioningState             *ProvisioningState   `json:"provisioningState,omitempty"`
	Type                          *VpnNatRuleType      `json:"type,omitempty"`
}
