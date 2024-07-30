package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnUserCreateStartExternalUsersSelfServiceSignUpUserTypeToCreate string

const (
	OnUserCreateStartExternalUsersSelfServiceSignUpUserTypeToCreate_Guest  OnUserCreateStartExternalUsersSelfServiceSignUpUserTypeToCreate = "guest"
	OnUserCreateStartExternalUsersSelfServiceSignUpUserTypeToCreate_Member OnUserCreateStartExternalUsersSelfServiceSignUpUserTypeToCreate = "member"
)

func PossibleValuesForOnUserCreateStartExternalUsersSelfServiceSignUpUserTypeToCreate() []string {
	return []string{
		string(OnUserCreateStartExternalUsersSelfServiceSignUpUserTypeToCreate_Guest),
		string(OnUserCreateStartExternalUsersSelfServiceSignUpUserTypeToCreate_Member),
	}
}

func (s *OnUserCreateStartExternalUsersSelfServiceSignUpUserTypeToCreate) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOnUserCreateStartExternalUsersSelfServiceSignUpUserTypeToCreate(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOnUserCreateStartExternalUsersSelfServiceSignUpUserTypeToCreate(input string) (*OnUserCreateStartExternalUsersSelfServiceSignUpUserTypeToCreate, error) {
	vals := map[string]OnUserCreateStartExternalUsersSelfServiceSignUpUserTypeToCreate{
		"guest":  OnUserCreateStartExternalUsersSelfServiceSignUpUserTypeToCreate_Guest,
		"member": OnUserCreateStartExternalUsersSelfServiceSignUpUserTypeToCreate_Member,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OnUserCreateStartExternalUsersSelfServiceSignUpUserTypeToCreate(input)
	return &out, nil
}
