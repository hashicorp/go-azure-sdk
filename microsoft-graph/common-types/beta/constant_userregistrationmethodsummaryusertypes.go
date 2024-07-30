package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserRegistrationMethodSummaryUserTypes string

const (
	UserRegistrationMethodSummaryUserTypes_All    UserRegistrationMethodSummaryUserTypes = "all"
	UserRegistrationMethodSummaryUserTypes_Guest  UserRegistrationMethodSummaryUserTypes = "guest"
	UserRegistrationMethodSummaryUserTypes_Member UserRegistrationMethodSummaryUserTypes = "member"
)

func PossibleValuesForUserRegistrationMethodSummaryUserTypes() []string {
	return []string{
		string(UserRegistrationMethodSummaryUserTypes_All),
		string(UserRegistrationMethodSummaryUserTypes_Guest),
		string(UserRegistrationMethodSummaryUserTypes_Member),
	}
}

func (s *UserRegistrationMethodSummaryUserTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserRegistrationMethodSummaryUserTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserRegistrationMethodSummaryUserTypes(input string) (*UserRegistrationMethodSummaryUserTypes, error) {
	vals := map[string]UserRegistrationMethodSummaryUserTypes{
		"all":    UserRegistrationMethodSummaryUserTypes_All,
		"guest":  UserRegistrationMethodSummaryUserTypes_Guest,
		"member": UserRegistrationMethodSummaryUserTypes_Member,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserRegistrationMethodSummaryUserTypes(input)
	return &out, nil
}
