package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserSecurityStateEmailRole string

const (
	UserSecurityStateEmailRole_Recipient UserSecurityStateEmailRole = "recipient"
	UserSecurityStateEmailRole_Sender    UserSecurityStateEmailRole = "sender"
	UserSecurityStateEmailRole_Unknown   UserSecurityStateEmailRole = "unknown"
)

func PossibleValuesForUserSecurityStateEmailRole() []string {
	return []string{
		string(UserSecurityStateEmailRole_Recipient),
		string(UserSecurityStateEmailRole_Sender),
		string(UserSecurityStateEmailRole_Unknown),
	}
}

func (s *UserSecurityStateEmailRole) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserSecurityStateEmailRole(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserSecurityStateEmailRole(input string) (*UserSecurityStateEmailRole, error) {
	vals := map[string]UserSecurityStateEmailRole{
		"recipient": UserSecurityStateEmailRole_Recipient,
		"sender":    UserSecurityStateEmailRole_Sender,
		"unknown":   UserSecurityStateEmailRole_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserSecurityStateEmailRole(input)
	return &out, nil
}
