package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserRegistrationFeatureSummaryUserTypes string

const (
	UserRegistrationFeatureSummaryUserTypes_All    UserRegistrationFeatureSummaryUserTypes = "all"
	UserRegistrationFeatureSummaryUserTypes_Guest  UserRegistrationFeatureSummaryUserTypes = "guest"
	UserRegistrationFeatureSummaryUserTypes_Member UserRegistrationFeatureSummaryUserTypes = "member"
)

func PossibleValuesForUserRegistrationFeatureSummaryUserTypes() []string {
	return []string{
		string(UserRegistrationFeatureSummaryUserTypes_All),
		string(UserRegistrationFeatureSummaryUserTypes_Guest),
		string(UserRegistrationFeatureSummaryUserTypes_Member),
	}
}

func (s *UserRegistrationFeatureSummaryUserTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserRegistrationFeatureSummaryUserTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserRegistrationFeatureSummaryUserTypes(input string) (*UserRegistrationFeatureSummaryUserTypes, error) {
	vals := map[string]UserRegistrationFeatureSummaryUserTypes{
		"all":    UserRegistrationFeatureSummaryUserTypes_All,
		"guest":  UserRegistrationFeatureSummaryUserTypes_Guest,
		"member": UserRegistrationFeatureSummaryUserTypes_Member,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserRegistrationFeatureSummaryUserTypes(input)
	return &out, nil
}
