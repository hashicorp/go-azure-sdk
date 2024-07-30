package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubjectRightsRequestDataSubjectType string

const (
	SubjectRightsRequestDataSubjectType_CurrentEmployee     SubjectRightsRequestDataSubjectType = "currentEmployee"
	SubjectRightsRequestDataSubjectType_Customer            SubjectRightsRequestDataSubjectType = "customer"
	SubjectRightsRequestDataSubjectType_Faculty             SubjectRightsRequestDataSubjectType = "faculty"
	SubjectRightsRequestDataSubjectType_FormerEmployee      SubjectRightsRequestDataSubjectType = "formerEmployee"
	SubjectRightsRequestDataSubjectType_Other               SubjectRightsRequestDataSubjectType = "other"
	SubjectRightsRequestDataSubjectType_ProspectiveEmployee SubjectRightsRequestDataSubjectType = "prospectiveEmployee"
	SubjectRightsRequestDataSubjectType_Student             SubjectRightsRequestDataSubjectType = "student"
	SubjectRightsRequestDataSubjectType_Teacher             SubjectRightsRequestDataSubjectType = "teacher"
)

func PossibleValuesForSubjectRightsRequestDataSubjectType() []string {
	return []string{
		string(SubjectRightsRequestDataSubjectType_CurrentEmployee),
		string(SubjectRightsRequestDataSubjectType_Customer),
		string(SubjectRightsRequestDataSubjectType_Faculty),
		string(SubjectRightsRequestDataSubjectType_FormerEmployee),
		string(SubjectRightsRequestDataSubjectType_Other),
		string(SubjectRightsRequestDataSubjectType_ProspectiveEmployee),
		string(SubjectRightsRequestDataSubjectType_Student),
		string(SubjectRightsRequestDataSubjectType_Teacher),
	}
}

func (s *SubjectRightsRequestDataSubjectType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSubjectRightsRequestDataSubjectType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSubjectRightsRequestDataSubjectType(input string) (*SubjectRightsRequestDataSubjectType, error) {
	vals := map[string]SubjectRightsRequestDataSubjectType{
		"currentemployee":     SubjectRightsRequestDataSubjectType_CurrentEmployee,
		"customer":            SubjectRightsRequestDataSubjectType_Customer,
		"faculty":             SubjectRightsRequestDataSubjectType_Faculty,
		"formeremployee":      SubjectRightsRequestDataSubjectType_FormerEmployee,
		"other":               SubjectRightsRequestDataSubjectType_Other,
		"prospectiveemployee": SubjectRightsRequestDataSubjectType_ProspectiveEmployee,
		"student":             SubjectRightsRequestDataSubjectType_Student,
		"teacher":             SubjectRightsRequestDataSubjectType_Teacher,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SubjectRightsRequestDataSubjectType(input)
	return &out, nil
}
