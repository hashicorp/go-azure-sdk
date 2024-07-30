package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkConnectionStatus string

const (
	NetworkConnectionStatus_Attempted NetworkConnectionStatus = "attempted"
	NetworkConnectionStatus_Blocked   NetworkConnectionStatus = "blocked"
	NetworkConnectionStatus_Failed    NetworkConnectionStatus = "failed"
	NetworkConnectionStatus_Succeeded NetworkConnectionStatus = "succeeded"
	NetworkConnectionStatus_Unknown   NetworkConnectionStatus = "unknown"
)

func PossibleValuesForNetworkConnectionStatus() []string {
	return []string{
		string(NetworkConnectionStatus_Attempted),
		string(NetworkConnectionStatus_Blocked),
		string(NetworkConnectionStatus_Failed),
		string(NetworkConnectionStatus_Succeeded),
		string(NetworkConnectionStatus_Unknown),
	}
}

func (s *NetworkConnectionStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkConnectionStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkConnectionStatus(input string) (*NetworkConnectionStatus, error) {
	vals := map[string]NetworkConnectionStatus{
		"attempted": NetworkConnectionStatus_Attempted,
		"blocked":   NetworkConnectionStatus_Blocked,
		"failed":    NetworkConnectionStatus_Failed,
		"succeeded": NetworkConnectionStatus_Succeeded,
		"unknown":   NetworkConnectionStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkConnectionStatus(input)
	return &out, nil
}
