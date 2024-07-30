package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerTaskPreviewType string

const (
	PlannerTaskPreviewType_Automatic   PlannerTaskPreviewType = "automatic"
	PlannerTaskPreviewType_Checklist   PlannerTaskPreviewType = "checklist"
	PlannerTaskPreviewType_Description PlannerTaskPreviewType = "description"
	PlannerTaskPreviewType_NoPreview   PlannerTaskPreviewType = "noPreview"
	PlannerTaskPreviewType_Reference   PlannerTaskPreviewType = "reference"
)

func PossibleValuesForPlannerTaskPreviewType() []string {
	return []string{
		string(PlannerTaskPreviewType_Automatic),
		string(PlannerTaskPreviewType_Checklist),
		string(PlannerTaskPreviewType_Description),
		string(PlannerTaskPreviewType_NoPreview),
		string(PlannerTaskPreviewType_Reference),
	}
}

func (s *PlannerTaskPreviewType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePlannerTaskPreviewType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePlannerTaskPreviewType(input string) (*PlannerTaskPreviewType, error) {
	vals := map[string]PlannerTaskPreviewType{
		"automatic":   PlannerTaskPreviewType_Automatic,
		"checklist":   PlannerTaskPreviewType_Checklist,
		"description": PlannerTaskPreviewType_Description,
		"nopreview":   PlannerTaskPreviewType_NoPreview,
		"reference":   PlannerTaskPreviewType_Reference,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerTaskPreviewType(input)
	return &out, nil
}
