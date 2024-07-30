package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationSchoolExternalSource string

const (
	EducationSchoolExternalSource_Lms    EducationSchoolExternalSource = "lms"
	EducationSchoolExternalSource_Manual EducationSchoolExternalSource = "manual"
	EducationSchoolExternalSource_Sis    EducationSchoolExternalSource = "sis"
)

func PossibleValuesForEducationSchoolExternalSource() []string {
	return []string{
		string(EducationSchoolExternalSource_Lms),
		string(EducationSchoolExternalSource_Manual),
		string(EducationSchoolExternalSource_Sis),
	}
}

func (s *EducationSchoolExternalSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEducationSchoolExternalSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEducationSchoolExternalSource(input string) (*EducationSchoolExternalSource, error) {
	vals := map[string]EducationSchoolExternalSource{
		"lms":    EducationSchoolExternalSource_Lms,
		"manual": EducationSchoolExternalSource_Manual,
		"sis":    EducationSchoolExternalSource_Sis,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EducationSchoolExternalSource(input)
	return &out, nil
}
