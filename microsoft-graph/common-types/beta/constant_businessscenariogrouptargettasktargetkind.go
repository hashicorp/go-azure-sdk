package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BusinessScenarioGroupTargetTaskTargetKind string

const (
	BusinessScenarioGroupTargetTaskTargetKind_Group BusinessScenarioGroupTargetTaskTargetKind = "group"
)

func PossibleValuesForBusinessScenarioGroupTargetTaskTargetKind() []string {
	return []string{
		string(BusinessScenarioGroupTargetTaskTargetKind_Group),
	}
}

func (s *BusinessScenarioGroupTargetTaskTargetKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBusinessScenarioGroupTargetTaskTargetKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBusinessScenarioGroupTargetTaskTargetKind(input string) (*BusinessScenarioGroupTargetTaskTargetKind, error) {
	vals := map[string]BusinessScenarioGroupTargetTaskTargetKind{
		"group": BusinessScenarioGroupTargetTaskTargetKind_Group,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BusinessScenarioGroupTargetTaskTargetKind(input)
	return &out, nil
}
