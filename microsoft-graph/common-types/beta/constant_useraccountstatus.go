package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserAccountStatus string

const (
	UserAccountStatus_Active    UserAccountStatus = "active"
	UserAccountStatus_Deleted   UserAccountStatus = "deleted"
	UserAccountStatus_Staged    UserAccountStatus = "staged"
	UserAccountStatus_Suspended UserAccountStatus = "suspended"
	UserAccountStatus_Unknown   UserAccountStatus = "unknown"
)

func PossibleValuesForUserAccountStatus() []string {
	return []string{
		string(UserAccountStatus_Active),
		string(UserAccountStatus_Deleted),
		string(UserAccountStatus_Staged),
		string(UserAccountStatus_Suspended),
		string(UserAccountStatus_Unknown),
	}
}

func (s *UserAccountStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserAccountStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserAccountStatus(input string) (*UserAccountStatus, error) {
	vals := map[string]UserAccountStatus{
		"active":    UserAccountStatus_Active,
		"deleted":   UserAccountStatus_Deleted,
		"staged":    UserAccountStatus_Staged,
		"suspended": UserAccountStatus_Suspended,
		"unknown":   UserAccountStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserAccountStatus(input)
	return &out, nil
}
