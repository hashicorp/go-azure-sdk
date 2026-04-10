package webapplicationfirewallpolicies

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationGatewayRoutingRulePropertiesFormat struct {
	BackendAddressPool *CommonSubResource                        `json:"backendAddressPool,omitempty"`
	BackendSettings    *CommonSubResource                        `json:"backendSettings,omitempty"`
	Listener           *CommonSubResource                        `json:"listener,omitempty"`
	Priority           int64                                     `json:"priority"`
	ProvisioningState  *ProvisioningState                        `json:"provisioningState,omitempty"`
	RuleType           *ApplicationGatewayRequestRoutingRuleType `json:"ruleType,omitempty"`
}
