package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationUserPrimaryRole string

const (
	EducationUserPrimaryRole_Faculty EducationUserPrimaryRole = "faculty"
	EducationUserPrimaryRole_None    EducationUserPrimaryRole = "none"
	EducationUserPrimaryRole_Student EducationUserPrimaryRole = "student"
	EducationUserPrimaryRole_Teacher EducationUserPrimaryRole = "teacher"
)

func PossibleValuesForEducationUserPrimaryRole() []string {
	return []string{
		string(EducationUserPrimaryRole_Faculty),
		string(EducationUserPrimaryRole_None),
		string(EducationUserPrimaryRole_Student),
		string(EducationUserPrimaryRole_Teacher),
	}
}

func (s *EducationUserPrimaryRole) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEducationUserPrimaryRole(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEducationUserPrimaryRole(input string) (*EducationUserPrimaryRole, error) {
	vals := map[string]EducationUserPrimaryRole{
		"faculty": EducationUserPrimaryRole_Faculty,
		"none":    EducationUserPrimaryRole_None,
		"student": EducationUserPrimaryRole_Student,
		"teacher": EducationUserPrimaryRole_Teacher,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EducationUserPrimaryRole(input)
	return &out, nil
}
