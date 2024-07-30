package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebAccountAllowedAudiences string

const (
	WebAccountAllowedAudiences_Contacts               WebAccountAllowedAudiences = "contacts"
	WebAccountAllowedAudiences_Everyone               WebAccountAllowedAudiences = "everyone"
	WebAccountAllowedAudiences_Family                 WebAccountAllowedAudiences = "family"
	WebAccountAllowedAudiences_FederatedOrganizations WebAccountAllowedAudiences = "federatedOrganizations"
	WebAccountAllowedAudiences_GroupMembers           WebAccountAllowedAudiences = "groupMembers"
	WebAccountAllowedAudiences_Me                     WebAccountAllowedAudiences = "me"
	WebAccountAllowedAudiences_Organization           WebAccountAllowedAudiences = "organization"
)

func PossibleValuesForWebAccountAllowedAudiences() []string {
	return []string{
		string(WebAccountAllowedAudiences_Contacts),
		string(WebAccountAllowedAudiences_Everyone),
		string(WebAccountAllowedAudiences_Family),
		string(WebAccountAllowedAudiences_FederatedOrganizations),
		string(WebAccountAllowedAudiences_GroupMembers),
		string(WebAccountAllowedAudiences_Me),
		string(WebAccountAllowedAudiences_Organization),
	}
}

func (s *WebAccountAllowedAudiences) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWebAccountAllowedAudiences(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWebAccountAllowedAudiences(input string) (*WebAccountAllowedAudiences, error) {
	vals := map[string]WebAccountAllowedAudiences{
		"contacts":               WebAccountAllowedAudiences_Contacts,
		"everyone":               WebAccountAllowedAudiences_Everyone,
		"family":                 WebAccountAllowedAudiences_Family,
		"federatedorganizations": WebAccountAllowedAudiences_FederatedOrganizations,
		"groupmembers":           WebAccountAllowedAudiences_GroupMembers,
		"me":                     WebAccountAllowedAudiences_Me,
		"organization":           WebAccountAllowedAudiences_Organization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WebAccountAllowedAudiences(input)
	return &out, nil
}
