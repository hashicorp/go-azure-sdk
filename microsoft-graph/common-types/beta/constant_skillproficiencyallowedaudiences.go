package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SkillProficiencyAllowedAudiences string

const (
	SkillProficiencyAllowedAudiences_Contacts               SkillProficiencyAllowedAudiences = "contacts"
	SkillProficiencyAllowedAudiences_Everyone               SkillProficiencyAllowedAudiences = "everyone"
	SkillProficiencyAllowedAudiences_Family                 SkillProficiencyAllowedAudiences = "family"
	SkillProficiencyAllowedAudiences_FederatedOrganizations SkillProficiencyAllowedAudiences = "federatedOrganizations"
	SkillProficiencyAllowedAudiences_GroupMembers           SkillProficiencyAllowedAudiences = "groupMembers"
	SkillProficiencyAllowedAudiences_Me                     SkillProficiencyAllowedAudiences = "me"
	SkillProficiencyAllowedAudiences_Organization           SkillProficiencyAllowedAudiences = "organization"
)

func PossibleValuesForSkillProficiencyAllowedAudiences() []string {
	return []string{
		string(SkillProficiencyAllowedAudiences_Contacts),
		string(SkillProficiencyAllowedAudiences_Everyone),
		string(SkillProficiencyAllowedAudiences_Family),
		string(SkillProficiencyAllowedAudiences_FederatedOrganizations),
		string(SkillProficiencyAllowedAudiences_GroupMembers),
		string(SkillProficiencyAllowedAudiences_Me),
		string(SkillProficiencyAllowedAudiences_Organization),
	}
}

func (s *SkillProficiencyAllowedAudiences) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSkillProficiencyAllowedAudiences(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSkillProficiencyAllowedAudiences(input string) (*SkillProficiencyAllowedAudiences, error) {
	vals := map[string]SkillProficiencyAllowedAudiences{
		"contacts":               SkillProficiencyAllowedAudiences_Contacts,
		"everyone":               SkillProficiencyAllowedAudiences_Everyone,
		"family":                 SkillProficiencyAllowedAudiences_Family,
		"federatedorganizations": SkillProficiencyAllowedAudiences_FederatedOrganizations,
		"groupmembers":           SkillProficiencyAllowedAudiences_GroupMembers,
		"me":                     SkillProficiencyAllowedAudiences_Me,
		"organization":           SkillProficiencyAllowedAudiences_Organization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SkillProficiencyAllowedAudiences(input)
	return &out, nil
}
