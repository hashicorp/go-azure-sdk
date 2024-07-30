package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsFirewallRuleProfileTypes string

const (
	WindowsFirewallRuleProfileTypes_Domain        WindowsFirewallRuleProfileTypes = "domain"
	WindowsFirewallRuleProfileTypes_NotConfigured WindowsFirewallRuleProfileTypes = "notConfigured"
	WindowsFirewallRuleProfileTypes_Private       WindowsFirewallRuleProfileTypes = "private"
	WindowsFirewallRuleProfileTypes_Public        WindowsFirewallRuleProfileTypes = "public"
)

func PossibleValuesForWindowsFirewallRuleProfileTypes() []string {
	return []string{
		string(WindowsFirewallRuleProfileTypes_Domain),
		string(WindowsFirewallRuleProfileTypes_NotConfigured),
		string(WindowsFirewallRuleProfileTypes_Private),
		string(WindowsFirewallRuleProfileTypes_Public),
	}
}

func (s *WindowsFirewallRuleProfileTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsFirewallRuleProfileTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsFirewallRuleProfileTypes(input string) (*WindowsFirewallRuleProfileTypes, error) {
	vals := map[string]WindowsFirewallRuleProfileTypes{
		"domain":        WindowsFirewallRuleProfileTypes_Domain,
		"notconfigured": WindowsFirewallRuleProfileTypes_NotConfigured,
		"private":       WindowsFirewallRuleProfileTypes_Private,
		"public":        WindowsFirewallRuleProfileTypes_Public,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsFirewallRuleProfileTypes(input)
	return &out, nil
}
