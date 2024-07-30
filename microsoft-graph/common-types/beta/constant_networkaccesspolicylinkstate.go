package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessPolicyLinkState string

const (
	NetworkaccessPolicyLinkState_Disabled NetworkaccessPolicyLinkState = "disabled"
	NetworkaccessPolicyLinkState_Enabled  NetworkaccessPolicyLinkState = "enabled"
)

func PossibleValuesForNetworkaccessPolicyLinkState() []string {
	return []string{
		string(NetworkaccessPolicyLinkState_Disabled),
		string(NetworkaccessPolicyLinkState_Enabled),
	}
}

func (s *NetworkaccessPolicyLinkState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessPolicyLinkState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessPolicyLinkState(input string) (*NetworkaccessPolicyLinkState, error) {
	vals := map[string]NetworkaccessPolicyLinkState{
		"disabled": NetworkaccessPolicyLinkState_Disabled,
		"enabled":  NetworkaccessPolicyLinkState_Enabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessPolicyLinkState(input)
	return &out, nil
}
