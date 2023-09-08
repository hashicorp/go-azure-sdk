package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsFirewallNetworkProfileFirewallEnabled string

const (
	WindowsFirewallNetworkProfileFirewallEnabledallowed       WindowsFirewallNetworkProfileFirewallEnabled = "Allowed"
	WindowsFirewallNetworkProfileFirewallEnabledblocked       WindowsFirewallNetworkProfileFirewallEnabled = "Blocked"
	WindowsFirewallNetworkProfileFirewallEnablednotConfigured WindowsFirewallNetworkProfileFirewallEnabled = "NotConfigured"
)

func PossibleValuesForWindowsFirewallNetworkProfileFirewallEnabled() []string {
	return []string{
		string(WindowsFirewallNetworkProfileFirewallEnabledallowed),
		string(WindowsFirewallNetworkProfileFirewallEnabledblocked),
		string(WindowsFirewallNetworkProfileFirewallEnablednotConfigured),
	}
}

func parseWindowsFirewallNetworkProfileFirewallEnabled(input string) (*WindowsFirewallNetworkProfileFirewallEnabled, error) {
	vals := map[string]WindowsFirewallNetworkProfileFirewallEnabled{
		"allowed":       WindowsFirewallNetworkProfileFirewallEnabledallowed,
		"blocked":       WindowsFirewallNetworkProfileFirewallEnabledblocked,
		"notconfigured": WindowsFirewallNetworkProfileFirewallEnablednotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsFirewallNetworkProfileFirewallEnabled(input)
	return &out, nil
}
