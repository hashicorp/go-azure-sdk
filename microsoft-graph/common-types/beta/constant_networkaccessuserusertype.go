package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessUserUserType string

const (
	NetworkaccessUserUserType_Guest  NetworkaccessUserUserType = "guest"
	NetworkaccessUserUserType_Member NetworkaccessUserUserType = "member"
)

func PossibleValuesForNetworkaccessUserUserType() []string {
	return []string{
		string(NetworkaccessUserUserType_Guest),
		string(NetworkaccessUserUserType_Member),
	}
}

func (s *NetworkaccessUserUserType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessUserUserType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessUserUserType(input string) (*NetworkaccessUserUserType, error) {
	vals := map[string]NetworkaccessUserUserType{
		"guest":  NetworkaccessUserUserType_Guest,
		"member": NetworkaccessUserUserType_Member,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessUserUserType(input)
	return &out, nil
}
