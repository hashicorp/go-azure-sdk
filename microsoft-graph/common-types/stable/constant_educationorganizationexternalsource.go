package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationOrganizationExternalSource string

const (
	EducationOrganizationExternalSource_Manual EducationOrganizationExternalSource = "manual"
	EducationOrganizationExternalSource_Sis    EducationOrganizationExternalSource = "sis"
)

func PossibleValuesForEducationOrganizationExternalSource() []string {
	return []string{
		string(EducationOrganizationExternalSource_Manual),
		string(EducationOrganizationExternalSource_Sis),
	}
}

func (s *EducationOrganizationExternalSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEducationOrganizationExternalSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEducationOrganizationExternalSource(input string) (*EducationOrganizationExternalSource, error) {
	vals := map[string]EducationOrganizationExternalSource{
		"manual": EducationOrganizationExternalSource_Manual,
		"sis":    EducationOrganizationExternalSource_Sis,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EducationOrganizationExternalSource(input)
	return &out, nil
}
