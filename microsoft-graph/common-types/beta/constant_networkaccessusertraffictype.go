package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessUserTrafficType string

const (
	NetworkaccessUserTrafficType_All          NetworkaccessUserTrafficType = "all"
	NetworkaccessUserTrafficType_Internet     NetworkaccessUserTrafficType = "internet"
	NetworkaccessUserTrafficType_Microsoft365 NetworkaccessUserTrafficType = "microsoft365"
	NetworkaccessUserTrafficType_Private      NetworkaccessUserTrafficType = "private"
)

func PossibleValuesForNetworkaccessUserTrafficType() []string {
	return []string{
		string(NetworkaccessUserTrafficType_All),
		string(NetworkaccessUserTrafficType_Internet),
		string(NetworkaccessUserTrafficType_Microsoft365),
		string(NetworkaccessUserTrafficType_Private),
	}
}

func (s *NetworkaccessUserTrafficType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessUserTrafficType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessUserTrafficType(input string) (*NetworkaccessUserTrafficType, error) {
	vals := map[string]NetworkaccessUserTrafficType{
		"all":          NetworkaccessUserTrafficType_All,
		"internet":     NetworkaccessUserTrafficType_Internet,
		"microsoft365": NetworkaccessUserTrafficType_Microsoft365,
		"private":      NetworkaccessUserTrafficType_Private,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessUserTrafficType(input)
	return &out, nil
}
