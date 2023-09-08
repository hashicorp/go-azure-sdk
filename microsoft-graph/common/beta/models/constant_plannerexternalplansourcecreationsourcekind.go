package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerExternalPlanSourceCreationSourceKind string

const (
	PlannerExternalPlanSourceCreationSourceKindexternal    PlannerExternalPlanSourceCreationSourceKind = "External"
	PlannerExternalPlanSourceCreationSourceKindnone        PlannerExternalPlanSourceCreationSourceKind = "None"
	PlannerExternalPlanSourceCreationSourceKindpublication PlannerExternalPlanSourceCreationSourceKind = "Publication"
)

func PossibleValuesForPlannerExternalPlanSourceCreationSourceKind() []string {
	return []string{
		string(PlannerExternalPlanSourceCreationSourceKindexternal),
		string(PlannerExternalPlanSourceCreationSourceKindnone),
		string(PlannerExternalPlanSourceCreationSourceKindpublication),
	}
}

func parsePlannerExternalPlanSourceCreationSourceKind(input string) (*PlannerExternalPlanSourceCreationSourceKind, error) {
	vals := map[string]PlannerExternalPlanSourceCreationSourceKind{
		"external":    PlannerExternalPlanSourceCreationSourceKindexternal,
		"none":        PlannerExternalPlanSourceCreationSourceKindnone,
		"publication": PlannerExternalPlanSourceCreationSourceKindpublication,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerExternalPlanSourceCreationSourceKind(input)
	return &out, nil
}
