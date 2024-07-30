package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationClassExternalSource string

const (
	EducationClassExternalSource_Lms    EducationClassExternalSource = "lms"
	EducationClassExternalSource_Manual EducationClassExternalSource = "manual"
	EducationClassExternalSource_Sis    EducationClassExternalSource = "sis"
)

func PossibleValuesForEducationClassExternalSource() []string {
	return []string{
		string(EducationClassExternalSource_Lms),
		string(EducationClassExternalSource_Manual),
		string(EducationClassExternalSource_Sis),
	}
}

func (s *EducationClassExternalSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEducationClassExternalSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEducationClassExternalSource(input string) (*EducationClassExternalSource, error) {
	vals := map[string]EducationClassExternalSource{
		"lms":    EducationClassExternalSource_Lms,
		"manual": EducationClassExternalSource_Manual,
		"sis":    EducationClassExternalSource_Sis,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EducationClassExternalSource(input)
	return &out, nil
}
