package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VpnTrafficRuleVpnTrafficDirection string

const (
	VpnTrafficRuleVpnTrafficDirection_Inbound  VpnTrafficRuleVpnTrafficDirection = "inbound"
	VpnTrafficRuleVpnTrafficDirection_Outbound VpnTrafficRuleVpnTrafficDirection = "outbound"
)

func PossibleValuesForVpnTrafficRuleVpnTrafficDirection() []string {
	return []string{
		string(VpnTrafficRuleVpnTrafficDirection_Inbound),
		string(VpnTrafficRuleVpnTrafficDirection_Outbound),
	}
}

func (s *VpnTrafficRuleVpnTrafficDirection) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVpnTrafficRuleVpnTrafficDirection(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVpnTrafficRuleVpnTrafficDirection(input string) (*VpnTrafficRuleVpnTrafficDirection, error) {
	vals := map[string]VpnTrafficRuleVpnTrafficDirection{
		"inbound":  VpnTrafficRuleVpnTrafficDirection_Inbound,
		"outbound": VpnTrafficRuleVpnTrafficDirection_Outbound,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VpnTrafficRuleVpnTrafficDirection(input)
	return &out, nil
}
