package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessEntitiesSummaryTrafficType string

const (
	NetworkaccessEntitiesSummaryTrafficType_All          NetworkaccessEntitiesSummaryTrafficType = "all"
	NetworkaccessEntitiesSummaryTrafficType_Internet     NetworkaccessEntitiesSummaryTrafficType = "internet"
	NetworkaccessEntitiesSummaryTrafficType_Microsoft365 NetworkaccessEntitiesSummaryTrafficType = "microsoft365"
	NetworkaccessEntitiesSummaryTrafficType_Private      NetworkaccessEntitiesSummaryTrafficType = "private"
)

func PossibleValuesForNetworkaccessEntitiesSummaryTrafficType() []string {
	return []string{
		string(NetworkaccessEntitiesSummaryTrafficType_All),
		string(NetworkaccessEntitiesSummaryTrafficType_Internet),
		string(NetworkaccessEntitiesSummaryTrafficType_Microsoft365),
		string(NetworkaccessEntitiesSummaryTrafficType_Private),
	}
}

func (s *NetworkaccessEntitiesSummaryTrafficType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessEntitiesSummaryTrafficType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessEntitiesSummaryTrafficType(input string) (*NetworkaccessEntitiesSummaryTrafficType, error) {
	vals := map[string]NetworkaccessEntitiesSummaryTrafficType{
		"all":          NetworkaccessEntitiesSummaryTrafficType_All,
		"internet":     NetworkaccessEntitiesSummaryTrafficType_Internet,
		"microsoft365": NetworkaccessEntitiesSummaryTrafficType_Microsoft365,
		"private":      NetworkaccessEntitiesSummaryTrafficType_Private,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessEntitiesSummaryTrafficType(input)
	return &out, nil
}
