package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VpnOnDemandRuleAction string

const (
	VpnOnDemandRuleAction_Connect            VpnOnDemandRuleAction = "connect"
	VpnOnDemandRuleAction_Disconnect         VpnOnDemandRuleAction = "disconnect"
	VpnOnDemandRuleAction_EvaluateConnection VpnOnDemandRuleAction = "evaluateConnection"
	VpnOnDemandRuleAction_Ignore             VpnOnDemandRuleAction = "ignore"
)

func PossibleValuesForVpnOnDemandRuleAction() []string {
	return []string{
		string(VpnOnDemandRuleAction_Connect),
		string(VpnOnDemandRuleAction_Disconnect),
		string(VpnOnDemandRuleAction_EvaluateConnection),
		string(VpnOnDemandRuleAction_Ignore),
	}
}

func (s *VpnOnDemandRuleAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVpnOnDemandRuleAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVpnOnDemandRuleAction(input string) (*VpnOnDemandRuleAction, error) {
	vals := map[string]VpnOnDemandRuleAction{
		"connect":            VpnOnDemandRuleAction_Connect,
		"disconnect":         VpnOnDemandRuleAction_Disconnect,
		"evaluateconnection": VpnOnDemandRuleAction_EvaluateConnection,
		"ignore":             VpnOnDemandRuleAction_Ignore,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VpnOnDemandRuleAction(input)
	return &out, nil
}
