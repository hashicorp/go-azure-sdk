package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LanguageProficiencyAllowedAudiences string

const (
	LanguageProficiencyAllowedAudiences_Contacts               LanguageProficiencyAllowedAudiences = "contacts"
	LanguageProficiencyAllowedAudiences_Everyone               LanguageProficiencyAllowedAudiences = "everyone"
	LanguageProficiencyAllowedAudiences_Family                 LanguageProficiencyAllowedAudiences = "family"
	LanguageProficiencyAllowedAudiences_FederatedOrganizations LanguageProficiencyAllowedAudiences = "federatedOrganizations"
	LanguageProficiencyAllowedAudiences_GroupMembers           LanguageProficiencyAllowedAudiences = "groupMembers"
	LanguageProficiencyAllowedAudiences_Me                     LanguageProficiencyAllowedAudiences = "me"
	LanguageProficiencyAllowedAudiences_Organization           LanguageProficiencyAllowedAudiences = "organization"
)

func PossibleValuesForLanguageProficiencyAllowedAudiences() []string {
	return []string{
		string(LanguageProficiencyAllowedAudiences_Contacts),
		string(LanguageProficiencyAllowedAudiences_Everyone),
		string(LanguageProficiencyAllowedAudiences_Family),
		string(LanguageProficiencyAllowedAudiences_FederatedOrganizations),
		string(LanguageProficiencyAllowedAudiences_GroupMembers),
		string(LanguageProficiencyAllowedAudiences_Me),
		string(LanguageProficiencyAllowedAudiences_Organization),
	}
}

func (s *LanguageProficiencyAllowedAudiences) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLanguageProficiencyAllowedAudiences(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLanguageProficiencyAllowedAudiences(input string) (*LanguageProficiencyAllowedAudiences, error) {
	vals := map[string]LanguageProficiencyAllowedAudiences{
		"contacts":               LanguageProficiencyAllowedAudiences_Contacts,
		"everyone":               LanguageProficiencyAllowedAudiences_Everyone,
		"family":                 LanguageProficiencyAllowedAudiences_Family,
		"federatedorganizations": LanguageProficiencyAllowedAudiences_FederatedOrganizations,
		"groupmembers":           LanguageProficiencyAllowedAudiences_GroupMembers,
		"me":                     LanguageProficiencyAllowedAudiences_Me,
		"organization":           LanguageProficiencyAllowedAudiences_Organization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LanguageProficiencyAllowedAudiences(input)
	return &out, nil
}
