package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProjectParticipationAllowedAudiences string

const (
	ProjectParticipationAllowedAudiences_Contacts               ProjectParticipationAllowedAudiences = "contacts"
	ProjectParticipationAllowedAudiences_Everyone               ProjectParticipationAllowedAudiences = "everyone"
	ProjectParticipationAllowedAudiences_Family                 ProjectParticipationAllowedAudiences = "family"
	ProjectParticipationAllowedAudiences_FederatedOrganizations ProjectParticipationAllowedAudiences = "federatedOrganizations"
	ProjectParticipationAllowedAudiences_GroupMembers           ProjectParticipationAllowedAudiences = "groupMembers"
	ProjectParticipationAllowedAudiences_Me                     ProjectParticipationAllowedAudiences = "me"
	ProjectParticipationAllowedAudiences_Organization           ProjectParticipationAllowedAudiences = "organization"
)

func PossibleValuesForProjectParticipationAllowedAudiences() []string {
	return []string{
		string(ProjectParticipationAllowedAudiences_Contacts),
		string(ProjectParticipationAllowedAudiences_Everyone),
		string(ProjectParticipationAllowedAudiences_Family),
		string(ProjectParticipationAllowedAudiences_FederatedOrganizations),
		string(ProjectParticipationAllowedAudiences_GroupMembers),
		string(ProjectParticipationAllowedAudiences_Me),
		string(ProjectParticipationAllowedAudiences_Organization),
	}
}

func (s *ProjectParticipationAllowedAudiences) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProjectParticipationAllowedAudiences(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProjectParticipationAllowedAudiences(input string) (*ProjectParticipationAllowedAudiences, error) {
	vals := map[string]ProjectParticipationAllowedAudiences{
		"contacts":               ProjectParticipationAllowedAudiences_Contacts,
		"everyone":               ProjectParticipationAllowedAudiences_Everyone,
		"family":                 ProjectParticipationAllowedAudiences_Family,
		"federatedorganizations": ProjectParticipationAllowedAudiences_FederatedOrganizations,
		"groupmembers":           ProjectParticipationAllowedAudiences_GroupMembers,
		"me":                     ProjectParticipationAllowedAudiences_Me,
		"organization":           ProjectParticipationAllowedAudiences_Organization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProjectParticipationAllowedAudiences(input)
	return &out, nil
}
