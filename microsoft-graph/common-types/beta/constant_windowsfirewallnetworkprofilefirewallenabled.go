package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsFirewallNetworkProfileFirewallEnabled string

const (
	WindowsFirewallNetworkProfileFirewallEnabled_Allowed       WindowsFirewallNetworkProfileFirewallEnabled = "allowed"
	WindowsFirewallNetworkProfileFirewallEnabled_Blocked       WindowsFirewallNetworkProfileFirewallEnabled = "blocked"
	WindowsFirewallNetworkProfileFirewallEnabled_NotConfigured WindowsFirewallNetworkProfileFirewallEnabled = "notConfigured"
)

func PossibleValuesForWindowsFirewallNetworkProfileFirewallEnabled() []string {
	return []string{
		string(WindowsFirewallNetworkProfileFirewallEnabled_Allowed),
		string(WindowsFirewallNetworkProfileFirewallEnabled_Blocked),
		string(WindowsFirewallNetworkProfileFirewallEnabled_NotConfigured),
	}
}

func (s *WindowsFirewallNetworkProfileFirewallEnabled) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsFirewallNetworkProfileFirewallEnabled(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsFirewallNetworkProfileFirewallEnabled(input string) (*WindowsFirewallNetworkProfileFirewallEnabled, error) {
	vals := map[string]WindowsFirewallNetworkProfileFirewallEnabled{
		"allowed":       WindowsFirewallNetworkProfileFirewallEnabled_Allowed,
		"blocked":       WindowsFirewallNetworkProfileFirewallEnabled_Blocked,
		"notconfigured": WindowsFirewallNetworkProfileFirewallEnabled_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsFirewallNetworkProfileFirewallEnabled(input)
	return &out, nil
}
