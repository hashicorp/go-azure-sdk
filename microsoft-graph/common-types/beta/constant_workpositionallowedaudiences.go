package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkPositionAllowedAudiences string

const (
	WorkPositionAllowedAudiences_Contacts               WorkPositionAllowedAudiences = "contacts"
	WorkPositionAllowedAudiences_Everyone               WorkPositionAllowedAudiences = "everyone"
	WorkPositionAllowedAudiences_Family                 WorkPositionAllowedAudiences = "family"
	WorkPositionAllowedAudiences_FederatedOrganizations WorkPositionAllowedAudiences = "federatedOrganizations"
	WorkPositionAllowedAudiences_GroupMembers           WorkPositionAllowedAudiences = "groupMembers"
	WorkPositionAllowedAudiences_Me                     WorkPositionAllowedAudiences = "me"
	WorkPositionAllowedAudiences_Organization           WorkPositionAllowedAudiences = "organization"
)

func PossibleValuesForWorkPositionAllowedAudiences() []string {
	return []string{
		string(WorkPositionAllowedAudiences_Contacts),
		string(WorkPositionAllowedAudiences_Everyone),
		string(WorkPositionAllowedAudiences_Family),
		string(WorkPositionAllowedAudiences_FederatedOrganizations),
		string(WorkPositionAllowedAudiences_GroupMembers),
		string(WorkPositionAllowedAudiences_Me),
		string(WorkPositionAllowedAudiences_Organization),
	}
}

func (s *WorkPositionAllowedAudiences) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWorkPositionAllowedAudiences(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWorkPositionAllowedAudiences(input string) (*WorkPositionAllowedAudiences, error) {
	vals := map[string]WorkPositionAllowedAudiences{
		"contacts":               WorkPositionAllowedAudiences_Contacts,
		"everyone":               WorkPositionAllowedAudiences_Everyone,
		"family":                 WorkPositionAllowedAudiences_Family,
		"federatedorganizations": WorkPositionAllowedAudiences_FederatedOrganizations,
		"groupmembers":           WorkPositionAllowedAudiences_GroupMembers,
		"me":                     WorkPositionAllowedAudiences_Me,
		"organization":           WorkPositionAllowedAudiences_Organization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WorkPositionAllowedAudiences(input)
	return &out, nil
}
