package dnssecurityrules

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DnsSecurityRulePatchProperties struct {
	Action                 *DnsSecurityRuleAction `json:"action,omitempty"`
	DnsResolverDomainLists *[]SubResource         `json:"dnsResolverDomainLists,omitempty"`
	DnsSecurityRuleState   *DnsSecurityRuleState  `json:"dnsSecurityRuleState,omitempty"`
	Priority               *int64                 `json:"priority,omitempty"`
}
