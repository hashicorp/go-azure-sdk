package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerTaskCreationCreationSourceKind string

const (
	PlannerTaskCreationCreationSourceKindexternal    PlannerTaskCreationCreationSourceKind = "External"
	PlannerTaskCreationCreationSourceKindnone        PlannerTaskCreationCreationSourceKind = "None"
	PlannerTaskCreationCreationSourceKindpublication PlannerTaskCreationCreationSourceKind = "Publication"
)

func PossibleValuesForPlannerTaskCreationCreationSourceKind() []string {
	return []string{
		string(PlannerTaskCreationCreationSourceKindexternal),
		string(PlannerTaskCreationCreationSourceKindnone),
		string(PlannerTaskCreationCreationSourceKindpublication),
	}
}

func parsePlannerTaskCreationCreationSourceKind(input string) (*PlannerTaskCreationCreationSourceKind, error) {
	vals := map[string]PlannerTaskCreationCreationSourceKind{
		"external":    PlannerTaskCreationCreationSourceKindexternal,
		"none":        PlannerTaskCreationCreationSourceKindnone,
		"publication": PlannerTaskCreationCreationSourceKindpublication,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerTaskCreationCreationSourceKind(input)
	return &out, nil
}
