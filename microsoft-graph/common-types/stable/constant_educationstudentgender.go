package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationStudentGender string

const (
	EducationStudentGender_Female EducationStudentGender = "female"
	EducationStudentGender_Male   EducationStudentGender = "male"
	EducationStudentGender_Other  EducationStudentGender = "other"
)

func PossibleValuesForEducationStudentGender() []string {
	return []string{
		string(EducationStudentGender_Female),
		string(EducationStudentGender_Male),
		string(EducationStudentGender_Other),
	}
}

func (s *EducationStudentGender) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEducationStudentGender(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEducationStudentGender(input string) (*EducationStudentGender, error) {
	vals := map[string]EducationStudentGender{
		"female": EducationStudentGender_Female,
		"male":   EducationStudentGender_Male,
		"other":  EducationStudentGender_Other,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EducationStudentGender(input)
	return &out, nil
}
