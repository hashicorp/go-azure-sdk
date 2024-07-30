package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ItemFacetAllowedAudiences string

const (
	ItemFacetAllowedAudiences_Contacts               ItemFacetAllowedAudiences = "contacts"
	ItemFacetAllowedAudiences_Everyone               ItemFacetAllowedAudiences = "everyone"
	ItemFacetAllowedAudiences_Family                 ItemFacetAllowedAudiences = "family"
	ItemFacetAllowedAudiences_FederatedOrganizations ItemFacetAllowedAudiences = "federatedOrganizations"
	ItemFacetAllowedAudiences_GroupMembers           ItemFacetAllowedAudiences = "groupMembers"
	ItemFacetAllowedAudiences_Me                     ItemFacetAllowedAudiences = "me"
	ItemFacetAllowedAudiences_Organization           ItemFacetAllowedAudiences = "organization"
)

func PossibleValuesForItemFacetAllowedAudiences() []string {
	return []string{
		string(ItemFacetAllowedAudiences_Contacts),
		string(ItemFacetAllowedAudiences_Everyone),
		string(ItemFacetAllowedAudiences_Family),
		string(ItemFacetAllowedAudiences_FederatedOrganizations),
		string(ItemFacetAllowedAudiences_GroupMembers),
		string(ItemFacetAllowedAudiences_Me),
		string(ItemFacetAllowedAudiences_Organization),
	}
}

func (s *ItemFacetAllowedAudiences) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseItemFacetAllowedAudiences(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseItemFacetAllowedAudiences(input string) (*ItemFacetAllowedAudiences, error) {
	vals := map[string]ItemFacetAllowedAudiences{
		"contacts":               ItemFacetAllowedAudiences_Contacts,
		"everyone":               ItemFacetAllowedAudiences_Everyone,
		"family":                 ItemFacetAllowedAudiences_Family,
		"federatedorganizations": ItemFacetAllowedAudiences_FederatedOrganizations,
		"groupmembers":           ItemFacetAllowedAudiences_GroupMembers,
		"me":                     ItemFacetAllowedAudiences_Me,
		"organization":           ItemFacetAllowedAudiences_Organization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ItemFacetAllowedAudiences(input)
	return &out, nil
}
