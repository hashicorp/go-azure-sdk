package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessForwardingPolicyTrafficForwardingType string

const (
	NetworkaccessForwardingPolicyTrafficForwardingType_Internet NetworkaccessForwardingPolicyTrafficForwardingType = "internet"
	NetworkaccessForwardingPolicyTrafficForwardingType_M365     NetworkaccessForwardingPolicyTrafficForwardingType = "m365"
	NetworkaccessForwardingPolicyTrafficForwardingType_Private  NetworkaccessForwardingPolicyTrafficForwardingType = "private"
)

func PossibleValuesForNetworkaccessForwardingPolicyTrafficForwardingType() []string {
	return []string{
		string(NetworkaccessForwardingPolicyTrafficForwardingType_Internet),
		string(NetworkaccessForwardingPolicyTrafficForwardingType_M365),
		string(NetworkaccessForwardingPolicyTrafficForwardingType_Private),
	}
}

func (s *NetworkaccessForwardingPolicyTrafficForwardingType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessForwardingPolicyTrafficForwardingType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessForwardingPolicyTrafficForwardingType(input string) (*NetworkaccessForwardingPolicyTrafficForwardingType, error) {
	vals := map[string]NetworkaccessForwardingPolicyTrafficForwardingType{
		"internet": NetworkaccessForwardingPolicyTrafficForwardingType_Internet,
		"m365":     NetworkaccessForwardingPolicyTrafficForwardingType_M365,
		"private":  NetworkaccessForwardingPolicyTrafficForwardingType_Private,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessForwardingPolicyTrafficForwardingType(input)
	return &out, nil
}
