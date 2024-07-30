package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerSharedWithContainerType string

const (
	PlannerSharedWithContainerType_DriveItem PlannerSharedWithContainerType = "driveItem"
	PlannerSharedWithContainerType_Group     PlannerSharedWithContainerType = "group"
	PlannerSharedWithContainerType_Project   PlannerSharedWithContainerType = "project"
	PlannerSharedWithContainerType_Roster    PlannerSharedWithContainerType = "roster"
	PlannerSharedWithContainerType_User      PlannerSharedWithContainerType = "user"
)

func PossibleValuesForPlannerSharedWithContainerType() []string {
	return []string{
		string(PlannerSharedWithContainerType_DriveItem),
		string(PlannerSharedWithContainerType_Group),
		string(PlannerSharedWithContainerType_Project),
		string(PlannerSharedWithContainerType_Roster),
		string(PlannerSharedWithContainerType_User),
	}
}

func (s *PlannerSharedWithContainerType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePlannerSharedWithContainerType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePlannerSharedWithContainerType(input string) (*PlannerSharedWithContainerType, error) {
	vals := map[string]PlannerSharedWithContainerType{
		"driveitem": PlannerSharedWithContainerType_DriveItem,
		"group":     PlannerSharedWithContainerType_Group,
		"project":   PlannerSharedWithContainerType_Project,
		"roster":    PlannerSharedWithContainerType_Roster,
		"user":      PlannerSharedWithContainerType_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerSharedWithContainerType(input)
	return &out, nil
}
