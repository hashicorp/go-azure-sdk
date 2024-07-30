package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserRegistrationDetailsUserType string

const (
	UserRegistrationDetailsUserType_Guest  UserRegistrationDetailsUserType = "guest"
	UserRegistrationDetailsUserType_Member UserRegistrationDetailsUserType = "member"
)

func PossibleValuesForUserRegistrationDetailsUserType() []string {
	return []string{
		string(UserRegistrationDetailsUserType_Guest),
		string(UserRegistrationDetailsUserType_Member),
	}
}

func (s *UserRegistrationDetailsUserType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserRegistrationDetailsUserType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserRegistrationDetailsUserType(input string) (*UserRegistrationDetailsUserType, error) {
	vals := map[string]UserRegistrationDetailsUserType{
		"guest":  UserRegistrationDetailsUserType_Guest,
		"member": UserRegistrationDetailsUserType_Member,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserRegistrationDetailsUserType(input)
	return &out, nil
}
