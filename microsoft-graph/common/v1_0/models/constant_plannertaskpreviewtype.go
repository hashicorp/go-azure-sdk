package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerTaskPreviewType string

const (
	PlannerTaskPreviewTypeautomatic   PlannerTaskPreviewType = "Automatic"
	PlannerTaskPreviewTypechecklist   PlannerTaskPreviewType = "Checklist"
	PlannerTaskPreviewTypedescription PlannerTaskPreviewType = "Description"
	PlannerTaskPreviewTypenoPreview   PlannerTaskPreviewType = "NoPreview"
	PlannerTaskPreviewTypereference   PlannerTaskPreviewType = "Reference"
)

func PossibleValuesForPlannerTaskPreviewType() []string {
	return []string{
		string(PlannerTaskPreviewTypeautomatic),
		string(PlannerTaskPreviewTypechecklist),
		string(PlannerTaskPreviewTypedescription),
		string(PlannerTaskPreviewTypenoPreview),
		string(PlannerTaskPreviewTypereference),
	}
}

func parsePlannerTaskPreviewType(input string) (*PlannerTaskPreviewType, error) {
	vals := map[string]PlannerTaskPreviewType{
		"automatic":   PlannerTaskPreviewTypeautomatic,
		"checklist":   PlannerTaskPreviewTypechecklist,
		"description": PlannerTaskPreviewTypedescription,
		"nopreview":   PlannerTaskPreviewTypenoPreview,
		"reference":   PlannerTaskPreviewTypereference,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerTaskPreviewType(input)
	return &out, nil
}
