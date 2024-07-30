package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PersonAwardAllowedAudiences string

const (
	PersonAwardAllowedAudiences_Contacts               PersonAwardAllowedAudiences = "contacts"
	PersonAwardAllowedAudiences_Everyone               PersonAwardAllowedAudiences = "everyone"
	PersonAwardAllowedAudiences_Family                 PersonAwardAllowedAudiences = "family"
	PersonAwardAllowedAudiences_FederatedOrganizations PersonAwardAllowedAudiences = "federatedOrganizations"
	PersonAwardAllowedAudiences_GroupMembers           PersonAwardAllowedAudiences = "groupMembers"
	PersonAwardAllowedAudiences_Me                     PersonAwardAllowedAudiences = "me"
	PersonAwardAllowedAudiences_Organization           PersonAwardAllowedAudiences = "organization"
)

func PossibleValuesForPersonAwardAllowedAudiences() []string {
	return []string{
		string(PersonAwardAllowedAudiences_Contacts),
		string(PersonAwardAllowedAudiences_Everyone),
		string(PersonAwardAllowedAudiences_Family),
		string(PersonAwardAllowedAudiences_FederatedOrganizations),
		string(PersonAwardAllowedAudiences_GroupMembers),
		string(PersonAwardAllowedAudiences_Me),
		string(PersonAwardAllowedAudiences_Organization),
	}
}

func (s *PersonAwardAllowedAudiences) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePersonAwardAllowedAudiences(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePersonAwardAllowedAudiences(input string) (*PersonAwardAllowedAudiences, error) {
	vals := map[string]PersonAwardAllowedAudiences{
		"contacts":               PersonAwardAllowedAudiences_Contacts,
		"everyone":               PersonAwardAllowedAudiences_Everyone,
		"family":                 PersonAwardAllowedAudiences_Family,
		"federatedorganizations": PersonAwardAllowedAudiences_FederatedOrganizations,
		"groupmembers":           PersonAwardAllowedAudiences_GroupMembers,
		"me":                     PersonAwardAllowedAudiences_Me,
		"organization":           PersonAwardAllowedAudiences_Organization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PersonAwardAllowedAudiences(input)
	return &out, nil
}
