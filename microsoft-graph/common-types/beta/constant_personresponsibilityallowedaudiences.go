package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PersonResponsibilityAllowedAudiences string

const (
	PersonResponsibilityAllowedAudiences_Contacts               PersonResponsibilityAllowedAudiences = "contacts"
	PersonResponsibilityAllowedAudiences_Everyone               PersonResponsibilityAllowedAudiences = "everyone"
	PersonResponsibilityAllowedAudiences_Family                 PersonResponsibilityAllowedAudiences = "family"
	PersonResponsibilityAllowedAudiences_FederatedOrganizations PersonResponsibilityAllowedAudiences = "federatedOrganizations"
	PersonResponsibilityAllowedAudiences_GroupMembers           PersonResponsibilityAllowedAudiences = "groupMembers"
	PersonResponsibilityAllowedAudiences_Me                     PersonResponsibilityAllowedAudiences = "me"
	PersonResponsibilityAllowedAudiences_Organization           PersonResponsibilityAllowedAudiences = "organization"
)

func PossibleValuesForPersonResponsibilityAllowedAudiences() []string {
	return []string{
		string(PersonResponsibilityAllowedAudiences_Contacts),
		string(PersonResponsibilityAllowedAudiences_Everyone),
		string(PersonResponsibilityAllowedAudiences_Family),
		string(PersonResponsibilityAllowedAudiences_FederatedOrganizations),
		string(PersonResponsibilityAllowedAudiences_GroupMembers),
		string(PersonResponsibilityAllowedAudiences_Me),
		string(PersonResponsibilityAllowedAudiences_Organization),
	}
}

func (s *PersonResponsibilityAllowedAudiences) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePersonResponsibilityAllowedAudiences(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePersonResponsibilityAllowedAudiences(input string) (*PersonResponsibilityAllowedAudiences, error) {
	vals := map[string]PersonResponsibilityAllowedAudiences{
		"contacts":               PersonResponsibilityAllowedAudiences_Contacts,
		"everyone":               PersonResponsibilityAllowedAudiences_Everyone,
		"family":                 PersonResponsibilityAllowedAudiences_Family,
		"federatedorganizations": PersonResponsibilityAllowedAudiences_FederatedOrganizations,
		"groupmembers":           PersonResponsibilityAllowedAudiences_GroupMembers,
		"me":                     PersonResponsibilityAllowedAudiences_Me,
		"organization":           PersonResponsibilityAllowedAudiences_Organization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PersonResponsibilityAllowedAudiences(input)
	return &out, nil
}
