package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerTaskDetailsPreviewType string

const (
	PlannerTaskDetailsPreviewType_Automatic   PlannerTaskDetailsPreviewType = "automatic"
	PlannerTaskDetailsPreviewType_Checklist   PlannerTaskDetailsPreviewType = "checklist"
	PlannerTaskDetailsPreviewType_Description PlannerTaskDetailsPreviewType = "description"
	PlannerTaskDetailsPreviewType_NoPreview   PlannerTaskDetailsPreviewType = "noPreview"
	PlannerTaskDetailsPreviewType_Reference   PlannerTaskDetailsPreviewType = "reference"
)

func PossibleValuesForPlannerTaskDetailsPreviewType() []string {
	return []string{
		string(PlannerTaskDetailsPreviewType_Automatic),
		string(PlannerTaskDetailsPreviewType_Checklist),
		string(PlannerTaskDetailsPreviewType_Description),
		string(PlannerTaskDetailsPreviewType_NoPreview),
		string(PlannerTaskDetailsPreviewType_Reference),
	}
}

func (s *PlannerTaskDetailsPreviewType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePlannerTaskDetailsPreviewType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePlannerTaskDetailsPreviewType(input string) (*PlannerTaskDetailsPreviewType, error) {
	vals := map[string]PlannerTaskDetailsPreviewType{
		"automatic":   PlannerTaskDetailsPreviewType_Automatic,
		"checklist":   PlannerTaskDetailsPreviewType_Checklist,
		"description": PlannerTaskDetailsPreviewType_Description,
		"nopreview":   PlannerTaskDetailsPreviewType_NoPreview,
		"reference":   PlannerTaskDetailsPreviewType_Reference,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerTaskDetailsPreviewType(input)
	return &out, nil
}
