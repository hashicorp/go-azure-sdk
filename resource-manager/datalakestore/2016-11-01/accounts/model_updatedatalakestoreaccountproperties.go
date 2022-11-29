package accounts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateDataLakeStoreAccountProperties struct {
	DefaultGroup           *string                                          `json:"defaultGroup,omitempty"`
	EncryptionConfig       *UpdateEncryptionConfig                          `json:"encryptionConfig,omitempty"`
	FirewallAllowAzureIPs  *FirewallAllowAzureIPsState                      `json:"firewallAllowAzureIps,omitempty"`
	FirewallRules          *[]UpdateFirewallRuleWithAccountParameters       `json:"firewallRules,omitempty"`
	FirewallState          *FirewallState                                   `json:"firewallState,omitempty"`
	NewTier                *TierType                                        `json:"newTier,omitempty"`
	TrustedIdProviderState *TrustedIdProviderState                          `json:"trustedIdProviderState,omitempty"`
	TrustedIdProviders     *[]UpdateTrustedIdProviderWithAccountParameters  `json:"trustedIdProviders,omitempty"`
	VirtualNetworkRules    *[]UpdateVirtualNetworkRuleWithAccountParameters `json:"virtualNetworkRules,omitempty"`
}
