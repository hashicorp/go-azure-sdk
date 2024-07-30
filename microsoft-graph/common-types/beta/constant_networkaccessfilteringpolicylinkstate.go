package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessFilteringPolicyLinkState string

const (
	NetworkaccessFilteringPolicyLinkState_Disabled NetworkaccessFilteringPolicyLinkState = "disabled"
	NetworkaccessFilteringPolicyLinkState_Enabled  NetworkaccessFilteringPolicyLinkState = "enabled"
)

func PossibleValuesForNetworkaccessFilteringPolicyLinkState() []string {
	return []string{
		string(NetworkaccessFilteringPolicyLinkState_Disabled),
		string(NetworkaccessFilteringPolicyLinkState_Enabled),
	}
}

func (s *NetworkaccessFilteringPolicyLinkState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessFilteringPolicyLinkState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessFilteringPolicyLinkState(input string) (*NetworkaccessFilteringPolicyLinkState, error) {
	vals := map[string]NetworkaccessFilteringPolicyLinkState{
		"disabled": NetworkaccessFilteringPolicyLinkState_Disabled,
		"enabled":  NetworkaccessFilteringPolicyLinkState_Enabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessFilteringPolicyLinkState(input)
	return &out, nil
}
