package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsFirewallRuleInterfaceTypes string

const (
	WindowsFirewallRuleInterfaceTypeslan           WindowsFirewallRuleInterfaceTypes = "Lan"
	WindowsFirewallRuleInterfaceTypesnotConfigured WindowsFirewallRuleInterfaceTypes = "NotConfigured"
	WindowsFirewallRuleInterfaceTypesremoteAccess  WindowsFirewallRuleInterfaceTypes = "RemoteAccess"
	WindowsFirewallRuleInterfaceTypeswireless      WindowsFirewallRuleInterfaceTypes = "Wireless"
)

func PossibleValuesForWindowsFirewallRuleInterfaceTypes() []string {
	return []string{
		string(WindowsFirewallRuleInterfaceTypeslan),
		string(WindowsFirewallRuleInterfaceTypesnotConfigured),
		string(WindowsFirewallRuleInterfaceTypesremoteAccess),
		string(WindowsFirewallRuleInterfaceTypeswireless),
	}
}

func parseWindowsFirewallRuleInterfaceTypes(input string) (*WindowsFirewallRuleInterfaceTypes, error) {
	vals := map[string]WindowsFirewallRuleInterfaceTypes{
		"lan":           WindowsFirewallRuleInterfaceTypeslan,
		"notconfigured": WindowsFirewallRuleInterfaceTypesnotConfigured,
		"remoteaccess":  WindowsFirewallRuleInterfaceTypesremoteAccess,
		"wireless":      WindowsFirewallRuleInterfaceTypeswireless,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsFirewallRuleInterfaceTypes(input)
	return &out, nil
}
