package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerPlanCreationCreationSourceKind string

const (
	PlannerPlanCreationCreationSourceKind_External    PlannerPlanCreationCreationSourceKind = "external"
	PlannerPlanCreationCreationSourceKind_None        PlannerPlanCreationCreationSourceKind = "none"
	PlannerPlanCreationCreationSourceKind_Publication PlannerPlanCreationCreationSourceKind = "publication"
)

func PossibleValuesForPlannerPlanCreationCreationSourceKind() []string {
	return []string{
		string(PlannerPlanCreationCreationSourceKind_External),
		string(PlannerPlanCreationCreationSourceKind_None),
		string(PlannerPlanCreationCreationSourceKind_Publication),
	}
}

func (s *PlannerPlanCreationCreationSourceKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePlannerPlanCreationCreationSourceKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePlannerPlanCreationCreationSourceKind(input string) (*PlannerPlanCreationCreationSourceKind, error) {
	vals := map[string]PlannerPlanCreationCreationSourceKind{
		"external":    PlannerPlanCreationCreationSourceKind_External,
		"none":        PlannerPlanCreationCreationSourceKind_None,
		"publication": PlannerPlanCreationCreationSourceKind_Publication,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerPlanCreationCreationSourceKind(input)
	return &out, nil
}
