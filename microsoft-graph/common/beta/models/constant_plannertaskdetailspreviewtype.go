package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerTaskDetailsPreviewType string

const (
	PlannerTaskDetailsPreviewTypeautomatic   PlannerTaskDetailsPreviewType = "Automatic"
	PlannerTaskDetailsPreviewTypechecklist   PlannerTaskDetailsPreviewType = "Checklist"
	PlannerTaskDetailsPreviewTypedescription PlannerTaskDetailsPreviewType = "Description"
	PlannerTaskDetailsPreviewTypenoPreview   PlannerTaskDetailsPreviewType = "NoPreview"
	PlannerTaskDetailsPreviewTypereference   PlannerTaskDetailsPreviewType = "Reference"
)

func PossibleValuesForPlannerTaskDetailsPreviewType() []string {
	return []string{
		string(PlannerTaskDetailsPreviewTypeautomatic),
		string(PlannerTaskDetailsPreviewTypechecklist),
		string(PlannerTaskDetailsPreviewTypedescription),
		string(PlannerTaskDetailsPreviewTypenoPreview),
		string(PlannerTaskDetailsPreviewTypereference),
	}
}

func parsePlannerTaskDetailsPreviewType(input string) (*PlannerTaskDetailsPreviewType, error) {
	vals := map[string]PlannerTaskDetailsPreviewType{
		"automatic":   PlannerTaskDetailsPreviewTypeautomatic,
		"checklist":   PlannerTaskDetailsPreviewTypechecklist,
		"description": PlannerTaskDetailsPreviewTypedescription,
		"nopreview":   PlannerTaskDetailsPreviewTypenoPreview,
		"reference":   PlannerTaskDetailsPreviewTypereference,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerTaskDetailsPreviewType(input)
	return &out, nil
}
