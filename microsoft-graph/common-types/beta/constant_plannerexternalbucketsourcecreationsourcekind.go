package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerExternalBucketSourceCreationSourceKind string

const (
	PlannerExternalBucketSourceCreationSourceKind_External    PlannerExternalBucketSourceCreationSourceKind = "external"
	PlannerExternalBucketSourceCreationSourceKind_None        PlannerExternalBucketSourceCreationSourceKind = "none"
	PlannerExternalBucketSourceCreationSourceKind_Publication PlannerExternalBucketSourceCreationSourceKind = "publication"
)

func PossibleValuesForPlannerExternalBucketSourceCreationSourceKind() []string {
	return []string{
		string(PlannerExternalBucketSourceCreationSourceKind_External),
		string(PlannerExternalBucketSourceCreationSourceKind_None),
		string(PlannerExternalBucketSourceCreationSourceKind_Publication),
	}
}

func (s *PlannerExternalBucketSourceCreationSourceKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePlannerExternalBucketSourceCreationSourceKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePlannerExternalBucketSourceCreationSourceKind(input string) (*PlannerExternalBucketSourceCreationSourceKind, error) {
	vals := map[string]PlannerExternalBucketSourceCreationSourceKind{
		"external":    PlannerExternalBucketSourceCreationSourceKind_External,
		"none":        PlannerExternalBucketSourceCreationSourceKind_None,
		"publication": PlannerExternalBucketSourceCreationSourceKind_Publication,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerExternalBucketSourceCreationSourceKind(input)
	return &out, nil
}
