package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsFirewallRuleAction string

const (
	WindowsFirewallRuleAction_Allowed       WindowsFirewallRuleAction = "allowed"
	WindowsFirewallRuleAction_Blocked       WindowsFirewallRuleAction = "blocked"
	WindowsFirewallRuleAction_NotConfigured WindowsFirewallRuleAction = "notConfigured"
)

func PossibleValuesForWindowsFirewallRuleAction() []string {
	return []string{
		string(WindowsFirewallRuleAction_Allowed),
		string(WindowsFirewallRuleAction_Blocked),
		string(WindowsFirewallRuleAction_NotConfigured),
	}
}

func (s *WindowsFirewallRuleAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsFirewallRuleAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsFirewallRuleAction(input string) (*WindowsFirewallRuleAction, error) {
	vals := map[string]WindowsFirewallRuleAction{
		"allowed":       WindowsFirewallRuleAction_Allowed,
		"blocked":       WindowsFirewallRuleAction_Blocked,
		"notconfigured": WindowsFirewallRuleAction_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsFirewallRuleAction(input)
	return &out, nil
}
