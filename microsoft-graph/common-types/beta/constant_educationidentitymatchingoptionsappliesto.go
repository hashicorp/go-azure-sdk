package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationIdentityMatchingOptionsAppliesTo string

const (
	EducationIdentityMatchingOptionsAppliesTo_Faculty EducationIdentityMatchingOptionsAppliesTo = "faculty"
	EducationIdentityMatchingOptionsAppliesTo_None    EducationIdentityMatchingOptionsAppliesTo = "none"
	EducationIdentityMatchingOptionsAppliesTo_Student EducationIdentityMatchingOptionsAppliesTo = "student"
	EducationIdentityMatchingOptionsAppliesTo_Teacher EducationIdentityMatchingOptionsAppliesTo = "teacher"
)

func PossibleValuesForEducationIdentityMatchingOptionsAppliesTo() []string {
	return []string{
		string(EducationIdentityMatchingOptionsAppliesTo_Faculty),
		string(EducationIdentityMatchingOptionsAppliesTo_None),
		string(EducationIdentityMatchingOptionsAppliesTo_Student),
		string(EducationIdentityMatchingOptionsAppliesTo_Teacher),
	}
}

func (s *EducationIdentityMatchingOptionsAppliesTo) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEducationIdentityMatchingOptionsAppliesTo(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEducationIdentityMatchingOptionsAppliesTo(input string) (*EducationIdentityMatchingOptionsAppliesTo, error) {
	vals := map[string]EducationIdentityMatchingOptionsAppliesTo{
		"faculty": EducationIdentityMatchingOptionsAppliesTo_Faculty,
		"none":    EducationIdentityMatchingOptionsAppliesTo_None,
		"student": EducationIdentityMatchingOptionsAppliesTo_Student,
		"teacher": EducationIdentityMatchingOptionsAppliesTo_Teacher,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EducationIdentityMatchingOptionsAppliesTo(input)
	return &out, nil
}
