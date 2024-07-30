package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BusinessScenarioTaskTargetBaseTaskTargetKind string

const (
	BusinessScenarioTaskTargetBaseTaskTargetKind_Group BusinessScenarioTaskTargetBaseTaskTargetKind = "group"
)

func PossibleValuesForBusinessScenarioTaskTargetBaseTaskTargetKind() []string {
	return []string{
		string(BusinessScenarioTaskTargetBaseTaskTargetKind_Group),
	}
}

func (s *BusinessScenarioTaskTargetBaseTaskTargetKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBusinessScenarioTaskTargetBaseTaskTargetKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBusinessScenarioTaskTargetBaseTaskTargetKind(input string) (*BusinessScenarioTaskTargetBaseTaskTargetKind, error) {
	vals := map[string]BusinessScenarioTaskTargetBaseTaskTargetKind{
		"group": BusinessScenarioTaskTargetBaseTaskTargetKind_Group,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BusinessScenarioTaskTargetBaseTaskTargetKind(input)
	return &out, nil
}
