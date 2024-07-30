package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessInternetAccessForwardingRuleAction string

const (
	NetworkaccessInternetAccessForwardingRuleAction_Bypass  NetworkaccessInternetAccessForwardingRuleAction = "bypass"
	NetworkaccessInternetAccessForwardingRuleAction_Forward NetworkaccessInternetAccessForwardingRuleAction = "forward"
)

func PossibleValuesForNetworkaccessInternetAccessForwardingRuleAction() []string {
	return []string{
		string(NetworkaccessInternetAccessForwardingRuleAction_Bypass),
		string(NetworkaccessInternetAccessForwardingRuleAction_Forward),
	}
}

func (s *NetworkaccessInternetAccessForwardingRuleAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessInternetAccessForwardingRuleAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessInternetAccessForwardingRuleAction(input string) (*NetworkaccessInternetAccessForwardingRuleAction, error) {
	vals := map[string]NetworkaccessInternetAccessForwardingRuleAction{
		"bypass":  NetworkaccessInternetAccessForwardingRuleAction_Bypass,
		"forward": NetworkaccessInternetAccessForwardingRuleAction_Forward,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessInternetAccessForwardingRuleAction(input)
	return &out, nil
}
