package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PersonAnnualEventAllowedAudiences string

const (
	PersonAnnualEventAllowedAudiences_Contacts               PersonAnnualEventAllowedAudiences = "contacts"
	PersonAnnualEventAllowedAudiences_Everyone               PersonAnnualEventAllowedAudiences = "everyone"
	PersonAnnualEventAllowedAudiences_Family                 PersonAnnualEventAllowedAudiences = "family"
	PersonAnnualEventAllowedAudiences_FederatedOrganizations PersonAnnualEventAllowedAudiences = "federatedOrganizations"
	PersonAnnualEventAllowedAudiences_GroupMembers           PersonAnnualEventAllowedAudiences = "groupMembers"
	PersonAnnualEventAllowedAudiences_Me                     PersonAnnualEventAllowedAudiences = "me"
	PersonAnnualEventAllowedAudiences_Organization           PersonAnnualEventAllowedAudiences = "organization"
)

func PossibleValuesForPersonAnnualEventAllowedAudiences() []string {
	return []string{
		string(PersonAnnualEventAllowedAudiences_Contacts),
		string(PersonAnnualEventAllowedAudiences_Everyone),
		string(PersonAnnualEventAllowedAudiences_Family),
		string(PersonAnnualEventAllowedAudiences_FederatedOrganizations),
		string(PersonAnnualEventAllowedAudiences_GroupMembers),
		string(PersonAnnualEventAllowedAudiences_Me),
		string(PersonAnnualEventAllowedAudiences_Organization),
	}
}

func (s *PersonAnnualEventAllowedAudiences) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePersonAnnualEventAllowedAudiences(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePersonAnnualEventAllowedAudiences(input string) (*PersonAnnualEventAllowedAudiences, error) {
	vals := map[string]PersonAnnualEventAllowedAudiences{
		"contacts":               PersonAnnualEventAllowedAudiences_Contacts,
		"everyone":               PersonAnnualEventAllowedAudiences_Everyone,
		"family":                 PersonAnnualEventAllowedAudiences_Family,
		"federatedorganizations": PersonAnnualEventAllowedAudiences_FederatedOrganizations,
		"groupmembers":           PersonAnnualEventAllowedAudiences_GroupMembers,
		"me":                     PersonAnnualEventAllowedAudiences_Me,
		"organization":           PersonAnnualEventAllowedAudiences_Organization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PersonAnnualEventAllowedAudiences(input)
	return &out, nil
}
