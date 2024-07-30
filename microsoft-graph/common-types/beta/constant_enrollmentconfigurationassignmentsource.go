package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnrollmentConfigurationAssignmentSource string

const (
	EnrollmentConfigurationAssignmentSource_Direct     EnrollmentConfigurationAssignmentSource = "direct"
	EnrollmentConfigurationAssignmentSource_PolicySets EnrollmentConfigurationAssignmentSource = "policySets"
)

func PossibleValuesForEnrollmentConfigurationAssignmentSource() []string {
	return []string{
		string(EnrollmentConfigurationAssignmentSource_Direct),
		string(EnrollmentConfigurationAssignmentSource_PolicySets),
	}
}

func (s *EnrollmentConfigurationAssignmentSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEnrollmentConfigurationAssignmentSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEnrollmentConfigurationAssignmentSource(input string) (*EnrollmentConfigurationAssignmentSource, error) {
	vals := map[string]EnrollmentConfigurationAssignmentSource{
		"direct":     EnrollmentConfigurationAssignmentSource_Direct,
		"policysets": EnrollmentConfigurationAssignmentSource_PolicySets,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EnrollmentConfigurationAssignmentSource(input)
	return &out, nil
}
