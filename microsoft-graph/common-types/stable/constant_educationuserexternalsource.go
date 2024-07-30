package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationUserExternalSource string

const (
	EducationUserExternalSource_Manual EducationUserExternalSource = "manual"
	EducationUserExternalSource_Sis    EducationUserExternalSource = "sis"
)

func PossibleValuesForEducationUserExternalSource() []string {
	return []string{
		string(EducationUserExternalSource_Manual),
		string(EducationUserExternalSource_Sis),
	}
}

func (s *EducationUserExternalSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEducationUserExternalSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEducationUserExternalSource(input string) (*EducationUserExternalSource, error) {
	vals := map[string]EducationUserExternalSource{
		"manual": EducationUserExternalSource_Manual,
		"sis":    EducationUserExternalSource_Sis,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EducationUserExternalSource(input)
	return &out, nil
}
