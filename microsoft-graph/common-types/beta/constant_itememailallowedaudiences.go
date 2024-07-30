package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ItemEmailAllowedAudiences string

const (
	ItemEmailAllowedAudiences_Contacts               ItemEmailAllowedAudiences = "contacts"
	ItemEmailAllowedAudiences_Everyone               ItemEmailAllowedAudiences = "everyone"
	ItemEmailAllowedAudiences_Family                 ItemEmailAllowedAudiences = "family"
	ItemEmailAllowedAudiences_FederatedOrganizations ItemEmailAllowedAudiences = "federatedOrganizations"
	ItemEmailAllowedAudiences_GroupMembers           ItemEmailAllowedAudiences = "groupMembers"
	ItemEmailAllowedAudiences_Me                     ItemEmailAllowedAudiences = "me"
	ItemEmailAllowedAudiences_Organization           ItemEmailAllowedAudiences = "organization"
)

func PossibleValuesForItemEmailAllowedAudiences() []string {
	return []string{
		string(ItemEmailAllowedAudiences_Contacts),
		string(ItemEmailAllowedAudiences_Everyone),
		string(ItemEmailAllowedAudiences_Family),
		string(ItemEmailAllowedAudiences_FederatedOrganizations),
		string(ItemEmailAllowedAudiences_GroupMembers),
		string(ItemEmailAllowedAudiences_Me),
		string(ItemEmailAllowedAudiences_Organization),
	}
}

func (s *ItemEmailAllowedAudiences) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseItemEmailAllowedAudiences(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseItemEmailAllowedAudiences(input string) (*ItemEmailAllowedAudiences, error) {
	vals := map[string]ItemEmailAllowedAudiences{
		"contacts":               ItemEmailAllowedAudiences_Contacts,
		"everyone":               ItemEmailAllowedAudiences_Everyone,
		"family":                 ItemEmailAllowedAudiences_Family,
		"federatedorganizations": ItemEmailAllowedAudiences_FederatedOrganizations,
		"groupmembers":           ItemEmailAllowedAudiences_GroupMembers,
		"me":                     ItemEmailAllowedAudiences_Me,
		"organization":           ItemEmailAllowedAudiences_Organization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ItemEmailAllowedAudiences(input)
	return &out, nil
}
