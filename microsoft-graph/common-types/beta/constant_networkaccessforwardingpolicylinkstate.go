package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessForwardingPolicyLinkState string

const (
	NetworkaccessForwardingPolicyLinkState_Disabled NetworkaccessForwardingPolicyLinkState = "disabled"
	NetworkaccessForwardingPolicyLinkState_Enabled  NetworkaccessForwardingPolicyLinkState = "enabled"
)

func PossibleValuesForNetworkaccessForwardingPolicyLinkState() []string {
	return []string{
		string(NetworkaccessForwardingPolicyLinkState_Disabled),
		string(NetworkaccessForwardingPolicyLinkState_Enabled),
	}
}

func (s *NetworkaccessForwardingPolicyLinkState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessForwardingPolicyLinkState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessForwardingPolicyLinkState(input string) (*NetworkaccessForwardingPolicyLinkState, error) {
	vals := map[string]NetworkaccessForwardingPolicyLinkState{
		"disabled": NetworkaccessForwardingPolicyLinkState_Disabled,
		"enabled":  NetworkaccessForwardingPolicyLinkState_Enabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessForwardingPolicyLinkState(input)
	return &out, nil
}
