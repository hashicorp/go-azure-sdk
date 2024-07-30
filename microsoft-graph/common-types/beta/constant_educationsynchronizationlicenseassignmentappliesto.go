package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationSynchronizationLicenseAssignmentAppliesTo string

const (
	EducationSynchronizationLicenseAssignmentAppliesTo_Faculty EducationSynchronizationLicenseAssignmentAppliesTo = "faculty"
	EducationSynchronizationLicenseAssignmentAppliesTo_None    EducationSynchronizationLicenseAssignmentAppliesTo = "none"
	EducationSynchronizationLicenseAssignmentAppliesTo_Student EducationSynchronizationLicenseAssignmentAppliesTo = "student"
	EducationSynchronizationLicenseAssignmentAppliesTo_Teacher EducationSynchronizationLicenseAssignmentAppliesTo = "teacher"
)

func PossibleValuesForEducationSynchronizationLicenseAssignmentAppliesTo() []string {
	return []string{
		string(EducationSynchronizationLicenseAssignmentAppliesTo_Faculty),
		string(EducationSynchronizationLicenseAssignmentAppliesTo_None),
		string(EducationSynchronizationLicenseAssignmentAppliesTo_Student),
		string(EducationSynchronizationLicenseAssignmentAppliesTo_Teacher),
	}
}

func (s *EducationSynchronizationLicenseAssignmentAppliesTo) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEducationSynchronizationLicenseAssignmentAppliesTo(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEducationSynchronizationLicenseAssignmentAppliesTo(input string) (*EducationSynchronizationLicenseAssignmentAppliesTo, error) {
	vals := map[string]EducationSynchronizationLicenseAssignmentAppliesTo{
		"faculty": EducationSynchronizationLicenseAssignmentAppliesTo_Faculty,
		"none":    EducationSynchronizationLicenseAssignmentAppliesTo_None,
		"student": EducationSynchronizationLicenseAssignmentAppliesTo_Student,
		"teacher": EducationSynchronizationLicenseAssignmentAppliesTo_Teacher,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EducationSynchronizationLicenseAssignmentAppliesTo(input)
	return &out, nil
}
