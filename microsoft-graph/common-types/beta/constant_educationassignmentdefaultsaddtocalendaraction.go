package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationAssignmentDefaultsAddToCalendarAction string

const (
	EducationAssignmentDefaultsAddToCalendarAction_None                  EducationAssignmentDefaultsAddToCalendarAction = "none"
	EducationAssignmentDefaultsAddToCalendarAction_StudentsAndPublisher  EducationAssignmentDefaultsAddToCalendarAction = "studentsAndPublisher"
	EducationAssignmentDefaultsAddToCalendarAction_StudentsAndTeamOwners EducationAssignmentDefaultsAddToCalendarAction = "studentsAndTeamOwners"
	EducationAssignmentDefaultsAddToCalendarAction_StudentsOnly          EducationAssignmentDefaultsAddToCalendarAction = "studentsOnly"
)

func PossibleValuesForEducationAssignmentDefaultsAddToCalendarAction() []string {
	return []string{
		string(EducationAssignmentDefaultsAddToCalendarAction_None),
		string(EducationAssignmentDefaultsAddToCalendarAction_StudentsAndPublisher),
		string(EducationAssignmentDefaultsAddToCalendarAction_StudentsAndTeamOwners),
		string(EducationAssignmentDefaultsAddToCalendarAction_StudentsOnly),
	}
}

func (s *EducationAssignmentDefaultsAddToCalendarAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEducationAssignmentDefaultsAddToCalendarAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEducationAssignmentDefaultsAddToCalendarAction(input string) (*EducationAssignmentDefaultsAddToCalendarAction, error) {
	vals := map[string]EducationAssignmentDefaultsAddToCalendarAction{
		"none":                  EducationAssignmentDefaultsAddToCalendarAction_None,
		"studentsandpublisher":  EducationAssignmentDefaultsAddToCalendarAction_StudentsAndPublisher,
		"studentsandteamowners": EducationAssignmentDefaultsAddToCalendarAction_StudentsAndTeamOwners,
		"studentsonly":          EducationAssignmentDefaultsAddToCalendarAction_StudentsOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EducationAssignmentDefaultsAddToCalendarAction(input)
	return &out, nil
}
