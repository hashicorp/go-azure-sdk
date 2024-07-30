package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessFilteringPolicyLinkLoggingState string

const (
	NetworkaccessFilteringPolicyLinkLoggingState_Disabled NetworkaccessFilteringPolicyLinkLoggingState = "disabled"
	NetworkaccessFilteringPolicyLinkLoggingState_Enabled  NetworkaccessFilteringPolicyLinkLoggingState = "enabled"
)

func PossibleValuesForNetworkaccessFilteringPolicyLinkLoggingState() []string {
	return []string{
		string(NetworkaccessFilteringPolicyLinkLoggingState_Disabled),
		string(NetworkaccessFilteringPolicyLinkLoggingState_Enabled),
	}
}

func (s *NetworkaccessFilteringPolicyLinkLoggingState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessFilteringPolicyLinkLoggingState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessFilteringPolicyLinkLoggingState(input string) (*NetworkaccessFilteringPolicyLinkLoggingState, error) {
	vals := map[string]NetworkaccessFilteringPolicyLinkLoggingState{
		"disabled": NetworkaccessFilteringPolicyLinkLoggingState_Disabled,
		"enabled":  NetworkaccessFilteringPolicyLinkLoggingState_Enabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessFilteringPolicyLinkLoggingState(input)
	return &out, nil
}
