package dnssecurityrules

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DnsSecurityRuleProperties struct {
	Action                 DnsSecurityRuleAction `json:"action"`
	DnsResolverDomainLists []SubResource         `json:"dnsResolverDomainLists"`
	DnsSecurityRuleState   *DnsSecurityRuleState `json:"dnsSecurityRuleState,omitempty"`
	Priority               int64                 `json:"priority"`
	ProvisioningState      *ProvisioningState    `json:"provisioningState,omitempty"`
}
