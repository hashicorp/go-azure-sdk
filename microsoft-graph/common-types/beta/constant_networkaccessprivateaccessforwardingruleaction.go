package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessPrivateAccessForwardingRuleAction string

const (
	NetworkaccessPrivateAccessForwardingRuleAction_Bypass  NetworkaccessPrivateAccessForwardingRuleAction = "bypass"
	NetworkaccessPrivateAccessForwardingRuleAction_Forward NetworkaccessPrivateAccessForwardingRuleAction = "forward"
)

func PossibleValuesForNetworkaccessPrivateAccessForwardingRuleAction() []string {
	return []string{
		string(NetworkaccessPrivateAccessForwardingRuleAction_Bypass),
		string(NetworkaccessPrivateAccessForwardingRuleAction_Forward),
	}
}

func (s *NetworkaccessPrivateAccessForwardingRuleAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessPrivateAccessForwardingRuleAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessPrivateAccessForwardingRuleAction(input string) (*NetworkaccessPrivateAccessForwardingRuleAction, error) {
	vals := map[string]NetworkaccessPrivateAccessForwardingRuleAction{
		"bypass":  NetworkaccessPrivateAccessForwardingRuleAction_Bypass,
		"forward": NetworkaccessPrivateAccessForwardingRuleAction_Forward,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessPrivateAccessForwardingRuleAction(input)
	return &out, nil
}
