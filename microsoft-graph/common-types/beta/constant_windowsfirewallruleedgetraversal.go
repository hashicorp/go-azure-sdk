package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsFirewallRuleEdgeTraversal string

const (
	WindowsFirewallRuleEdgeTraversal_Allowed       WindowsFirewallRuleEdgeTraversal = "allowed"
	WindowsFirewallRuleEdgeTraversal_Blocked       WindowsFirewallRuleEdgeTraversal = "blocked"
	WindowsFirewallRuleEdgeTraversal_NotConfigured WindowsFirewallRuleEdgeTraversal = "notConfigured"
)

func PossibleValuesForWindowsFirewallRuleEdgeTraversal() []string {
	return []string{
		string(WindowsFirewallRuleEdgeTraversal_Allowed),
		string(WindowsFirewallRuleEdgeTraversal_Blocked),
		string(WindowsFirewallRuleEdgeTraversal_NotConfigured),
	}
}

func (s *WindowsFirewallRuleEdgeTraversal) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsFirewallRuleEdgeTraversal(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsFirewallRuleEdgeTraversal(input string) (*WindowsFirewallRuleEdgeTraversal, error) {
	vals := map[string]WindowsFirewallRuleEdgeTraversal{
		"allowed":       WindowsFirewallRuleEdgeTraversal_Allowed,
		"blocked":       WindowsFirewallRuleEdgeTraversal_Blocked,
		"notconfigured": WindowsFirewallRuleEdgeTraversal_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsFirewallRuleEdgeTraversal(input)
	return &out, nil
}
