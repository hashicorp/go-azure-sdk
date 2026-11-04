package ipv6firewallrules

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IPv6ServerFirewallRuleProperties struct {
	EndIPv6Address   *string `json:"endIPv6Address,omitempty"`
	StartIPv6Address *string `json:"startIPv6Address,omitempty"`
}
