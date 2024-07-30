package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessForwardingProfileTrafficForwardingType string

const (
	NetworkaccessForwardingProfileTrafficForwardingType_Internet NetworkaccessForwardingProfileTrafficForwardingType = "internet"
	NetworkaccessForwardingProfileTrafficForwardingType_M365     NetworkaccessForwardingProfileTrafficForwardingType = "m365"
	NetworkaccessForwardingProfileTrafficForwardingType_Private  NetworkaccessForwardingProfileTrafficForwardingType = "private"
)

func PossibleValuesForNetworkaccessForwardingProfileTrafficForwardingType() []string {
	return []string{
		string(NetworkaccessForwardingProfileTrafficForwardingType_Internet),
		string(NetworkaccessForwardingProfileTrafficForwardingType_M365),
		string(NetworkaccessForwardingProfileTrafficForwardingType_Private),
	}
}

func (s *NetworkaccessForwardingProfileTrafficForwardingType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessForwardingProfileTrafficForwardingType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessForwardingProfileTrafficForwardingType(input string) (*NetworkaccessForwardingProfileTrafficForwardingType, error) {
	vals := map[string]NetworkaccessForwardingProfileTrafficForwardingType{
		"internet": NetworkaccessForwardingProfileTrafficForwardingType_Internet,
		"m365":     NetworkaccessForwardingProfileTrafficForwardingType_M365,
		"private":  NetworkaccessForwardingProfileTrafficForwardingType_Private,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessForwardingProfileTrafficForwardingType(input)
	return &out, nil
}
