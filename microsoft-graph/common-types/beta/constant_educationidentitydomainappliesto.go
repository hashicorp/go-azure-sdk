package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationIdentityDomainAppliesTo string

const (
	EducationIdentityDomainAppliesTo_Faculty EducationIdentityDomainAppliesTo = "faculty"
	EducationIdentityDomainAppliesTo_None    EducationIdentityDomainAppliesTo = "none"
	EducationIdentityDomainAppliesTo_Student EducationIdentityDomainAppliesTo = "student"
	EducationIdentityDomainAppliesTo_Teacher EducationIdentityDomainAppliesTo = "teacher"
)

func PossibleValuesForEducationIdentityDomainAppliesTo() []string {
	return []string{
		string(EducationIdentityDomainAppliesTo_Faculty),
		string(EducationIdentityDomainAppliesTo_None),
		string(EducationIdentityDomainAppliesTo_Student),
		string(EducationIdentityDomainAppliesTo_Teacher),
	}
}

func (s *EducationIdentityDomainAppliesTo) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEducationIdentityDomainAppliesTo(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEducationIdentityDomainAppliesTo(input string) (*EducationIdentityDomainAppliesTo, error) {
	vals := map[string]EducationIdentityDomainAppliesTo{
		"faculty": EducationIdentityDomainAppliesTo_Faculty,
		"none":    EducationIdentityDomainAppliesTo_None,
		"student": EducationIdentityDomainAppliesTo_Student,
		"teacher": EducationIdentityDomainAppliesTo_Teacher,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EducationIdentityDomainAppliesTo(input)
	return &out, nil
}
