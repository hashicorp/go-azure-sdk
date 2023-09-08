package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerSharedWithContainerAccessLevel string

const (
	PlannerSharedWithContainerAccessLevelfullAccess      PlannerSharedWithContainerAccessLevel = "FullAccess"
	PlannerSharedWithContainerAccessLevelreadAccess      PlannerSharedWithContainerAccessLevel = "ReadAccess"
	PlannerSharedWithContainerAccessLevelreadWriteAccess PlannerSharedWithContainerAccessLevel = "ReadWriteAccess"
)

func PossibleValuesForPlannerSharedWithContainerAccessLevel() []string {
	return []string{
		string(PlannerSharedWithContainerAccessLevelfullAccess),
		string(PlannerSharedWithContainerAccessLevelreadAccess),
		string(PlannerSharedWithContainerAccessLevelreadWriteAccess),
	}
}

func parsePlannerSharedWithContainerAccessLevel(input string) (*PlannerSharedWithContainerAccessLevel, error) {
	vals := map[string]PlannerSharedWithContainerAccessLevel{
		"fullaccess":      PlannerSharedWithContainerAccessLevelfullAccess,
		"readaccess":      PlannerSharedWithContainerAccessLevelreadAccess,
		"readwriteaccess": PlannerSharedWithContainerAccessLevelreadWriteAccess,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerSharedWithContainerAccessLevel(input)
	return &out, nil
}
