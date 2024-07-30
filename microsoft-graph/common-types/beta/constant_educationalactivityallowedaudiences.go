package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationalActivityAllowedAudiences string

const (
	EducationalActivityAllowedAudiences_Contacts               EducationalActivityAllowedAudiences = "contacts"
	EducationalActivityAllowedAudiences_Everyone               EducationalActivityAllowedAudiences = "everyone"
	EducationalActivityAllowedAudiences_Family                 EducationalActivityAllowedAudiences = "family"
	EducationalActivityAllowedAudiences_FederatedOrganizations EducationalActivityAllowedAudiences = "federatedOrganizations"
	EducationalActivityAllowedAudiences_GroupMembers           EducationalActivityAllowedAudiences = "groupMembers"
	EducationalActivityAllowedAudiences_Me                     EducationalActivityAllowedAudiences = "me"
	EducationalActivityAllowedAudiences_Organization           EducationalActivityAllowedAudiences = "organization"
)

func PossibleValuesForEducationalActivityAllowedAudiences() []string {
	return []string{
		string(EducationalActivityAllowedAudiences_Contacts),
		string(EducationalActivityAllowedAudiences_Everyone),
		string(EducationalActivityAllowedAudiences_Family),
		string(EducationalActivityAllowedAudiences_FederatedOrganizations),
		string(EducationalActivityAllowedAudiences_GroupMembers),
		string(EducationalActivityAllowedAudiences_Me),
		string(EducationalActivityAllowedAudiences_Organization),
	}
}

func (s *EducationalActivityAllowedAudiences) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEducationalActivityAllowedAudiences(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEducationalActivityAllowedAudiences(input string) (*EducationalActivityAllowedAudiences, error) {
	vals := map[string]EducationalActivityAllowedAudiences{
		"contacts":               EducationalActivityAllowedAudiences_Contacts,
		"everyone":               EducationalActivityAllowedAudiences_Everyone,
		"family":                 EducationalActivityAllowedAudiences_Family,
		"federatedorganizations": EducationalActivityAllowedAudiences_FederatedOrganizations,
		"groupmembers":           EducationalActivityAllowedAudiences_GroupMembers,
		"me":                     EducationalActivityAllowedAudiences_Me,
		"organization":           EducationalActivityAllowedAudiences_Organization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EducationalActivityAllowedAudiences(input)
	return &out, nil
}
