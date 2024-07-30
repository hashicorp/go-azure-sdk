package plannerplan

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerPlanContainerType string

const (
	PlannerPlanContainerTypeDriveItem PlannerPlanContainerType = "driveItem"
	PlannerPlanContainerTypeGroup     PlannerPlanContainerType = "group"
	PlannerPlanContainerTypeProject   PlannerPlanContainerType = "project"
	PlannerPlanContainerTypeRoster    PlannerPlanContainerType = "roster"
	PlannerPlanContainerTypeUser      PlannerPlanContainerType = "user"
)

func PossibleValuesForPlannerPlanContainerType() []string {
	return []string{
		string(PlannerPlanContainerTypeDriveItem),
		string(PlannerPlanContainerTypeGroup),
		string(PlannerPlanContainerTypeProject),
		string(PlannerPlanContainerTypeRoster),
		string(PlannerPlanContainerTypeUser),
	}
}

func (s *PlannerPlanContainerType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePlannerPlanContainerType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePlannerPlanContainerType(input string) (*PlannerPlanContainerType, error) {
	vals := map[string]PlannerPlanContainerType{
		"driveitem": PlannerPlanContainerTypeDriveItem,
		"group":     PlannerPlanContainerTypeGroup,
		"project":   PlannerPlanContainerTypeProject,
		"roster":    PlannerPlanContainerTypeRoster,
		"user":      PlannerPlanContainerTypeUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerPlanContainerType(input)
	return &out, nil
}
