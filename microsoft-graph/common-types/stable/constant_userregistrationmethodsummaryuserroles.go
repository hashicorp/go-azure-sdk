package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserRegistrationMethodSummaryUserRoles string

const (
	UserRegistrationMethodSummaryUserRoles_Admin           UserRegistrationMethodSummaryUserRoles = "admin"
	UserRegistrationMethodSummaryUserRoles_All             UserRegistrationMethodSummaryUserRoles = "all"
	UserRegistrationMethodSummaryUserRoles_PrivilegedAdmin UserRegistrationMethodSummaryUserRoles = "privilegedAdmin"
	UserRegistrationMethodSummaryUserRoles_User            UserRegistrationMethodSummaryUserRoles = "user"
)

func PossibleValuesForUserRegistrationMethodSummaryUserRoles() []string {
	return []string{
		string(UserRegistrationMethodSummaryUserRoles_Admin),
		string(UserRegistrationMethodSummaryUserRoles_All),
		string(UserRegistrationMethodSummaryUserRoles_PrivilegedAdmin),
		string(UserRegistrationMethodSummaryUserRoles_User),
	}
}

func (s *UserRegistrationMethodSummaryUserRoles) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserRegistrationMethodSummaryUserRoles(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserRegistrationMethodSummaryUserRoles(input string) (*UserRegistrationMethodSummaryUserRoles, error) {
	vals := map[string]UserRegistrationMethodSummaryUserRoles{
		"admin":           UserRegistrationMethodSummaryUserRoles_Admin,
		"all":             UserRegistrationMethodSummaryUserRoles_All,
		"privilegedadmin": UserRegistrationMethodSummaryUserRoles_PrivilegedAdmin,
		"user":            UserRegistrationMethodSummaryUserRoles_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserRegistrationMethodSummaryUserRoles(input)
	return &out, nil
}
