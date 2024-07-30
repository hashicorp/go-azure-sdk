package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerBucketCreationCreationSourceKind string

const (
	PlannerBucketCreationCreationSourceKind_External    PlannerBucketCreationCreationSourceKind = "external"
	PlannerBucketCreationCreationSourceKind_None        PlannerBucketCreationCreationSourceKind = "none"
	PlannerBucketCreationCreationSourceKind_Publication PlannerBucketCreationCreationSourceKind = "publication"
)

func PossibleValuesForPlannerBucketCreationCreationSourceKind() []string {
	return []string{
		string(PlannerBucketCreationCreationSourceKind_External),
		string(PlannerBucketCreationCreationSourceKind_None),
		string(PlannerBucketCreationCreationSourceKind_Publication),
	}
}

func (s *PlannerBucketCreationCreationSourceKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePlannerBucketCreationCreationSourceKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePlannerBucketCreationCreationSourceKind(input string) (*PlannerBucketCreationCreationSourceKind, error) {
	vals := map[string]PlannerBucketCreationCreationSourceKind{
		"external":    PlannerBucketCreationCreationSourceKind_External,
		"none":        PlannerBucketCreationCreationSourceKind_None,
		"publication": PlannerBucketCreationCreationSourceKind_Publication,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerBucketCreationCreationSourceKind(input)
	return &out, nil
}
