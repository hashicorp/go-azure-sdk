package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessForwardingOptionsSkipDnsLookupState string

const (
	NetworkaccessForwardingOptionsSkipDnsLookupState_Disabled NetworkaccessForwardingOptionsSkipDnsLookupState = "disabled"
	NetworkaccessForwardingOptionsSkipDnsLookupState_Enabled  NetworkaccessForwardingOptionsSkipDnsLookupState = "enabled"
)

func PossibleValuesForNetworkaccessForwardingOptionsSkipDnsLookupState() []string {
	return []string{
		string(NetworkaccessForwardingOptionsSkipDnsLookupState_Disabled),
		string(NetworkaccessForwardingOptionsSkipDnsLookupState_Enabled),
	}
}

func (s *NetworkaccessForwardingOptionsSkipDnsLookupState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessForwardingOptionsSkipDnsLookupState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessForwardingOptionsSkipDnsLookupState(input string) (*NetworkaccessForwardingOptionsSkipDnsLookupState, error) {
	vals := map[string]NetworkaccessForwardingOptionsSkipDnsLookupState{
		"disabled": NetworkaccessForwardingOptionsSkipDnsLookupState_Disabled,
		"enabled":  NetworkaccessForwardingOptionsSkipDnsLookupState_Enabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessForwardingOptionsSkipDnsLookupState(input)
	return &out, nil
}
