package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessPolicyRuleDeltaAction string

const (
	NetworkaccessPolicyRuleDeltaAction_Bypass  NetworkaccessPolicyRuleDeltaAction = "bypass"
	NetworkaccessPolicyRuleDeltaAction_Forward NetworkaccessPolicyRuleDeltaAction = "forward"
)

func PossibleValuesForNetworkaccessPolicyRuleDeltaAction() []string {
	return []string{
		string(NetworkaccessPolicyRuleDeltaAction_Bypass),
		string(NetworkaccessPolicyRuleDeltaAction_Forward),
	}
}

func (s *NetworkaccessPolicyRuleDeltaAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessPolicyRuleDeltaAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessPolicyRuleDeltaAction(input string) (*NetworkaccessPolicyRuleDeltaAction, error) {
	vals := map[string]NetworkaccessPolicyRuleDeltaAction{
		"bypass":  NetworkaccessPolicyRuleDeltaAction_Bypass,
		"forward": NetworkaccessPolicyRuleDeltaAction_Forward,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessPolicyRuleDeltaAction(input)
	return &out, nil
}
