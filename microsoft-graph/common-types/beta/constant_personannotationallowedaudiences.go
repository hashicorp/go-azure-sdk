package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PersonAnnotationAllowedAudiences string

const (
	PersonAnnotationAllowedAudiences_Contacts               PersonAnnotationAllowedAudiences = "contacts"
	PersonAnnotationAllowedAudiences_Everyone               PersonAnnotationAllowedAudiences = "everyone"
	PersonAnnotationAllowedAudiences_Family                 PersonAnnotationAllowedAudiences = "family"
	PersonAnnotationAllowedAudiences_FederatedOrganizations PersonAnnotationAllowedAudiences = "federatedOrganizations"
	PersonAnnotationAllowedAudiences_GroupMembers           PersonAnnotationAllowedAudiences = "groupMembers"
	PersonAnnotationAllowedAudiences_Me                     PersonAnnotationAllowedAudiences = "me"
	PersonAnnotationAllowedAudiences_Organization           PersonAnnotationAllowedAudiences = "organization"
)

func PossibleValuesForPersonAnnotationAllowedAudiences() []string {
	return []string{
		string(PersonAnnotationAllowedAudiences_Contacts),
		string(PersonAnnotationAllowedAudiences_Everyone),
		string(PersonAnnotationAllowedAudiences_Family),
		string(PersonAnnotationAllowedAudiences_FederatedOrganizations),
		string(PersonAnnotationAllowedAudiences_GroupMembers),
		string(PersonAnnotationAllowedAudiences_Me),
		string(PersonAnnotationAllowedAudiences_Organization),
	}
}

func (s *PersonAnnotationAllowedAudiences) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePersonAnnotationAllowedAudiences(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePersonAnnotationAllowedAudiences(input string) (*PersonAnnotationAllowedAudiences, error) {
	vals := map[string]PersonAnnotationAllowedAudiences{
		"contacts":               PersonAnnotationAllowedAudiences_Contacts,
		"everyone":               PersonAnnotationAllowedAudiences_Everyone,
		"family":                 PersonAnnotationAllowedAudiences_Family,
		"federatedorganizations": PersonAnnotationAllowedAudiences_FederatedOrganizations,
		"groupmembers":           PersonAnnotationAllowedAudiences_GroupMembers,
		"me":                     PersonAnnotationAllowedAudiences_Me,
		"organization":           PersonAnnotationAllowedAudiences_Organization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PersonAnnotationAllowedAudiences(input)
	return &out, nil
}
