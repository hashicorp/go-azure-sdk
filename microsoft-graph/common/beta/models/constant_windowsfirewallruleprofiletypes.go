package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsFirewallRuleProfileTypes string

const (
	WindowsFirewallRuleProfileTypesdomain        WindowsFirewallRuleProfileTypes = "Domain"
	WindowsFirewallRuleProfileTypesnotConfigured WindowsFirewallRuleProfileTypes = "NotConfigured"
	WindowsFirewallRuleProfileTypesprivate       WindowsFirewallRuleProfileTypes = "Private"
	WindowsFirewallRuleProfileTypespublic        WindowsFirewallRuleProfileTypes = "Public"
)

func PossibleValuesForWindowsFirewallRuleProfileTypes() []string {
	return []string{
		string(WindowsFirewallRuleProfileTypesdomain),
		string(WindowsFirewallRuleProfileTypesnotConfigured),
		string(WindowsFirewallRuleProfileTypesprivate),
		string(WindowsFirewallRuleProfileTypespublic),
	}
}

func parseWindowsFirewallRuleProfileTypes(input string) (*WindowsFirewallRuleProfileTypes, error) {
	vals := map[string]WindowsFirewallRuleProfileTypes{
		"domain":        WindowsFirewallRuleProfileTypesdomain,
		"notconfigured": WindowsFirewallRuleProfileTypesnotConfigured,
		"private":       WindowsFirewallRuleProfileTypesprivate,
		"public":        WindowsFirewallRuleProfileTypespublic,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsFirewallRuleProfileTypes(input)
	return &out, nil
}
