package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessNetworkAccessTrafficTrafficType string

const (
	NetworkaccessNetworkAccessTrafficTrafficType_All          NetworkaccessNetworkAccessTrafficTrafficType = "all"
	NetworkaccessNetworkAccessTrafficTrafficType_Internet     NetworkaccessNetworkAccessTrafficTrafficType = "internet"
	NetworkaccessNetworkAccessTrafficTrafficType_Microsoft365 NetworkaccessNetworkAccessTrafficTrafficType = "microsoft365"
	NetworkaccessNetworkAccessTrafficTrafficType_Private      NetworkaccessNetworkAccessTrafficTrafficType = "private"
)

func PossibleValuesForNetworkaccessNetworkAccessTrafficTrafficType() []string {
	return []string{
		string(NetworkaccessNetworkAccessTrafficTrafficType_All),
		string(NetworkaccessNetworkAccessTrafficTrafficType_Internet),
		string(NetworkaccessNetworkAccessTrafficTrafficType_Microsoft365),
		string(NetworkaccessNetworkAccessTrafficTrafficType_Private),
	}
}

func (s *NetworkaccessNetworkAccessTrafficTrafficType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessNetworkAccessTrafficTrafficType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessNetworkAccessTrafficTrafficType(input string) (*NetworkaccessNetworkAccessTrafficTrafficType, error) {
	vals := map[string]NetworkaccessNetworkAccessTrafficTrafficType{
		"all":          NetworkaccessNetworkAccessTrafficTrafficType_All,
		"internet":     NetworkaccessNetworkAccessTrafficTrafficType_Internet,
		"microsoft365": NetworkaccessNetworkAccessTrafficTrafficType_Microsoft365,
		"private":      NetworkaccessNetworkAccessTrafficTrafficType_Private,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessNetworkAccessTrafficTrafficType(input)
	return &out, nil
}
