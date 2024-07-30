package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsFirewallRuleTrafficDirection string

const (
	WindowsFirewallRuleTrafficDirection_In            WindowsFirewallRuleTrafficDirection = "in"
	WindowsFirewallRuleTrafficDirection_NotConfigured WindowsFirewallRuleTrafficDirection = "notConfigured"
	WindowsFirewallRuleTrafficDirection_Out           WindowsFirewallRuleTrafficDirection = "out"
)

func PossibleValuesForWindowsFirewallRuleTrafficDirection() []string {
	return []string{
		string(WindowsFirewallRuleTrafficDirection_In),
		string(WindowsFirewallRuleTrafficDirection_NotConfigured),
		string(WindowsFirewallRuleTrafficDirection_Out),
	}
}

func (s *WindowsFirewallRuleTrafficDirection) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsFirewallRuleTrafficDirection(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsFirewallRuleTrafficDirection(input string) (*WindowsFirewallRuleTrafficDirection, error) {
	vals := map[string]WindowsFirewallRuleTrafficDirection{
		"in":            WindowsFirewallRuleTrafficDirection_In,
		"notconfigured": WindowsFirewallRuleTrafficDirection_NotConfigured,
		"out":           WindowsFirewallRuleTrafficDirection_Out,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsFirewallRuleTrafficDirection(input)
	return &out, nil
}
