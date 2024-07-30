package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PersonCertificationAllowedAudiences string

const (
	PersonCertificationAllowedAudiences_Contacts               PersonCertificationAllowedAudiences = "contacts"
	PersonCertificationAllowedAudiences_Everyone               PersonCertificationAllowedAudiences = "everyone"
	PersonCertificationAllowedAudiences_Family                 PersonCertificationAllowedAudiences = "family"
	PersonCertificationAllowedAudiences_FederatedOrganizations PersonCertificationAllowedAudiences = "federatedOrganizations"
	PersonCertificationAllowedAudiences_GroupMembers           PersonCertificationAllowedAudiences = "groupMembers"
	PersonCertificationAllowedAudiences_Me                     PersonCertificationAllowedAudiences = "me"
	PersonCertificationAllowedAudiences_Organization           PersonCertificationAllowedAudiences = "organization"
)

func PossibleValuesForPersonCertificationAllowedAudiences() []string {
	return []string{
		string(PersonCertificationAllowedAudiences_Contacts),
		string(PersonCertificationAllowedAudiences_Everyone),
		string(PersonCertificationAllowedAudiences_Family),
		string(PersonCertificationAllowedAudiences_FederatedOrganizations),
		string(PersonCertificationAllowedAudiences_GroupMembers),
		string(PersonCertificationAllowedAudiences_Me),
		string(PersonCertificationAllowedAudiences_Organization),
	}
}

func (s *PersonCertificationAllowedAudiences) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePersonCertificationAllowedAudiences(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePersonCertificationAllowedAudiences(input string) (*PersonCertificationAllowedAudiences, error) {
	vals := map[string]PersonCertificationAllowedAudiences{
		"contacts":               PersonCertificationAllowedAudiences_Contacts,
		"everyone":               PersonCertificationAllowedAudiences_Everyone,
		"family":                 PersonCertificationAllowedAudiences_Family,
		"federatedorganizations": PersonCertificationAllowedAudiences_FederatedOrganizations,
		"groupmembers":           PersonCertificationAllowedAudiences_GroupMembers,
		"me":                     PersonCertificationAllowedAudiences_Me,
		"organization":           PersonCertificationAllowedAudiences_Organization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PersonCertificationAllowedAudiences(input)
	return &out, nil
}
