package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsFirewallRuleEdgeTraversal string

const (
	WindowsFirewallRuleEdgeTraversalallowed       WindowsFirewallRuleEdgeTraversal = "Allowed"
	WindowsFirewallRuleEdgeTraversalblocked       WindowsFirewallRuleEdgeTraversal = "Blocked"
	WindowsFirewallRuleEdgeTraversalnotConfigured WindowsFirewallRuleEdgeTraversal = "NotConfigured"
)

func PossibleValuesForWindowsFirewallRuleEdgeTraversal() []string {
	return []string{
		string(WindowsFirewallRuleEdgeTraversalallowed),
		string(WindowsFirewallRuleEdgeTraversalblocked),
		string(WindowsFirewallRuleEdgeTraversalnotConfigured),
	}
}

func parseWindowsFirewallRuleEdgeTraversal(input string) (*WindowsFirewallRuleEdgeTraversal, error) {
	vals := map[string]WindowsFirewallRuleEdgeTraversal{
		"allowed":       WindowsFirewallRuleEdgeTraversalallowed,
		"blocked":       WindowsFirewallRuleEdgeTraversalblocked,
		"notconfigured": WindowsFirewallRuleEdgeTraversalnotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsFirewallRuleEdgeTraversal(input)
	return &out, nil
}
