package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessNetworkAccessTrafficAction string

const (
	NetworkaccessNetworkAccessTrafficAction_Allow NetworkaccessNetworkAccessTrafficAction = "allow"
	NetworkaccessNetworkAccessTrafficAction_Block NetworkaccessNetworkAccessTrafficAction = "block"
)

func PossibleValuesForNetworkaccessNetworkAccessTrafficAction() []string {
	return []string{
		string(NetworkaccessNetworkAccessTrafficAction_Allow),
		string(NetworkaccessNetworkAccessTrafficAction_Block),
	}
}

func (s *NetworkaccessNetworkAccessTrafficAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessNetworkAccessTrafficAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessNetworkAccessTrafficAction(input string) (*NetworkaccessNetworkAccessTrafficAction, error) {
	vals := map[string]NetworkaccessNetworkAccessTrafficAction{
		"allow": NetworkaccessNetworkAccessTrafficAction_Allow,
		"block": NetworkaccessNetworkAccessTrafficAction_Block,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessNetworkAccessTrafficAction(input)
	return &out, nil
}
