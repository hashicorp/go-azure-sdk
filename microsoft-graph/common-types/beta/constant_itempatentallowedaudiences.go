package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ItemPatentAllowedAudiences string

const (
	ItemPatentAllowedAudiences_Contacts               ItemPatentAllowedAudiences = "contacts"
	ItemPatentAllowedAudiences_Everyone               ItemPatentAllowedAudiences = "everyone"
	ItemPatentAllowedAudiences_Family                 ItemPatentAllowedAudiences = "family"
	ItemPatentAllowedAudiences_FederatedOrganizations ItemPatentAllowedAudiences = "federatedOrganizations"
	ItemPatentAllowedAudiences_GroupMembers           ItemPatentAllowedAudiences = "groupMembers"
	ItemPatentAllowedAudiences_Me                     ItemPatentAllowedAudiences = "me"
	ItemPatentAllowedAudiences_Organization           ItemPatentAllowedAudiences = "organization"
)

func PossibleValuesForItemPatentAllowedAudiences() []string {
	return []string{
		string(ItemPatentAllowedAudiences_Contacts),
		string(ItemPatentAllowedAudiences_Everyone),
		string(ItemPatentAllowedAudiences_Family),
		string(ItemPatentAllowedAudiences_FederatedOrganizations),
		string(ItemPatentAllowedAudiences_GroupMembers),
		string(ItemPatentAllowedAudiences_Me),
		string(ItemPatentAllowedAudiences_Organization),
	}
}

func (s *ItemPatentAllowedAudiences) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseItemPatentAllowedAudiences(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseItemPatentAllowedAudiences(input string) (*ItemPatentAllowedAudiences, error) {
	vals := map[string]ItemPatentAllowedAudiences{
		"contacts":               ItemPatentAllowedAudiences_Contacts,
		"everyone":               ItemPatentAllowedAudiences_Everyone,
		"family":                 ItemPatentAllowedAudiences_Family,
		"federatedorganizations": ItemPatentAllowedAudiences_FederatedOrganizations,
		"groupmembers":           ItemPatentAllowedAudiences_GroupMembers,
		"me":                     ItemPatentAllowedAudiences_Me,
		"organization":           ItemPatentAllowedAudiences_Organization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ItemPatentAllowedAudiences(input)
	return &out, nil
}
