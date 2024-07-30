package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PersonInterestAllowedAudiences string

const (
	PersonInterestAllowedAudiences_Contacts               PersonInterestAllowedAudiences = "contacts"
	PersonInterestAllowedAudiences_Everyone               PersonInterestAllowedAudiences = "everyone"
	PersonInterestAllowedAudiences_Family                 PersonInterestAllowedAudiences = "family"
	PersonInterestAllowedAudiences_FederatedOrganizations PersonInterestAllowedAudiences = "federatedOrganizations"
	PersonInterestAllowedAudiences_GroupMembers           PersonInterestAllowedAudiences = "groupMembers"
	PersonInterestAllowedAudiences_Me                     PersonInterestAllowedAudiences = "me"
	PersonInterestAllowedAudiences_Organization           PersonInterestAllowedAudiences = "organization"
)

func PossibleValuesForPersonInterestAllowedAudiences() []string {
	return []string{
		string(PersonInterestAllowedAudiences_Contacts),
		string(PersonInterestAllowedAudiences_Everyone),
		string(PersonInterestAllowedAudiences_Family),
		string(PersonInterestAllowedAudiences_FederatedOrganizations),
		string(PersonInterestAllowedAudiences_GroupMembers),
		string(PersonInterestAllowedAudiences_Me),
		string(PersonInterestAllowedAudiences_Organization),
	}
}

func (s *PersonInterestAllowedAudiences) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePersonInterestAllowedAudiences(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePersonInterestAllowedAudiences(input string) (*PersonInterestAllowedAudiences, error) {
	vals := map[string]PersonInterestAllowedAudiences{
		"contacts":               PersonInterestAllowedAudiences_Contacts,
		"everyone":               PersonInterestAllowedAudiences_Everyone,
		"family":                 PersonInterestAllowedAudiences_Family,
		"federatedorganizations": PersonInterestAllowedAudiences_FederatedOrganizations,
		"groupmembers":           PersonInterestAllowedAudiences_GroupMembers,
		"me":                     PersonInterestAllowedAudiences_Me,
		"organization":           PersonInterestAllowedAudiences_Organization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PersonInterestAllowedAudiences(input)
	return &out, nil
}
