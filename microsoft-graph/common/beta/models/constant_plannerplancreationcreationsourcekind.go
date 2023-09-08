package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerPlanCreationCreationSourceKind string

const (
	PlannerPlanCreationCreationSourceKindexternal    PlannerPlanCreationCreationSourceKind = "External"
	PlannerPlanCreationCreationSourceKindnone        PlannerPlanCreationCreationSourceKind = "None"
	PlannerPlanCreationCreationSourceKindpublication PlannerPlanCreationCreationSourceKind = "Publication"
)

func PossibleValuesForPlannerPlanCreationCreationSourceKind() []string {
	return []string{
		string(PlannerPlanCreationCreationSourceKindexternal),
		string(PlannerPlanCreationCreationSourceKindnone),
		string(PlannerPlanCreationCreationSourceKindpublication),
	}
}

func parsePlannerPlanCreationCreationSourceKind(input string) (*PlannerPlanCreationCreationSourceKind, error) {
	vals := map[string]PlannerPlanCreationCreationSourceKind{
		"external":    PlannerPlanCreationCreationSourceKindexternal,
		"none":        PlannerPlanCreationCreationSourceKindnone,
		"publication": PlannerPlanCreationCreationSourceKindpublication,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerPlanCreationCreationSourceKind(input)
	return &out, nil
}
