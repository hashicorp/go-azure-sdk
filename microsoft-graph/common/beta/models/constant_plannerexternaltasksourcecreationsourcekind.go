package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerExternalTaskSourceCreationSourceKind string

const (
	PlannerExternalTaskSourceCreationSourceKindexternal    PlannerExternalTaskSourceCreationSourceKind = "External"
	PlannerExternalTaskSourceCreationSourceKindnone        PlannerExternalTaskSourceCreationSourceKind = "None"
	PlannerExternalTaskSourceCreationSourceKindpublication PlannerExternalTaskSourceCreationSourceKind = "Publication"
)

func PossibleValuesForPlannerExternalTaskSourceCreationSourceKind() []string {
	return []string{
		string(PlannerExternalTaskSourceCreationSourceKindexternal),
		string(PlannerExternalTaskSourceCreationSourceKindnone),
		string(PlannerExternalTaskSourceCreationSourceKindpublication),
	}
}

func parsePlannerExternalTaskSourceCreationSourceKind(input string) (*PlannerExternalTaskSourceCreationSourceKind, error) {
	vals := map[string]PlannerExternalTaskSourceCreationSourceKind{
		"external":    PlannerExternalTaskSourceCreationSourceKindexternal,
		"none":        PlannerExternalTaskSourceCreationSourceKindnone,
		"publication": PlannerExternalTaskSourceCreationSourceKindpublication,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerExternalTaskSourceCreationSourceKind(input)
	return &out, nil
}
