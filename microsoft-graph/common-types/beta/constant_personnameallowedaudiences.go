package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PersonNameAllowedAudiences string

const (
	PersonNameAllowedAudiences_Contacts               PersonNameAllowedAudiences = "contacts"
	PersonNameAllowedAudiences_Everyone               PersonNameAllowedAudiences = "everyone"
	PersonNameAllowedAudiences_Family                 PersonNameAllowedAudiences = "family"
	PersonNameAllowedAudiences_FederatedOrganizations PersonNameAllowedAudiences = "federatedOrganizations"
	PersonNameAllowedAudiences_GroupMembers           PersonNameAllowedAudiences = "groupMembers"
	PersonNameAllowedAudiences_Me                     PersonNameAllowedAudiences = "me"
	PersonNameAllowedAudiences_Organization           PersonNameAllowedAudiences = "organization"
)

func PossibleValuesForPersonNameAllowedAudiences() []string {
	return []string{
		string(PersonNameAllowedAudiences_Contacts),
		string(PersonNameAllowedAudiences_Everyone),
		string(PersonNameAllowedAudiences_Family),
		string(PersonNameAllowedAudiences_FederatedOrganizations),
		string(PersonNameAllowedAudiences_GroupMembers),
		string(PersonNameAllowedAudiences_Me),
		string(PersonNameAllowedAudiences_Organization),
	}
}

func (s *PersonNameAllowedAudiences) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePersonNameAllowedAudiences(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePersonNameAllowedAudiences(input string) (*PersonNameAllowedAudiences, error) {
	vals := map[string]PersonNameAllowedAudiences{
		"contacts":               PersonNameAllowedAudiences_Contacts,
		"everyone":               PersonNameAllowedAudiences_Everyone,
		"family":                 PersonNameAllowedAudiences_Family,
		"federatedorganizations": PersonNameAllowedAudiences_FederatedOrganizations,
		"groupmembers":           PersonNameAllowedAudiences_GroupMembers,
		"me":                     PersonNameAllowedAudiences_Me,
		"organization":           PersonNameAllowedAudiences_Organization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PersonNameAllowedAudiences(input)
	return &out, nil
}
