package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ItemPublicationAllowedAudiences string

const (
	ItemPublicationAllowedAudiences_Contacts               ItemPublicationAllowedAudiences = "contacts"
	ItemPublicationAllowedAudiences_Everyone               ItemPublicationAllowedAudiences = "everyone"
	ItemPublicationAllowedAudiences_Family                 ItemPublicationAllowedAudiences = "family"
	ItemPublicationAllowedAudiences_FederatedOrganizations ItemPublicationAllowedAudiences = "federatedOrganizations"
	ItemPublicationAllowedAudiences_GroupMembers           ItemPublicationAllowedAudiences = "groupMembers"
	ItemPublicationAllowedAudiences_Me                     ItemPublicationAllowedAudiences = "me"
	ItemPublicationAllowedAudiences_Organization           ItemPublicationAllowedAudiences = "organization"
)

func PossibleValuesForItemPublicationAllowedAudiences() []string {
	return []string{
		string(ItemPublicationAllowedAudiences_Contacts),
		string(ItemPublicationAllowedAudiences_Everyone),
		string(ItemPublicationAllowedAudiences_Family),
		string(ItemPublicationAllowedAudiences_FederatedOrganizations),
		string(ItemPublicationAllowedAudiences_GroupMembers),
		string(ItemPublicationAllowedAudiences_Me),
		string(ItemPublicationAllowedAudiences_Organization),
	}
}

func (s *ItemPublicationAllowedAudiences) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseItemPublicationAllowedAudiences(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseItemPublicationAllowedAudiences(input string) (*ItemPublicationAllowedAudiences, error) {
	vals := map[string]ItemPublicationAllowedAudiences{
		"contacts":               ItemPublicationAllowedAudiences_Contacts,
		"everyone":               ItemPublicationAllowedAudiences_Everyone,
		"family":                 ItemPublicationAllowedAudiences_Family,
		"federatedorganizations": ItemPublicationAllowedAudiences_FederatedOrganizations,
		"groupmembers":           ItemPublicationAllowedAudiences_GroupMembers,
		"me":                     ItemPublicationAllowedAudiences_Me,
		"organization":           ItemPublicationAllowedAudiences_Organization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ItemPublicationAllowedAudiences(input)
	return &out, nil
}
