package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerTaskCreationCreationSourceKind string

const (
	PlannerTaskCreationCreationSourceKind_External    PlannerTaskCreationCreationSourceKind = "external"
	PlannerTaskCreationCreationSourceKind_None        PlannerTaskCreationCreationSourceKind = "none"
	PlannerTaskCreationCreationSourceKind_Publication PlannerTaskCreationCreationSourceKind = "publication"
)

func PossibleValuesForPlannerTaskCreationCreationSourceKind() []string {
	return []string{
		string(PlannerTaskCreationCreationSourceKind_External),
		string(PlannerTaskCreationCreationSourceKind_None),
		string(PlannerTaskCreationCreationSourceKind_Publication),
	}
}

func (s *PlannerTaskCreationCreationSourceKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePlannerTaskCreationCreationSourceKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePlannerTaskCreationCreationSourceKind(input string) (*PlannerTaskCreationCreationSourceKind, error) {
	vals := map[string]PlannerTaskCreationCreationSourceKind{
		"external":    PlannerTaskCreationCreationSourceKind_External,
		"none":        PlannerTaskCreationCreationSourceKind_None,
		"publication": PlannerTaskCreationCreationSourceKind_Publication,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerTaskCreationCreationSourceKind(input)
	return &out, nil
}
