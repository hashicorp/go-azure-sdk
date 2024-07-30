package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessM365ForwardingRuleAction string

const (
	NetworkaccessM365ForwardingRuleAction_Bypass  NetworkaccessM365ForwardingRuleAction = "bypass"
	NetworkaccessM365ForwardingRuleAction_Forward NetworkaccessM365ForwardingRuleAction = "forward"
)

func PossibleValuesForNetworkaccessM365ForwardingRuleAction() []string {
	return []string{
		string(NetworkaccessM365ForwardingRuleAction_Bypass),
		string(NetworkaccessM365ForwardingRuleAction_Forward),
	}
}

func (s *NetworkaccessM365ForwardingRuleAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessM365ForwardingRuleAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessM365ForwardingRuleAction(input string) (*NetworkaccessM365ForwardingRuleAction, error) {
	vals := map[string]NetworkaccessM365ForwardingRuleAction{
		"bypass":  NetworkaccessM365ForwardingRuleAction_Bypass,
		"forward": NetworkaccessM365ForwardingRuleAction_Forward,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessM365ForwardingRuleAction(input)
	return &out, nil
}
