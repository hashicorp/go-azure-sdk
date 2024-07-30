package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationAssignmentDefaultsAddedStudentAction string

const (
	EducationAssignmentDefaultsAddedStudentAction_AssignIfOpen EducationAssignmentDefaultsAddedStudentAction = "assignIfOpen"
	EducationAssignmentDefaultsAddedStudentAction_None         EducationAssignmentDefaultsAddedStudentAction = "none"
)

func PossibleValuesForEducationAssignmentDefaultsAddedStudentAction() []string {
	return []string{
		string(EducationAssignmentDefaultsAddedStudentAction_AssignIfOpen),
		string(EducationAssignmentDefaultsAddedStudentAction_None),
	}
}

func (s *EducationAssignmentDefaultsAddedStudentAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEducationAssignmentDefaultsAddedStudentAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEducationAssignmentDefaultsAddedStudentAction(input string) (*EducationAssignmentDefaultsAddedStudentAction, error) {
	vals := map[string]EducationAssignmentDefaultsAddedStudentAction{
		"assignifopen": EducationAssignmentDefaultsAddedStudentAction_AssignIfOpen,
		"none":         EducationAssignmentDefaultsAddedStudentAction_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EducationAssignmentDefaultsAddedStudentAction(input)
	return &out, nil
}
