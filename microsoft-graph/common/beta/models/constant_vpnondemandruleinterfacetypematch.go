package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VpnOnDemandRuleInterfaceTypeMatch string

const (
	VpnOnDemandRuleInterfaceTypeMatchcellular      VpnOnDemandRuleInterfaceTypeMatch = "Cellular"
	VpnOnDemandRuleInterfaceTypeMatchethernet      VpnOnDemandRuleInterfaceTypeMatch = "Ethernet"
	VpnOnDemandRuleInterfaceTypeMatchnotConfigured VpnOnDemandRuleInterfaceTypeMatch = "NotConfigured"
	VpnOnDemandRuleInterfaceTypeMatchwiFi          VpnOnDemandRuleInterfaceTypeMatch = "WiFi"
)

func PossibleValuesForVpnOnDemandRuleInterfaceTypeMatch() []string {
	return []string{
		string(VpnOnDemandRuleInterfaceTypeMatchcellular),
		string(VpnOnDemandRuleInterfaceTypeMatchethernet),
		string(VpnOnDemandRuleInterfaceTypeMatchnotConfigured),
		string(VpnOnDemandRuleInterfaceTypeMatchwiFi),
	}
}

func parseVpnOnDemandRuleInterfaceTypeMatch(input string) (*VpnOnDemandRuleInterfaceTypeMatch, error) {
	vals := map[string]VpnOnDemandRuleInterfaceTypeMatch{
		"cellular":      VpnOnDemandRuleInterfaceTypeMatchcellular,
		"ethernet":      VpnOnDemandRuleInterfaceTypeMatchethernet,
		"notconfigured": VpnOnDemandRuleInterfaceTypeMatchnotConfigured,
		"wifi":          VpnOnDemandRuleInterfaceTypeMatchwiFi,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VpnOnDemandRuleInterfaceTypeMatch(input)
	return &out, nil
}
