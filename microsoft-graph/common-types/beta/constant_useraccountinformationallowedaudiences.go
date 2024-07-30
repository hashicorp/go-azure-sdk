package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserAccountInformationAllowedAudiences string

const (
	UserAccountInformationAllowedAudiences_Contacts               UserAccountInformationAllowedAudiences = "contacts"
	UserAccountInformationAllowedAudiences_Everyone               UserAccountInformationAllowedAudiences = "everyone"
	UserAccountInformationAllowedAudiences_Family                 UserAccountInformationAllowedAudiences = "family"
	UserAccountInformationAllowedAudiences_FederatedOrganizations UserAccountInformationAllowedAudiences = "federatedOrganizations"
	UserAccountInformationAllowedAudiences_GroupMembers           UserAccountInformationAllowedAudiences = "groupMembers"
	UserAccountInformationAllowedAudiences_Me                     UserAccountInformationAllowedAudiences = "me"
	UserAccountInformationAllowedAudiences_Organization           UserAccountInformationAllowedAudiences = "organization"
)

func PossibleValuesForUserAccountInformationAllowedAudiences() []string {
	return []string{
		string(UserAccountInformationAllowedAudiences_Contacts),
		string(UserAccountInformationAllowedAudiences_Everyone),
		string(UserAccountInformationAllowedAudiences_Family),
		string(UserAccountInformationAllowedAudiences_FederatedOrganizations),
		string(UserAccountInformationAllowedAudiences_GroupMembers),
		string(UserAccountInformationAllowedAudiences_Me),
		string(UserAccountInformationAllowedAudiences_Organization),
	}
}

func (s *UserAccountInformationAllowedAudiences) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserAccountInformationAllowedAudiences(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserAccountInformationAllowedAudiences(input string) (*UserAccountInformationAllowedAudiences, error) {
	vals := map[string]UserAccountInformationAllowedAudiences{
		"contacts":               UserAccountInformationAllowedAudiences_Contacts,
		"everyone":               UserAccountInformationAllowedAudiences_Everyone,
		"family":                 UserAccountInformationAllowedAudiences_Family,
		"federatedorganizations": UserAccountInformationAllowedAudiences_FederatedOrganizations,
		"groupmembers":           UserAccountInformationAllowedAudiences_GroupMembers,
		"me":                     UserAccountInformationAllowedAudiences_Me,
		"organization":           UserAccountInformationAllowedAudiences_Organization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserAccountInformationAllowedAudiences(input)
	return &out, nil
}
