package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerSharedWithContainerType string

const (
	PlannerSharedWithContainerTypedriveItem PlannerSharedWithContainerType = "DriveItem"
	PlannerSharedWithContainerTypegroup     PlannerSharedWithContainerType = "Group"
	PlannerSharedWithContainerTypeproject   PlannerSharedWithContainerType = "Project"
	PlannerSharedWithContainerTyperoster    PlannerSharedWithContainerType = "Roster"
)

func PossibleValuesForPlannerSharedWithContainerType() []string {
	return []string{
		string(PlannerSharedWithContainerTypedriveItem),
		string(PlannerSharedWithContainerTypegroup),
		string(PlannerSharedWithContainerTypeproject),
		string(PlannerSharedWithContainerTyperoster),
	}
}

func parsePlannerSharedWithContainerType(input string) (*PlannerSharedWithContainerType, error) {
	vals := map[string]PlannerSharedWithContainerType{
		"driveitem": PlannerSharedWithContainerTypedriveItem,
		"group":     PlannerSharedWithContainerTypegroup,
		"project":   PlannerSharedWithContainerTypeproject,
		"roster":    PlannerSharedWithContainerTyperoster,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerSharedWithContainerType(input)
	return &out, nil
}
