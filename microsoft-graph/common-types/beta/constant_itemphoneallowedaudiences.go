package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ItemPhoneAllowedAudiences string

const (
	ItemPhoneAllowedAudiences_Contacts               ItemPhoneAllowedAudiences = "contacts"
	ItemPhoneAllowedAudiences_Everyone               ItemPhoneAllowedAudiences = "everyone"
	ItemPhoneAllowedAudiences_Family                 ItemPhoneAllowedAudiences = "family"
	ItemPhoneAllowedAudiences_FederatedOrganizations ItemPhoneAllowedAudiences = "federatedOrganizations"
	ItemPhoneAllowedAudiences_GroupMembers           ItemPhoneAllowedAudiences = "groupMembers"
	ItemPhoneAllowedAudiences_Me                     ItemPhoneAllowedAudiences = "me"
	ItemPhoneAllowedAudiences_Organization           ItemPhoneAllowedAudiences = "organization"
)

func PossibleValuesForItemPhoneAllowedAudiences() []string {
	return []string{
		string(ItemPhoneAllowedAudiences_Contacts),
		string(ItemPhoneAllowedAudiences_Everyone),
		string(ItemPhoneAllowedAudiences_Family),
		string(ItemPhoneAllowedAudiences_FederatedOrganizations),
		string(ItemPhoneAllowedAudiences_GroupMembers),
		string(ItemPhoneAllowedAudiences_Me),
		string(ItemPhoneAllowedAudiences_Organization),
	}
}

func (s *ItemPhoneAllowedAudiences) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseItemPhoneAllowedAudiences(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseItemPhoneAllowedAudiences(input string) (*ItemPhoneAllowedAudiences, error) {
	vals := map[string]ItemPhoneAllowedAudiences{
		"contacts":               ItemPhoneAllowedAudiences_Contacts,
		"everyone":               ItemPhoneAllowedAudiences_Everyone,
		"family":                 ItemPhoneAllowedAudiences_Family,
		"federatedorganizations": ItemPhoneAllowedAudiences_FederatedOrganizations,
		"groupmembers":           ItemPhoneAllowedAudiences_GroupMembers,
		"me":                     ItemPhoneAllowedAudiences_Me,
		"organization":           ItemPhoneAllowedAudiences_Organization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ItemPhoneAllowedAudiences(input)
	return &out, nil
}
