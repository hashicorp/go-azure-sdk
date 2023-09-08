package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DecisionItemPrincipalResourceMembershipMembershipType string

const (
	DecisionItemPrincipalResourceMembershipMembershipTypedirect   DecisionItemPrincipalResourceMembershipMembershipType = "Direct"
	DecisionItemPrincipalResourceMembershipMembershipTypeindirect DecisionItemPrincipalResourceMembershipMembershipType = "Indirect"
)

func PossibleValuesForDecisionItemPrincipalResourceMembershipMembershipType() []string {
	return []string{
		string(DecisionItemPrincipalResourceMembershipMembershipTypedirect),
		string(DecisionItemPrincipalResourceMembershipMembershipTypeindirect),
	}
}

func parseDecisionItemPrincipalResourceMembershipMembershipType(input string) (*DecisionItemPrincipalResourceMembershipMembershipType, error) {
	vals := map[string]DecisionItemPrincipalResourceMembershipMembershipType{
		"direct":   DecisionItemPrincipalResourceMembershipMembershipTypedirect,
		"indirect": DecisionItemPrincipalResourceMembershipMembershipTypeindirect,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DecisionItemPrincipalResourceMembershipMembershipType(input)
	return &out, nil
}
