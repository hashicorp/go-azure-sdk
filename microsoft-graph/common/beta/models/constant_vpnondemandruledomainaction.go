package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VpnOnDemandRuleDomainAction string

const (
	VpnOnDemandRuleDomainActionconnectIfNeeded VpnOnDemandRuleDomainAction = "ConnectIfNeeded"
	VpnOnDemandRuleDomainActionneverConnect    VpnOnDemandRuleDomainAction = "NeverConnect"
)

func PossibleValuesForVpnOnDemandRuleDomainAction() []string {
	return []string{
		string(VpnOnDemandRuleDomainActionconnectIfNeeded),
		string(VpnOnDemandRuleDomainActionneverConnect),
	}
}

func parseVpnOnDemandRuleDomainAction(input string) (*VpnOnDemandRuleDomainAction, error) {
	vals := map[string]VpnOnDemandRuleDomainAction{
		"connectifneeded": VpnOnDemandRuleDomainActionconnectIfNeeded,
		"neverconnect":    VpnOnDemandRuleDomainActionneverConnect,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VpnOnDemandRuleDomainAction(input)
	return &out, nil
}
