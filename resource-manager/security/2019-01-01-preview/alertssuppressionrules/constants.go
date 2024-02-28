package alertssuppressionrules

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RuleState string

const (
	RuleStateDisabled RuleState = "Disabled"
	RuleStateEnabled  RuleState = "Enabled"
	RuleStateExpired  RuleState = "Expired"
)

func PossibleValuesForRuleState() []string {
	return []string{
		string(RuleStateDisabled),
		string(RuleStateEnabled),
		string(RuleStateExpired),
	}
}

func (s *RuleState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRuleState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRuleState(input string) (*RuleState, error) {
	vals := map[string]RuleState{
		"disabled": RuleStateDisabled,
		"enabled":  RuleStateEnabled,
		"expired":  RuleStateExpired,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RuleState(input)
	return &out, nil
}
