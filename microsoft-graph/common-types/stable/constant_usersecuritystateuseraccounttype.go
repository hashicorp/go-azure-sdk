package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserSecurityStateUserAccountType string

const (
	UserSecurityStateUserAccountType_Administrator UserSecurityStateUserAccountType = "administrator"
	UserSecurityStateUserAccountType_Power         UserSecurityStateUserAccountType = "power"
	UserSecurityStateUserAccountType_Standard      UserSecurityStateUserAccountType = "standard"
	UserSecurityStateUserAccountType_Unknown       UserSecurityStateUserAccountType = "unknown"
)

func PossibleValuesForUserSecurityStateUserAccountType() []string {
	return []string{
		string(UserSecurityStateUserAccountType_Administrator),
		string(UserSecurityStateUserAccountType_Power),
		string(UserSecurityStateUserAccountType_Standard),
		string(UserSecurityStateUserAccountType_Unknown),
	}
}

func (s *UserSecurityStateUserAccountType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserSecurityStateUserAccountType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserSecurityStateUserAccountType(input string) (*UserSecurityStateUserAccountType, error) {
	vals := map[string]UserSecurityStateUserAccountType{
		"administrator": UserSecurityStateUserAccountType_Administrator,
		"power":         UserSecurityStateUserAccountType_Power,
		"standard":      UserSecurityStateUserAccountType_Standard,
		"unknown":       UserSecurityStateUserAccountType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserSecurityStateUserAccountType(input)
	return &out, nil
}
