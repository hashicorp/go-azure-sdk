package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationAssignmentAddToCalendarAction string

const (
	EducationAssignmentAddToCalendarAction_None                  EducationAssignmentAddToCalendarAction = "none"
	EducationAssignmentAddToCalendarAction_StudentsAndPublisher  EducationAssignmentAddToCalendarAction = "studentsAndPublisher"
	EducationAssignmentAddToCalendarAction_StudentsAndTeamOwners EducationAssignmentAddToCalendarAction = "studentsAndTeamOwners"
	EducationAssignmentAddToCalendarAction_StudentsOnly          EducationAssignmentAddToCalendarAction = "studentsOnly"
)

func PossibleValuesForEducationAssignmentAddToCalendarAction() []string {
	return []string{
		string(EducationAssignmentAddToCalendarAction_None),
		string(EducationAssignmentAddToCalendarAction_StudentsAndPublisher),
		string(EducationAssignmentAddToCalendarAction_StudentsAndTeamOwners),
		string(EducationAssignmentAddToCalendarAction_StudentsOnly),
	}
}

func (s *EducationAssignmentAddToCalendarAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEducationAssignmentAddToCalendarAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEducationAssignmentAddToCalendarAction(input string) (*EducationAssignmentAddToCalendarAction, error) {
	vals := map[string]EducationAssignmentAddToCalendarAction{
		"none":                  EducationAssignmentAddToCalendarAction_None,
		"studentsandpublisher":  EducationAssignmentAddToCalendarAction_StudentsAndPublisher,
		"studentsandteamowners": EducationAssignmentAddToCalendarAction_StudentsAndTeamOwners,
		"studentsonly":          EducationAssignmentAddToCalendarAction_StudentsOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EducationAssignmentAddToCalendarAction(input)
	return &out, nil
}
