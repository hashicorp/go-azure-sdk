package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkConnectionDirection string

const (
	NetworkConnectionDirection_Inbound  NetworkConnectionDirection = "inbound"
	NetworkConnectionDirection_Outbound NetworkConnectionDirection = "outbound"
	NetworkConnectionDirection_Unknown  NetworkConnectionDirection = "unknown"
)

func PossibleValuesForNetworkConnectionDirection() []string {
	return []string{
		string(NetworkConnectionDirection_Inbound),
		string(NetworkConnectionDirection_Outbound),
		string(NetworkConnectionDirection_Unknown),
	}
}

func (s *NetworkConnectionDirection) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkConnectionDirection(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkConnectionDirection(input string) (*NetworkConnectionDirection, error) {
	vals := map[string]NetworkConnectionDirection{
		"inbound":  NetworkConnectionDirection_Inbound,
		"outbound": NetworkConnectionDirection_Outbound,
		"unknown":  NetworkConnectionDirection_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkConnectionDirection(input)
	return &out, nil
}
