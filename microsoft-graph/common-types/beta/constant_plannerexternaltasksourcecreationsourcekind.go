package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerExternalTaskSourceCreationSourceKind string

const (
	PlannerExternalTaskSourceCreationSourceKind_External    PlannerExternalTaskSourceCreationSourceKind = "external"
	PlannerExternalTaskSourceCreationSourceKind_None        PlannerExternalTaskSourceCreationSourceKind = "none"
	PlannerExternalTaskSourceCreationSourceKind_Publication PlannerExternalTaskSourceCreationSourceKind = "publication"
)

func PossibleValuesForPlannerExternalTaskSourceCreationSourceKind() []string {
	return []string{
		string(PlannerExternalTaskSourceCreationSourceKind_External),
		string(PlannerExternalTaskSourceCreationSourceKind_None),
		string(PlannerExternalTaskSourceCreationSourceKind_Publication),
	}
}

func (s *PlannerExternalTaskSourceCreationSourceKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePlannerExternalTaskSourceCreationSourceKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePlannerExternalTaskSourceCreationSourceKind(input string) (*PlannerExternalTaskSourceCreationSourceKind, error) {
	vals := map[string]PlannerExternalTaskSourceCreationSourceKind{
		"external":    PlannerExternalTaskSourceCreationSourceKind_External,
		"none":        PlannerExternalTaskSourceCreationSourceKind_None,
		"publication": PlannerExternalTaskSourceCreationSourceKind_Publication,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerExternalTaskSourceCreationSourceKind(input)
	return &out, nil
}
