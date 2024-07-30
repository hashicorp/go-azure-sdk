package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ItemAddressAllowedAudiences string

const (
	ItemAddressAllowedAudiences_Contacts               ItemAddressAllowedAudiences = "contacts"
	ItemAddressAllowedAudiences_Everyone               ItemAddressAllowedAudiences = "everyone"
	ItemAddressAllowedAudiences_Family                 ItemAddressAllowedAudiences = "family"
	ItemAddressAllowedAudiences_FederatedOrganizations ItemAddressAllowedAudiences = "federatedOrganizations"
	ItemAddressAllowedAudiences_GroupMembers           ItemAddressAllowedAudiences = "groupMembers"
	ItemAddressAllowedAudiences_Me                     ItemAddressAllowedAudiences = "me"
	ItemAddressAllowedAudiences_Organization           ItemAddressAllowedAudiences = "organization"
)

func PossibleValuesForItemAddressAllowedAudiences() []string {
	return []string{
		string(ItemAddressAllowedAudiences_Contacts),
		string(ItemAddressAllowedAudiences_Everyone),
		string(ItemAddressAllowedAudiences_Family),
		string(ItemAddressAllowedAudiences_FederatedOrganizations),
		string(ItemAddressAllowedAudiences_GroupMembers),
		string(ItemAddressAllowedAudiences_Me),
		string(ItemAddressAllowedAudiences_Organization),
	}
}

func (s *ItemAddressAllowedAudiences) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseItemAddressAllowedAudiences(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseItemAddressAllowedAudiences(input string) (*ItemAddressAllowedAudiences, error) {
	vals := map[string]ItemAddressAllowedAudiences{
		"contacts":               ItemAddressAllowedAudiences_Contacts,
		"everyone":               ItemAddressAllowedAudiences_Everyone,
		"family":                 ItemAddressAllowedAudiences_Family,
		"federatedorganizations": ItemAddressAllowedAudiences_FederatedOrganizations,
		"groupmembers":           ItemAddressAllowedAudiences_GroupMembers,
		"me":                     ItemAddressAllowedAudiences_Me,
		"organization":           ItemAddressAllowedAudiences_Organization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ItemAddressAllowedAudiences(input)
	return &out, nil
}
