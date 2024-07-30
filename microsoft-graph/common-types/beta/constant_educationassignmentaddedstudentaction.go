package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationAssignmentAddedStudentAction string

const (
	EducationAssignmentAddedStudentAction_AssignIfOpen EducationAssignmentAddedStudentAction = "assignIfOpen"
	EducationAssignmentAddedStudentAction_None         EducationAssignmentAddedStudentAction = "none"
)

func PossibleValuesForEducationAssignmentAddedStudentAction() []string {
	return []string{
		string(EducationAssignmentAddedStudentAction_AssignIfOpen),
		string(EducationAssignmentAddedStudentAction_None),
	}
}

func (s *EducationAssignmentAddedStudentAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEducationAssignmentAddedStudentAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEducationAssignmentAddedStudentAction(input string) (*EducationAssignmentAddedStudentAction, error) {
	vals := map[string]EducationAssignmentAddedStudentAction{
		"assignifopen": EducationAssignmentAddedStudentAction_AssignIfOpen,
		"none":         EducationAssignmentAddedStudentAction_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EducationAssignmentAddedStudentAction(input)
	return &out, nil
}
