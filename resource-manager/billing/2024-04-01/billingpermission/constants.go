package billingpermission

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessDecision string

const (
	AccessDecisionAllowed    AccessDecision = "Allowed"
	AccessDecisionNotAllowed AccessDecision = "NotAllowed"
	AccessDecisionOther      AccessDecision = "Other"
)

func PossibleValuesForAccessDecision() []string {
	return []string{
		string(AccessDecisionAllowed),
		string(AccessDecisionNotAllowed),
		string(AccessDecisionOther),
	}
}

func (s *AccessDecision) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAccessDecision(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAccessDecision(input string) (*AccessDecision, error) {
	vals := map[string]AccessDecision{
		"allowed":    AccessDecisionAllowed,
		"notallowed": AccessDecisionNotAllowed,
		"other":      AccessDecisionOther,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessDecision(input)
	return &out, nil
}
