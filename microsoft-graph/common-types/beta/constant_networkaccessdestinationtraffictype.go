package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessDestinationTrafficType string

const (
	NetworkaccessDestinationTrafficType_All          NetworkaccessDestinationTrafficType = "all"
	NetworkaccessDestinationTrafficType_Internet     NetworkaccessDestinationTrafficType = "internet"
	NetworkaccessDestinationTrafficType_Microsoft365 NetworkaccessDestinationTrafficType = "microsoft365"
	NetworkaccessDestinationTrafficType_Private      NetworkaccessDestinationTrafficType = "private"
)

func PossibleValuesForNetworkaccessDestinationTrafficType() []string {
	return []string{
		string(NetworkaccessDestinationTrafficType_All),
		string(NetworkaccessDestinationTrafficType_Internet),
		string(NetworkaccessDestinationTrafficType_Microsoft365),
		string(NetworkaccessDestinationTrafficType_Private),
	}
}

func (s *NetworkaccessDestinationTrafficType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessDestinationTrafficType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessDestinationTrafficType(input string) (*NetworkaccessDestinationTrafficType, error) {
	vals := map[string]NetworkaccessDestinationTrafficType{
		"all":          NetworkaccessDestinationTrafficType_All,
		"internet":     NetworkaccessDestinationTrafficType_Internet,
		"microsoft365": NetworkaccessDestinationTrafficType_Microsoft365,
		"private":      NetworkaccessDestinationTrafficType_Private,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessDestinationTrafficType(input)
	return &out, nil
}
