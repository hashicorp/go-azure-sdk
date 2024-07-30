package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VpnOnDemandRule struct {
	Action                *VpnOnDemandRuleAction             `json:"action,omitempty"`
	DnsSearchDomains      *[]string                          `json:"dnsSearchDomains,omitempty"`
	DnsServerAddressMatch *[]string                          `json:"dnsServerAddressMatch,omitempty"`
	DomainAction          *VpnOnDemandRuleDomainAction       `json:"domainAction,omitempty"`
	Domains               *[]string                          `json:"domains,omitempty"`
	InterfaceTypeMatch    *VpnOnDemandRuleInterfaceTypeMatch `json:"interfaceTypeMatch,omitempty"`
	ODataType             *string                            `json:"@odata.type,omitempty"`
	ProbeRequiredUrl      *string                            `json:"probeRequiredUrl,omitempty"`
	ProbeUrl              *string                            `json:"probeUrl,omitempty"`
	Ssids                 *[]string                          `json:"ssids,omitempty"`
}
