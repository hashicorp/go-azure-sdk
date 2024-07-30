package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerExternalTaskSourceDisplayLinkType string

const (
	PlannerExternalTaskSourceDisplayLinkType_Default PlannerExternalTaskSourceDisplayLinkType = "default"
	PlannerExternalTaskSourceDisplayLinkType_None    PlannerExternalTaskSourceDisplayLinkType = "none"
)

func PossibleValuesForPlannerExternalTaskSourceDisplayLinkType() []string {
	return []string{
		string(PlannerExternalTaskSourceDisplayLinkType_Default),
		string(PlannerExternalTaskSourceDisplayLinkType_None),
	}
}

func (s *PlannerExternalTaskSourceDisplayLinkType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePlannerExternalTaskSourceDisplayLinkType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePlannerExternalTaskSourceDisplayLinkType(input string) (*PlannerExternalTaskSourceDisplayLinkType, error) {
	vals := map[string]PlannerExternalTaskSourceDisplayLinkType{
		"default": PlannerExternalTaskSourceDisplayLinkType_Default,
		"none":    PlannerExternalTaskSourceDisplayLinkType_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerExternalTaskSourceDisplayLinkType(input)
	return &out, nil
}
