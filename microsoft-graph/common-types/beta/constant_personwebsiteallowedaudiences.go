package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PersonWebsiteAllowedAudiences string

const (
	PersonWebsiteAllowedAudiences_Contacts               PersonWebsiteAllowedAudiences = "contacts"
	PersonWebsiteAllowedAudiences_Everyone               PersonWebsiteAllowedAudiences = "everyone"
	PersonWebsiteAllowedAudiences_Family                 PersonWebsiteAllowedAudiences = "family"
	PersonWebsiteAllowedAudiences_FederatedOrganizations PersonWebsiteAllowedAudiences = "federatedOrganizations"
	PersonWebsiteAllowedAudiences_GroupMembers           PersonWebsiteAllowedAudiences = "groupMembers"
	PersonWebsiteAllowedAudiences_Me                     PersonWebsiteAllowedAudiences = "me"
	PersonWebsiteAllowedAudiences_Organization           PersonWebsiteAllowedAudiences = "organization"
)

func PossibleValuesForPersonWebsiteAllowedAudiences() []string {
	return []string{
		string(PersonWebsiteAllowedAudiences_Contacts),
		string(PersonWebsiteAllowedAudiences_Everyone),
		string(PersonWebsiteAllowedAudiences_Family),
		string(PersonWebsiteAllowedAudiences_FederatedOrganizations),
		string(PersonWebsiteAllowedAudiences_GroupMembers),
		string(PersonWebsiteAllowedAudiences_Me),
		string(PersonWebsiteAllowedAudiences_Organization),
	}
}

func (s *PersonWebsiteAllowedAudiences) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePersonWebsiteAllowedAudiences(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePersonWebsiteAllowedAudiences(input string) (*PersonWebsiteAllowedAudiences, error) {
	vals := map[string]PersonWebsiteAllowedAudiences{
		"contacts":               PersonWebsiteAllowedAudiences_Contacts,
		"everyone":               PersonWebsiteAllowedAudiences_Everyone,
		"family":                 PersonWebsiteAllowedAudiences_Family,
		"federatedorganizations": PersonWebsiteAllowedAudiences_FederatedOrganizations,
		"groupmembers":           PersonWebsiteAllowedAudiences_GroupMembers,
		"me":                     PersonWebsiteAllowedAudiences_Me,
		"organization":           PersonWebsiteAllowedAudiences_Organization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PersonWebsiteAllowedAudiences(input)
	return &out, nil
}
