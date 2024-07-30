package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessForwardingProfileState string

const (
	NetworkaccessForwardingProfileState_Disabled NetworkaccessForwardingProfileState = "disabled"
	NetworkaccessForwardingProfileState_Enabled  NetworkaccessForwardingProfileState = "enabled"
)

func PossibleValuesForNetworkaccessForwardingProfileState() []string {
	return []string{
		string(NetworkaccessForwardingProfileState_Disabled),
		string(NetworkaccessForwardingProfileState_Enabled),
	}
}

func (s *NetworkaccessForwardingProfileState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessForwardingProfileState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessForwardingProfileState(input string) (*NetworkaccessForwardingProfileState, error) {
	vals := map[string]NetworkaccessForwardingProfileState{
		"disabled": NetworkaccessForwardingProfileState_Disabled,
		"enabled":  NetworkaccessForwardingProfileState_Enabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessForwardingProfileState(input)
	return &out, nil
}
