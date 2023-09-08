package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VpnTrafficRuleVpnTrafficDirection string

const (
	VpnTrafficRuleVpnTrafficDirectioninbound  VpnTrafficRuleVpnTrafficDirection = "Inbound"
	VpnTrafficRuleVpnTrafficDirectionoutbound VpnTrafficRuleVpnTrafficDirection = "Outbound"
)

func PossibleValuesForVpnTrafficRuleVpnTrafficDirection() []string {
	return []string{
		string(VpnTrafficRuleVpnTrafficDirectioninbound),
		string(VpnTrafficRuleVpnTrafficDirectionoutbound),
	}
}

func parseVpnTrafficRuleVpnTrafficDirection(input string) (*VpnTrafficRuleVpnTrafficDirection, error) {
	vals := map[string]VpnTrafficRuleVpnTrafficDirection{
		"inbound":  VpnTrafficRuleVpnTrafficDirectioninbound,
		"outbound": VpnTrafficRuleVpnTrafficDirectionoutbound,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VpnTrafficRuleVpnTrafficDirection(input)
	return &out, nil
}
