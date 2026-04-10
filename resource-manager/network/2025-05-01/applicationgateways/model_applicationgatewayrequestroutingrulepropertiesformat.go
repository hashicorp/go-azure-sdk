package applicationgateways

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationGatewayRequestRoutingRulePropertiesFormat struct {
	BackendAddressPool       *CommonSubResource                        `json:"backendAddressPool,omitempty"`
	BackendHTTPSettings      *CommonSubResource                        `json:"backendHttpSettings,omitempty"`
	EntraJWTValidationConfig *CommonSubResource                        `json:"entraJWTValidationConfig,omitempty"`
	HTTPListener             *CommonSubResource                        `json:"httpListener,omitempty"`
	LoadDistributionPolicy   *CommonSubResource                        `json:"loadDistributionPolicy,omitempty"`
	Priority                 *int64                                    `json:"priority,omitempty"`
	ProvisioningState        *ProvisioningState                        `json:"provisioningState,omitempty"`
	RedirectConfiguration    *CommonSubResource                        `json:"redirectConfiguration,omitempty"`
	RewriteRuleSet           *CommonSubResource                        `json:"rewriteRuleSet,omitempty"`
	RuleType                 *ApplicationGatewayRequestRoutingRuleType `json:"ruleType,omitempty"`
	UrlPathMap               *CommonSubResource                        `json:"urlPathMap,omitempty"`
}
