package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerPlanContainerType string

const (
	PlannerPlanContainerType_DriveItem PlannerPlanContainerType = "driveItem"
	PlannerPlanContainerType_Group     PlannerPlanContainerType = "group"
	PlannerPlanContainerType_Project   PlannerPlanContainerType = "project"
	PlannerPlanContainerType_Roster    PlannerPlanContainerType = "roster"
	PlannerPlanContainerType_User      PlannerPlanContainerType = "user"
)

func PossibleValuesForPlannerPlanContainerType() []string {
	return []string{
		string(PlannerPlanContainerType_DriveItem),
		string(PlannerPlanContainerType_Group),
		string(PlannerPlanContainerType_Project),
		string(PlannerPlanContainerType_Roster),
		string(PlannerPlanContainerType_User),
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
		"driveitem": PlannerPlanContainerType_DriveItem,
		"group":     PlannerPlanContainerType_Group,
		"project":   PlannerPlanContainerType_Project,
		"roster":    PlannerPlanContainerType_Roster,
		"user":      PlannerPlanContainerType_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerPlanContainerType(input)
	return &out, nil
}
