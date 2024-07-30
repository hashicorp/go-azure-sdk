package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessDestinationSummaryTrafficType string

const (
	NetworkaccessDestinationSummaryTrafficType_All          NetworkaccessDestinationSummaryTrafficType = "all"
	NetworkaccessDestinationSummaryTrafficType_Internet     NetworkaccessDestinationSummaryTrafficType = "internet"
	NetworkaccessDestinationSummaryTrafficType_Microsoft365 NetworkaccessDestinationSummaryTrafficType = "microsoft365"
	NetworkaccessDestinationSummaryTrafficType_Private      NetworkaccessDestinationSummaryTrafficType = "private"
)

func PossibleValuesForNetworkaccessDestinationSummaryTrafficType() []string {
	return []string{
		string(NetworkaccessDestinationSummaryTrafficType_All),
		string(NetworkaccessDestinationSummaryTrafficType_Internet),
		string(NetworkaccessDestinationSummaryTrafficType_Microsoft365),
		string(NetworkaccessDestinationSummaryTrafficType_Private),
	}
}

func (s *NetworkaccessDestinationSummaryTrafficType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessDestinationSummaryTrafficType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessDestinationSummaryTrafficType(input string) (*NetworkaccessDestinationSummaryTrafficType, error) {
	vals := map[string]NetworkaccessDestinationSummaryTrafficType{
		"all":          NetworkaccessDestinationSummaryTrafficType_All,
		"internet":     NetworkaccessDestinationSummaryTrafficType_Internet,
		"microsoft365": NetworkaccessDestinationSummaryTrafficType_Microsoft365,
		"private":      NetworkaccessDestinationSummaryTrafficType_Private,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessDestinationSummaryTrafficType(input)
	return &out, nil
}
