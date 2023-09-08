package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerPlanContainerType string

const (
	PlannerPlanContainerTypegroup  PlannerPlanContainerType = "Group"
	PlannerPlanContainerTyperoster PlannerPlanContainerType = "Roster"
)

func PossibleValuesForPlannerPlanContainerType() []string {
	return []string{
		string(PlannerPlanContainerTypegroup),
		string(PlannerPlanContainerTyperoster),
	}
}

func parsePlannerPlanContainerType(input string) (*PlannerPlanContainerType, error) {
	vals := map[string]PlannerPlanContainerType{
		"group":  PlannerPlanContainerTypegroup,
		"roster": PlannerPlanContainerTyperoster,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerPlanContainerType(input)
	return &out, nil
}
