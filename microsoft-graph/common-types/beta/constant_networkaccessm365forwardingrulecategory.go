package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessM365ForwardingRuleCategory string

const (
	NetworkaccessM365ForwardingRuleCategory_Allow     NetworkaccessM365ForwardingRuleCategory = "allow"
	NetworkaccessM365ForwardingRuleCategory_Default   NetworkaccessM365ForwardingRuleCategory = "default"
	NetworkaccessM365ForwardingRuleCategory_Optimized NetworkaccessM365ForwardingRuleCategory = "optimized"
)

func PossibleValuesForNetworkaccessM365ForwardingRuleCategory() []string {
	return []string{
		string(NetworkaccessM365ForwardingRuleCategory_Allow),
		string(NetworkaccessM365ForwardingRuleCategory_Default),
		string(NetworkaccessM365ForwardingRuleCategory_Optimized),
	}
}

func (s *NetworkaccessM365ForwardingRuleCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessM365ForwardingRuleCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessM365ForwardingRuleCategory(input string) (*NetworkaccessM365ForwardingRuleCategory, error) {
	vals := map[string]NetworkaccessM365ForwardingRuleCategory{
		"allow":     NetworkaccessM365ForwardingRuleCategory_Allow,
		"default":   NetworkaccessM365ForwardingRuleCategory_Default,
		"optimized": NetworkaccessM365ForwardingRuleCategory_Optimized,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessM365ForwardingRuleCategory(input)
	return &out, nil
}
