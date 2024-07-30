package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserRegistrationFeatureSummaryUserRoles string

const (
	UserRegistrationFeatureSummaryUserRoles_Admin           UserRegistrationFeatureSummaryUserRoles = "admin"
	UserRegistrationFeatureSummaryUserRoles_All             UserRegistrationFeatureSummaryUserRoles = "all"
	UserRegistrationFeatureSummaryUserRoles_PrivilegedAdmin UserRegistrationFeatureSummaryUserRoles = "privilegedAdmin"
	UserRegistrationFeatureSummaryUserRoles_User            UserRegistrationFeatureSummaryUserRoles = "user"
)

func PossibleValuesForUserRegistrationFeatureSummaryUserRoles() []string {
	return []string{
		string(UserRegistrationFeatureSummaryUserRoles_Admin),
		string(UserRegistrationFeatureSummaryUserRoles_All),
		string(UserRegistrationFeatureSummaryUserRoles_PrivilegedAdmin),
		string(UserRegistrationFeatureSummaryUserRoles_User),
	}
}

func (s *UserRegistrationFeatureSummaryUserRoles) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserRegistrationFeatureSummaryUserRoles(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserRegistrationFeatureSummaryUserRoles(input string) (*UserRegistrationFeatureSummaryUserRoles, error) {
	vals := map[string]UserRegistrationFeatureSummaryUserRoles{
		"admin":           UserRegistrationFeatureSummaryUserRoles_Admin,
		"all":             UserRegistrationFeatureSummaryUserRoles_All,
		"privilegedadmin": UserRegistrationFeatureSummaryUserRoles_PrivilegedAdmin,
		"user":            UserRegistrationFeatureSummaryUserRoles_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserRegistrationFeatureSummaryUserRoles(input)
	return &out, nil
}
