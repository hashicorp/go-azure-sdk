package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthorizationPolicyAllowInvitesFrom string

const (
	AuthorizationPolicyAllowInvitesFrom_AdminsAndGuestInviters           AuthorizationPolicyAllowInvitesFrom = "adminsAndGuestInviters"
	AuthorizationPolicyAllowInvitesFrom_AdminsGuestInvitersAndAllMembers AuthorizationPolicyAllowInvitesFrom = "adminsGuestInvitersAndAllMembers"
	AuthorizationPolicyAllowInvitesFrom_Everyone                         AuthorizationPolicyAllowInvitesFrom = "everyone"
	AuthorizationPolicyAllowInvitesFrom_None                             AuthorizationPolicyAllowInvitesFrom = "none"
)

func PossibleValuesForAuthorizationPolicyAllowInvitesFrom() []string {
	return []string{
		string(AuthorizationPolicyAllowInvitesFrom_AdminsAndGuestInviters),
		string(AuthorizationPolicyAllowInvitesFrom_AdminsGuestInvitersAndAllMembers),
		string(AuthorizationPolicyAllowInvitesFrom_Everyone),
		string(AuthorizationPolicyAllowInvitesFrom_None),
	}
}

func (s *AuthorizationPolicyAllowInvitesFrom) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAuthorizationPolicyAllowInvitesFrom(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAuthorizationPolicyAllowInvitesFrom(input string) (*AuthorizationPolicyAllowInvitesFrom, error) {
	vals := map[string]AuthorizationPolicyAllowInvitesFrom{
		"adminsandguestinviters":           AuthorizationPolicyAllowInvitesFrom_AdminsAndGuestInviters,
		"adminsguestinvitersandallmembers": AuthorizationPolicyAllowInvitesFrom_AdminsGuestInvitersAndAllMembers,
		"everyone":                         AuthorizationPolicyAllowInvitesFrom_Everyone,
		"none":                             AuthorizationPolicyAllowInvitesFrom_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthorizationPolicyAllowInvitesFrom(input)
	return &out, nil
}
