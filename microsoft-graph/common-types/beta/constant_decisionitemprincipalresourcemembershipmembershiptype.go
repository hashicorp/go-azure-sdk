package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DecisionItemPrincipalResourceMembershipMembershipType string

const (
	DecisionItemPrincipalResourceMembershipMembershipType_Direct   DecisionItemPrincipalResourceMembershipMembershipType = "direct"
	DecisionItemPrincipalResourceMembershipMembershipType_Indirect DecisionItemPrincipalResourceMembershipMembershipType = "indirect"
)

func PossibleValuesForDecisionItemPrincipalResourceMembershipMembershipType() []string {
	return []string{
		string(DecisionItemPrincipalResourceMembershipMembershipType_Direct),
		string(DecisionItemPrincipalResourceMembershipMembershipType_Indirect),
	}
}

func (s *DecisionItemPrincipalResourceMembershipMembershipType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDecisionItemPrincipalResourceMembershipMembershipType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDecisionItemPrincipalResourceMembershipMembershipType(input string) (*DecisionItemPrincipalResourceMembershipMembershipType, error) {
	vals := map[string]DecisionItemPrincipalResourceMembershipMembershipType{
		"direct":   DecisionItemPrincipalResourceMembershipMembershipType_Direct,
		"indirect": DecisionItemPrincipalResourceMembershipMembershipType_Indirect,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DecisionItemPrincipalResourceMembershipMembershipType(input)
	return &out, nil
}
