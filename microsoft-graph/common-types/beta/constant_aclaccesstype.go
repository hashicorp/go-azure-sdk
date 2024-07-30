package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AclAccessType string

const (
	AclAccessType_Deny  AclAccessType = "deny"
	AclAccessType_Grant AclAccessType = "grant"
)

func PossibleValuesForAclAccessType() []string {
	return []string{
		string(AclAccessType_Deny),
		string(AclAccessType_Grant),
	}
}

func (s *AclAccessType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAclAccessType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAclAccessType(input string) (*AclAccessType, error) {
	vals := map[string]AclAccessType{
		"deny":  AclAccessType_Deny,
		"grant": AclAccessType_Grant,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AclAccessType(input)
	return &out, nil
}
