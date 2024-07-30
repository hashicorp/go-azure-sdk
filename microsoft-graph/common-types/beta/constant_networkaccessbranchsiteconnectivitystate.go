package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessBranchSiteConnectivityState string

const (
	NetworkaccessBranchSiteConnectivityState_Connected NetworkaccessBranchSiteConnectivityState = "connected"
	NetworkaccessBranchSiteConnectivityState_Error     NetworkaccessBranchSiteConnectivityState = "error"
	NetworkaccessBranchSiteConnectivityState_Inactive  NetworkaccessBranchSiteConnectivityState = "inactive"
	NetworkaccessBranchSiteConnectivityState_Pending   NetworkaccessBranchSiteConnectivityState = "pending"
)

func PossibleValuesForNetworkaccessBranchSiteConnectivityState() []string {
	return []string{
		string(NetworkaccessBranchSiteConnectivityState_Connected),
		string(NetworkaccessBranchSiteConnectivityState_Error),
		string(NetworkaccessBranchSiteConnectivityState_Inactive),
		string(NetworkaccessBranchSiteConnectivityState_Pending),
	}
}

func (s *NetworkaccessBranchSiteConnectivityState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessBranchSiteConnectivityState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessBranchSiteConnectivityState(input string) (*NetworkaccessBranchSiteConnectivityState, error) {
	vals := map[string]NetworkaccessBranchSiteConnectivityState{
		"connected": NetworkaccessBranchSiteConnectivityState_Connected,
		"error":     NetworkaccessBranchSiteConnectivityState_Error,
		"inactive":  NetworkaccessBranchSiteConnectivityState_Inactive,
		"pending":   NetworkaccessBranchSiteConnectivityState_Pending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessBranchSiteConnectivityState(input)
	return &out, nil
}
