package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthorizationPolicyAllowInvitesFrom string

const (
	AuthorizationPolicyAllowInvitesFromadminsAndGuestInviters           AuthorizationPolicyAllowInvitesFrom = "AdminsAndGuestInviters"
	AuthorizationPolicyAllowInvitesFromadminsGuestInvitersAndAllMembers AuthorizationPolicyAllowInvitesFrom = "AdminsGuestInvitersAndAllMembers"
	AuthorizationPolicyAllowInvitesFromeveryone                         AuthorizationPolicyAllowInvitesFrom = "Everyone"
	AuthorizationPolicyAllowInvitesFromnone                             AuthorizationPolicyAllowInvitesFrom = "None"
)

func PossibleValuesForAuthorizationPolicyAllowInvitesFrom() []string {
	return []string{
		string(AuthorizationPolicyAllowInvitesFromadminsAndGuestInviters),
		string(AuthorizationPolicyAllowInvitesFromadminsGuestInvitersAndAllMembers),
		string(AuthorizationPolicyAllowInvitesFromeveryone),
		string(AuthorizationPolicyAllowInvitesFromnone),
	}
}

func parseAuthorizationPolicyAllowInvitesFrom(input string) (*AuthorizationPolicyAllowInvitesFrom, error) {
	vals := map[string]AuthorizationPolicyAllowInvitesFrom{
		"adminsandguestinviters":           AuthorizationPolicyAllowInvitesFromadminsAndGuestInviters,
		"adminsguestinvitersandallmembers": AuthorizationPolicyAllowInvitesFromadminsGuestInvitersAndAllMembers,
		"everyone":                         AuthorizationPolicyAllowInvitesFromeveryone,
		"none":                             AuthorizationPolicyAllowInvitesFromnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthorizationPolicyAllowInvitesFrom(input)
	return &out, nil
}
