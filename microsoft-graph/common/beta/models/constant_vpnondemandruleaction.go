package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VpnOnDemandRuleAction string

const (
	VpnOnDemandRuleActionconnect            VpnOnDemandRuleAction = "Connect"
	VpnOnDemandRuleActiondisconnect         VpnOnDemandRuleAction = "Disconnect"
	VpnOnDemandRuleActionevaluateConnection VpnOnDemandRuleAction = "EvaluateConnection"
	VpnOnDemandRuleActionignore             VpnOnDemandRuleAction = "Ignore"
)

func PossibleValuesForVpnOnDemandRuleAction() []string {
	return []string{
		string(VpnOnDemandRuleActionconnect),
		string(VpnOnDemandRuleActiondisconnect),
		string(VpnOnDemandRuleActionevaluateConnection),
		string(VpnOnDemandRuleActionignore),
	}
}

func parseVpnOnDemandRuleAction(input string) (*VpnOnDemandRuleAction, error) {
	vals := map[string]VpnOnDemandRuleAction{
		"connect":            VpnOnDemandRuleActionconnect,
		"disconnect":         VpnOnDemandRuleActiondisconnect,
		"evaluateconnection": VpnOnDemandRuleActionevaluateConnection,
		"ignore":             VpnOnDemandRuleActionignore,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VpnOnDemandRuleAction(input)
	return &out, nil
}
