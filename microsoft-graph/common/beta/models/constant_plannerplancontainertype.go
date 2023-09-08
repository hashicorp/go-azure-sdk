package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerPlanContainerType string

const (
	PlannerPlanContainerTypedriveItem PlannerPlanContainerType = "DriveItem"
	PlannerPlanContainerTypegroup     PlannerPlanContainerType = "Group"
	PlannerPlanContainerTypeproject   PlannerPlanContainerType = "Project"
	PlannerPlanContainerTyperoster    PlannerPlanContainerType = "Roster"
)

func PossibleValuesForPlannerPlanContainerType() []string {
	return []string{
		string(PlannerPlanContainerTypedriveItem),
		string(PlannerPlanContainerTypegroup),
		string(PlannerPlanContainerTypeproject),
		string(PlannerPlanContainerTyperoster),
	}
}

func parsePlannerPlanContainerType(input string) (*PlannerPlanContainerType, error) {
	vals := map[string]PlannerPlanContainerType{
		"driveitem": PlannerPlanContainerTypedriveItem,
		"group":     PlannerPlanContainerTypegroup,
		"project":   PlannerPlanContainerTypeproject,
		"roster":    PlannerPlanContainerTyperoster,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerPlanContainerType(input)
	return &out, nil
}
